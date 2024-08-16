package main

import (
	"fmt"
	"log"

	"api/api"
	red "api/api/handler/budgeting"
	"api/config"

	"github.com/go-redis/redis/v8"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	cnf := config.Load()
	authservice, err := grpc.NewClient(cnf.AUTH_PORT, grpc.WithTransportCredentials(insecure.NewCredentials())) // Update the address
	if err != nil {
		log.Fatalf("Failed to connect to authservice service %v", err)
	}
	defer authservice.Close()

	companyService, err := grpc.NewClient("budgeting_service"+ cnf.BUDGETING_PORT, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Failed to connect to learning service service %v", err)
	}

	CL:=redis.NewClient(&redis.Options{
		Addr: "redis1:6370",
	})

	c := red.NewInMemoryStorage(CL)

	router := api.NewGin(companyService, c)

	fmt.Println("API Gateway running on http://localhost:8082")
	if err := router.Run(":8082"); err != nil {
		log.Fatalf("Failed to connect to gin engine: %v", err)
	}

}
