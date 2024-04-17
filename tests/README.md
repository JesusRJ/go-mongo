# Run Tests

```bash
cd tests
# prepare test database
docker compose up -d

# run all tests
go test -race -count=1 ./...
```

> Ensure the port `27017` is free to use or change `docker-compose.yaml` and `tests_test.go` files to appropriate port.