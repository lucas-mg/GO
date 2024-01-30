
# Using GO create a deploy file

Go API that receives parameters (client_id, number_of_nodes, and client_name) and generates a Kubernetes manifest YAML file with the specified values:

## Rodando localmente

Run the Go API:

```bash
 go run main.go
```

Make a POST request to the API:
```bash
 curl -X POST -H "Content-Type: application/json" -d '{"client_id": "myclient", "number_of_nodes": 3, "client_name": "myapp"}' http://localhost:8080/generate-manifest

```


