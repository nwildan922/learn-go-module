package main

import (
	"log"
	"net"
	"os"

	"github.com/google/uuid"
	"github.com/joho/godotenv"
	h "github.com/nwildan922/learn-go-module/handler"
	"github.com/nwildan922/learn-go-module/pkg/db"
	"github.com/nwildan922/learn-go-module/proto/counterpb"
	"github.com/nwildan922/learn-go-module/repository"
	"github.com/nwildan922/learn-go-module/service"
	"google.golang.org/grpc"
)

func main() {

	Load()

	appId := uuid.New()

	port := os.Getenv("GRPC_PORT")

	if port == "" {
		port = "50051"
	}

	dsn := os.Getenv("DB_DSN")
	log.Println("dsn")
	log.Println(dsn)
	database := db.NewDatabase(dsn)
	log.Println("🚀 succes db connect", port)

	repo := repository.NewCounterRepository(database)
	svc := service.NewCounterService(repo, appId.String())
	handler := h.NewCounterHandler(svc)

	log.Println("🚀 try listen port", port)

	lis, err := net.Listen("tcp", ":"+port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	log.Println("🚀 try listen port", port)

	// 🔥 Create gRPC server
	grpcServer := grpc.NewServer()

	// 🔧 Register service
	counterpb.RegisterCounterServiceServer(grpcServer, handler)

	log.Println("🚀 gRPC server running on port", port)

	// 🔥 Start server
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

func Load() {
	env := os.Getenv("APP_ENV")
	if env == "" {
		env = "develop"
	}

	file := ".env." + env

	// try specific env first
	if err := godotenv.Load(file); err != nil {
		// fallback to default .env
		if err := godotenv.Load(); err != nil {
			log.Println("No env file found, using system env")
		}
	}

	log.Println("Environment loaded:", env)
}
