package grpc

import (
	"context"
	"log"
	"net"
	"sync"

	"github.com/vegitobluefan/task-manager/domain"
	pb "github.com/vegitobluefan/task-manager/proto"
	"github.com/vegitobluefan/task-manager/usecase"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type Server struct {
	pb.UnimplementedTaskServiceServer
	uc usecase.TaskUseCase
	mu sync.Mutex
}

func NewServer(uc usecase.TaskUseCase) *Server {
	return &Server{uc: uc}
}

func (s *Server) GetTask(ctx context.Context, req *pb.TaskRequest) (*pb.TaskResponse, error) {
	task, err := s.uc.GetTask(req.Id)
	if err != nil {
		return nil, err
	}
	return &pb.TaskResponse{
		Id:      task.ID,
		Type:    task.Type,
		Status:  task.Status,
		Payload: task.Payload,
		Result:  task.Result,
	}, nil
}

func (s *Server) ListTasks(ctx context.Context, req *pb.TaskListRequest) (*pb.TaskListResponse, error) {
	tasks, err := s.uc.ListTasks()
	if err != nil {
		return nil, err
	}
	resp := &pb.TaskListResponse{}
	for _, t := range tasks {
		resp.Tasks = append(resp.Tasks, &pb.TaskResponse{
			Id:      t.ID,
			Type:    t.Type,
			Status:  t.Status,
			Payload: t.Payload,
			Result:  t.Result,
		})
	}
	return resp, nil
}

func (s *Server) CreateTask(ctx context.Context, req *pb.CreateTaskRequest) (*pb.CreateTaskResponse, error) {
	task := &domain.Task{
		Type:    req.Type,
		Payload: req.Payload,
	}
	id, err := s.uc.Enqueue(task)
	if err != nil {
		return nil, err
	}
	return &pb.CreateTaskResponse{Id: id}, nil
}

func RunGRPCServer(uc usecase.TaskUseCase, port string) error {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		return err
	}
	grpcServer := grpc.NewServer()
	pb.RegisterTaskServiceServer(grpcServer, NewServer(uc))

	reflection.Register(grpcServer)

	log.Printf("gRPC сервер запущен на %s", port)
	return grpcServer.Serve(lis)
}
