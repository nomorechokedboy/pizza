## Getting Started

First, run the development server:

```bash
go mod download # install deps if run on local

docker compose up postgres minio-svc -d # run database and other service for local dev

pn api:dev # run this at root level
# OR
air # run this at apps/api
```

Open [http://localhost:3001](http://localhost:3001) with your browser to see the result.

You can start editing the api by modifying `main.go`. The api auto-updates as you edit the file.

[API routes](http://localhost:3001/docs) can be accessed on [http://localhost:3001/docs](http://localhost:3001/docs)

## Testing

Test files should be place next to it implementation for easier to maintain

- You should write unit test for use cases
- For the repository, we will use integration test
- You can refer to my test files for more information
- Finally, we will use e2e to test the endpoints

## Learn More

To learn more about the hexagonal/clean architecture, take a look at the following resources:

- [Fiber Examples](https://github.com/gofiber/recipes) - learn about clean architecture layers and how they work.
- [Learn To Write Test In Clean Architecture](https://github.com/arielizuardi/golang-backend-blog) - a clean architecture realworld tutorial/example from gojek developer with unit/integration/e2e tests.

## Deploy on Okteto

The easiest way to deploy this api is to use the [Okteto Platform](https://okteto.com). I will support deploy with kubernetes in the future.
