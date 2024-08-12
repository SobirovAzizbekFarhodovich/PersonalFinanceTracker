package main

import (
	"log"
	"net"

	pb "budgeting/genprotos"
	"budgeting/service"
	"budgeting/storage/mongo"
	"google.golang.org/grpc"
)

func main(){
	db, err := mongo.ConnectMongo()
	if err != nil {
		log.Fatal(err)
	}

	liss, err := net.Listen("tcp", ":50055")
	if err != nil {
		log.Fatal(err)
	}

	s := grpc.NewServer()
	pb.RegisterAccountServiceServer(s, service.NewAccountService(&db))
	pb.RegisterBudgetServiceServer(s, service.NewBudgetService(&db))
	pb.RegisterCategoryServiceServer(s, service.NewCategoryService(&db))
	pb.RegisterGoalServiceServer(s, service.NewGoalService(&db))
	pb.RegisterTransactionServiceServer(s, service.NewTransactionService(&db))


	log.Printf("server listening at %v", liss.Addr())
	if err := s.Serve(liss); err != nil {
		log.Fatalf("failed to serve %v", err)
	}
}
