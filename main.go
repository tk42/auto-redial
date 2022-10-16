package main

import (
	"database/sql"
	"fmt"
	"log"
	"net"

	bufbuild "github.com/tk42/auto-redial/gen/proto/golang/github.com/tk42/auto-redial"
	"github.com/tk42/auto-redial/service"

	_ "github.com/jackc/pgx/v5/stdlib"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	log.Print("server is starting...")
	dsn := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable&TimeZone=Asia/Tokyo",
		"postgres", "e8a48653851e28c69d0506508fb27fc5", "db", "5432", "postgres",
	)
	client, err := sql.Open("pgx", dsn)
	if err != nil {
		log.Fatalf("failed to connect to db: %s", err)
	}
	defer client.Close()

	metric_svc := service.NewMetricServiceServer(client)
	scammer_svc := service.NewScammerServiceServer(client)
	callhistory_svc := service.NewCallHistoryServiceServer(client)
	matching_svc := service.NewMatchingServiceServer(client)

	server := grpc.NewServer()
	reflection.Register(server) // Failed to list services: server does not support the reflection API
	bufbuild.RegisterMetricStoreServiceServer(server, metric_svc)
	bufbuild.RegisterScammerStoreServiceServer(server, scammer_svc)
	bufbuild.RegisterCallHistoryStoreServiceServer(server, callhistory_svc)
	bufbuild.RegisterMatchingStoreServiceServer(server, matching_svc)

	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %s", err)
	}

	if err := server.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %s", err)
	}
}
