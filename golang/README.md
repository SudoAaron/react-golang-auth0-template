# Golang Backend for Auth0 Integration

## Introduction

This Golang application serves as a backend API for demonstrating Auth0 integration. It provides secure endpoints that work in conjunction with the React frontend, showcasing user authentication and authorization using Auth0.

## Prerequisites

Before starting, ensure you have the following:

1. Go (version 1.13 or later)
2. An Auth0 account for creating applications and APIs
3. A postgres database

## Setting Up Auth0

\*\* Be sure to use the same application and API as the React application.
Create an Auth0 application and API on the Auth0 dashboard.
Note down the domain and audience from your Auth0 setup.

## Setting Up Postgres

1. Install postgres via your preferred method. 
2. Create postgres user and password.
3. Modify the .env values as needed.

\*\* Note that the tables will automatically be set up when starting the application.

## Configuration

Rename .env.example to .env and fill in the values:

```bash
AUTH0_DOMAIN={domain} # Your Auth0 application domain
AUTH0_AUDIENCE={audience} # Your Auth0 API audience
PG_HOST="localhost"
PG_PORT="5432"
PG_USER={postgres_user}
PG_PASSWORD={postgres_password}
PG_DATABASE="react-golang-auth0-template"
```

## Installation

Install necessary dependencies by running:

```bash
go mod tidy
```

## Running the Application

Start the server with:

```bash
go run cmd/main.go
```

## API Endpoints

The API provides several endpoints for authentication and data retrieval. These endpoints are secured and require valid Auth0 tokens.

## Contributing

We welcome contributions to this project!

## Support and Contact

If you have any queries or need support, please open an issue.
