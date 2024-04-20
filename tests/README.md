# Run Tests

```bash
cd tests
# prepare test database
docker compose up -d

# run all tests
go test -race -count=1 ./...
```

> Make sure port `27017` is available for use, or modify the `docker-compose.yaml` and `tests_test.go` files to use an appropriate port.