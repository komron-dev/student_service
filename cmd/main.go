package main

import (
	"net"

	"student_service/config"
	pb "student_service/genproto/student_service"
	"student_service/pkg/db"
	"student_service/pkg/logger"
	"student_service/services"
	"student_service/services/grpcClient"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	cfg := config.Load()

	log := logger.New(cfg.LogLevel, "student-service")
	defer logger.Cleanup(log)

	log.Info("main: sqlxConfig",
		logger.String("host", cfg.PostgresHost),
		logger.Int("port", cfg.PostgresPort),
		logger.String("database", cfg.PostgresDatabase),
	)

	grpcClient, err := grpcClient.New(cfg)
	if err != nil {
		log.Error("grpc dial error", logger.Error(err))
	}

	connDB, err, _ := db.ConnectToDbAndAlsoForSuite(cfg)
	if err != nil {
		log.Fatal("sqlx connection to postgres error", logger.Error(err))
	}

	studentService := services.NewStudentService(connDB, *grpcClient, log)

	lis, err := net.Listen("tcp", cfg.RPCPort)
	if err != nil {
		log.Fatal("Error while listening: %v", logger.Error(err))
	}

	s := grpc.NewServer()
	reflection.Register(s)

	pb.RegisterStudentServiceServer(s, studentService)
	log.Info("main: server running", logger.String("port:", cfg.RPCPort))

	if err := s.Serve(lis); err != nil {
		log.Fatal("Error while listening: %v", logger.Error(err))
	}
}
