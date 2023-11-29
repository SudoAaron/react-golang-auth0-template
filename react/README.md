## Getting Started with Auth0 Frontend

First, create an Auth0 application and API. I would recommend using `http://localhost:3001/` for the API audience.

Next, rename .env.example to .env.local. Fill out the relevant data from your Auth0 Application.

```
NEXT_PUBLIC_AUTH0_DOMAIN={domain} ## This data is associated with the Auth0 application
NEXT_PUBLIC_AUTH0_CLIENT_ID={client_id}  ## This data is associated with the Auth0 application
NEXT_PUBLIC_AUTH0_AUDIENCE={audience} ## This data is associated with the Auth0 API
NEXT_PUBLIC_API_ENDPOINT="http://localhost:3001" ## This should be the default endpoint for the Golang API.
```

Next, install any dependencies:

```bash
npm i
```

Next, run the development server:

```bash
npm run dev
```

Open [http://localhost:3000](http://localhost:3000) with your browser to see the result.
