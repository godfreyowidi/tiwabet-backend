package main

import (
	"fmt"
	"log"
	"net"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/handler/extension"
	"github.com/99designs/gqlgen/graphql/handler/lru"
	"github.com/99designs/gqlgen/graphql/handler/transport"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/godfreyowidi/tiwabet-backend/api"
	"github.com/godfreyowidi/tiwabet-backend/gql-gateway/graph"
	"github.com/godfreyowidi/tiwabet-backend/gql-gateway/resolvers"
	"github.com/godfreyowidi/tiwabet-backend/infra"
	"github.com/godfreyowidi/tiwabet-backend/infra/db"
	"github.com/godfreyowidi/tiwabet-backend/proto/userpb"
	"github.com/vektah/gqlparser/v2/ast"

	"google.golang.org/grpc"
)

func startGRPCServer() {
	// Load database URL from environment variables
	dbURL := os.Getenv("DATABASE_URL")
	if dbURL == "" {
		log.Fatal("DATABASE_URL is not set in the environment")
	} else {
		fmt.Println("✅ DATABASE_URL loaded successfully")
	}

	// Initialize database connection
	database, err := db.NewPostgresDB()
	if err != nil {
		log.Fatalf("❌ Failed to connect to database: %v", err)
	}
	defer database.Close()
	fmt.Println("✅ Connected to PostgreSQL successfully!")

	// Define gRPC server listener
	port := os.Getenv("GRPC_PORT")
	if port == "" {
		port = "50051"
	}
	listener, err := net.Listen("tcp", ":"+port)
	if err != nil {
		log.Fatalf("❌ Failed to listen: %v", err)
	}

	userRepo := infra.NewUserRepository(database.Pool)

	userServer := api.NewUserService(userRepo)
	// Create new gRPC server
	grpcServer := grpc.NewServer()

	// Register UserService
	userpb.RegisterUserServiceServer(grpcServer, userServer)

	log.Printf("🚀 gRPC Server running on port %s", port)
	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("❌ Failed to serve: %v", err)
	}
}

func startGraphServer() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8081"
	}

	srv := handler.New(graph.NewExecutableSchema(graph.Config{Resolvers: &resolvers.Resolver{}}))

	srv.AddTransport(transport.Options{})
	srv.AddTransport(transport.GET{})
	srv.AddTransport(transport.POST{})

	srv.SetQueryCache(lru.New[*ast.QueryDocument](1000))

	srv.Use(extension.Introspection{})
	srv.Use(extension.AutomaticPersistedQuery{
		Cache: lru.New[string](100),
	})

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	log.Printf("connect to http://0.0.0.0:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe("0.0.0.0:"+port, nil))
}

func main() {
	go startGraphServer()
	startGRPCServer()
}
