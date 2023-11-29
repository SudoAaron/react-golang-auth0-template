# React Frontend for Auth0 Integration

## Introduction

This React application serves as a frontend for demonstrating Auth0 integration. It showcases user authentication and secure API integration using Auth0 services. The application is designed to be paired with a Golang API backend.

## Prerequisites

Before you begin, ensure you have the following installed:

1. Node.js (version 12 or later)
2. npm (usually comes with Node.js)
3. An Auth0 account for creating applications and APIs

## Setting Up Auth0

Create an Auth0 application and API on the Auth0 dashboard.
For the API, use http://localhost:3001/ as the audience.
Note down the domain, client ID, and audience from your Auth0 setup.

## Configuration

Rename .env.example to .env.local and fill in the values:

```bash
NEXT_PUBLIC_AUTH0_DOMAIN={domain} # Your Auth0 application domain
NEXT_PUBLIC_AUTH0_CLIENT_ID={client_id} # Your Auth0 application client ID
NEXT_PUBLIC_AUTH0_AUDIENCE={audience} # Your Auth0 API audience
NEXT_PUBLIC_API_ENDPOINT="http://localhost:3001" # Endpoint for the Golang API
```

## Installation

Install the necessary dependencies by running:

```bash
npm install
```

## Running the Application

Start the development server with:

```bash
npm run dev
```

Access the application at http://localhost:3000.

## Usage

Once the application is running:

You'll be presented with an authentication screen.
After logging in, you can interact with secured API endpoints.
Explore features demonstrating secure API calls and user authentication.

## Contributing

We welcome contributions to this project!

## Support and Contact

If you have any queries or need support, please open an issue.
