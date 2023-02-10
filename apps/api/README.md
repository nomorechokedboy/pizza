## Getting Started

First, run the development server:

```bash
<<<<<<< HEAD
yarn dev
```

Open [http://localhost:3000](http://localhost:3000) with your browser to see the result.

You can start editing the page by modifying `pages/index.js`. The page auto-updates as you edit the file.

[API routes](https://nextjs.org/docs/api-routes/introduction) can be accessed on [http://localhost:3000/api/hello](http://localhost:3000/api/hello). This endpoint can be edited in `pages/api/hello.js`.

The `pages/api` directory is mapped to `/api/*`. Files in this directory are treated as [API routes](https://nextjs.org/docs/api-routes/introduction) instead of React pages.

## Learn More

To learn more about Next.js, take a look at the following resources:

- [Next.js Documentation](https://nextjs.org/docs) - learn about Next.js features and API.
- [Learn Next.js](https://nextjs.org/learn) - an interactive Next.js tutorial.

You can check out [the Next.js GitHub repository](https://github.com/vercel/next.js/) - your feedback and contributions are welcome!

## Deploy on Vercel

The easiest way to deploy your Next.js app is to use the [Vercel Platform](https://vercel.com/new?utm_source=github.com&utm_medium=referral&utm_campaign=turborepo-readme) from the creators of Next.js.

Check out our [Next.js deployment documentation](https://nextjs.org/docs/deployment) for more details.
=======
docker-compose up -d # run database and other service for local dev

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
>>>>>>> a8dc111 (chore(api/README): add instruction to run dev and test)
