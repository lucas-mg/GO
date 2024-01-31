package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"sync"

	"github.com/gorilla/mux"
)

// NFT represents a Non-Fungible Token
type NFT struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

// NFTStore represents an in-memory storage for NFTs
type NFTStore struct {
	mu   sync.Mutex
	nfts map[string]NFT
}

// CreateNFT creates a new NFT
func (store *NFTStore) CreateNFT(w http.ResponseWriter, r *http.Request) {
	var nft NFT
	if err := json.NewDecoder(r.Body).Decode(&nft); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	store.mu.Lock()
	defer store.mu.Unlock()

	// Generate a unique ID for the NFT (you might want to use a library for this in production)
	nft.ID = fmt.Sprintf("nft%d", len(store.nfts)+1)

	store.nfts[nft.ID] = nft

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(nft)
}

// GetNFTs returns a list of all NFTs
func (store *NFTStore) GetNFTs(w http.ResponseWriter, r *http.Request) {
	store.mu.Lock()
	defer store.mu.Unlock()

	nftsList := make([]NFT, 0, len(store.nfts))
	for _, nft := range store.nfts {
		nftsList = append(nftsList, nft)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(nftsList)
}

func main() {
	router := mux.NewRouter()

	nftStore := &NFTStore{
		nfts: make(map[string]NFT),
	}

	router.HandleFunc("/nfts", nftStore.CreateNFT).Methods("POST")
	router.HandleFunc("/nfts", nftStore.GetNFTs).Methods("GET")

	fmt.Println("Server is running on :8080")
	http.ListenAndServe(":8080", router)
}
