package grpcClient

import (
	"fmt"
	"student_service/config"
	pbs "student_service/genproto/student_service"
	pbt "student_service/genproto/task_service"

	"google.golang.org/grpc"
	// "google.golang.org/grpc"
)

type IService interface {
	StudentService() pbs.StudentServiceClient
	TaskService() pbt.TaskServiceClient
}

type GrpcClient struct {
	cfg         config.Config
	connections map[string]interface{}
}

func New(cfg config.Config) (*GrpcClient, error) {
	connTask, err := grpc.Dial(
		fmt.Sprintf("%s:%d", cfg.TaskServiceHost, cfg.TaskServicePort),
		grpc.WithInsecure())
	if err != nil {
		return nil, fmt.Errorf("taks service dial host: %s port:%d, err:%s",
			cfg.TaskServiceHost, cfg.TaskServicePort, err)
	}
	return &GrpcClient{
		cfg: cfg,
		connections: map[string]interface{}{
			"student_task_service": pbt.NewTaskServiceClient(connTask),
		},
	}, nil
}

func (g *GrpcClient) TaskService() pbs.StudentServiceClient {
	return g.connections["student_task_service"].(pbs.StudentServiceClient)
}
