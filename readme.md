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

## Problems

1. I want to reference _user.proto_ with _bet.proto_ for gRPC and have tried (several variations) to import user into bet like _import "tiwabet-backend/proto/bet.proto"_ and I cant get it to work

2. You can comment out the bets table in the meantime - its buggy

## Ivan's Review

1. This file is not well-formatted. You can download a Markdown Linter and clear out all the warnings you have
2. When you run the command `docker-compose up --build`, you'll get back this error:

    ```text
    tiwabet-migrate    | error: migration failed: relation "events" does not exist in line 0: CREATE TABLE IF NOT EXISTS bets (
    tiwabet-migrate    |     id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    tiwabet-migrate    |     user_id UUID NOT NULL,
    tiwabet-migrate    |     event_id UUID NOT NULL,  -- Links to a sports event - events table TODO
    tiwabet-migrate    |     amount NUMERIC(10,2) NOT NULL CHECK (amount > 0),
    tiwabet-migrate    |     odds NUMERIC(5,2) NOT NULL CHECK (odds > 0), -- Betting odds (1.50, 2.75) - these are to be mapped from an external api
    tiwabet-migrate    |     bet_type TEXT NOT NULL CHECK (bet_type IN ('SINGLE', 'MULTI', 'SYSTEM')),
    tiwabet-migrate    |     status TEXT NOT NULL DEFAULT 'PENDING' CHECK (status IN ('PENDING', 'WON', 'LOST', 'CANCELLED')),
    tiwabet-migrate    |     outcome TEXT CHECK (outcome IN ('WIN', 'LOSE', 'VOID')),
    tiwabet-migrate    |     placed_at TIMESTAMP DEFAULT NOW(),
    tiwabet-migrate    |     updated_at TIMESTAMP DEFAULT NOW(),
    tiwabet-migrate    | 
    tiwabet-migrate    |     -- Foreign Key Constraints - these were suppose to be the same as .proto descriptions
    tiwabet-migrate    |     CONSTRAINT fk_bet_user FOREIGN KEY (user_id) REFERENCES users (id) ON DELETE CASCADE,
    tiwabet-migrate    |     CONSTRAINT fk_bet_event FOREIGN KEY (event_id) REFERENCES events (id) ON DELETE CASCADE
    tiwabet-migrate    | );
    ```

3. The container `tiwabet-migrate` messed up with migrations (due to the previous point), and yield back `tiwabet-migrate    | error: Dirty database version 2. Fix and force version.`
4. Add a section in the `readme.md` file to explain how to run the code without using a container. This is helpful for debugging purposes (since it's easier to debug something local instead of remote)
5. Add a section in the `readme.md` file where you tell how to invoke the endpoints you're exposing from the project. Could be the `cURL` requests to start with. This will guide the reader on how to test out the application (beyond a successful build) and it will also provide a starting point for looking at the code
