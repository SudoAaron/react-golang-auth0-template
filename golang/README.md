## Getting Started with Auth0 Backend

First, rename .env.example to .env Fill out the relevant data from your Auth0 Application.

```
AUTH0_DOMAIN={domain} ## This data is associated with the Auth0 application
AUTH0_AUDIENCE={audience} ## This data is associated with the Auth0 API
```

Next, install any dependencies:

```bash
go mod tidy
```

Next, start the server:

```bash
go run cmd/main.go
```
