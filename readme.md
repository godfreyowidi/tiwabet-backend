# Demo Project

### Demo Project Showcasing Domain-Driven Design (DDD) with gRPC, GraphQL, and PGX for Scalable Backend Development

This project demonstrates the implementation of Domain-Driven Design (DDD) using gRPC for efficient communication, GraphQL for flexible data querying, and PGX for high-performance PostgreSQL interactions, creating a scalable and maintainable backend architecture.

### Technologies
- Go
- gqlgen - Go graphql library
- gRPC
- Docker
- PGX - Go toolkit for PostgresQL

## Project Setup
1. Clone the repo:
```
git clone github.com/godfreyowidi/tiwabet-backend
```

2. CD into the directory and run:
```
go mod tidy
```
3. Verify no missing dependencies:
```
go mod verify
```

4. On the project root, create a .env file (or config.yml) and add the following secrets:
- `POSTGRES_USER={user}`
- `POSTGRES_PASSWORD={user_password}`
- `POSTGRES_DB={db_name}`
- `DATABASE_URL={your_database_url}`
- `GRPC_PORT=50051`

5. If you dont have docker installed, you can follow this link to install, other run:
```
docker-compose up --build   
```

6. Check for the error on the terminal and sort them out appropriately