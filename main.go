package main

import (
	"task/handler"
	pb "task/proto"

	"github.com/micro/micro/v3/service"
	"github.com/micro/micro/v3/service/logger"
)

func main() {
	// Create service
	srv := service.New(
		service.Name("task"),
		service.Version("latest"),
	)

	// Register handler
	pb.RegisterTaskHandler(srv.Server(), new(handler.Task))

	// Run service
	if err := srv.Run(); err != nil {
		logger.Fatal(err)
	}
}
