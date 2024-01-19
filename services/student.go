package services

import (
	"context"

	pbs "student_service/genproto/student_service"
	l "student_service/pkg/logger"
	"student_service/services/grpcClient"
	"student_service/storage"

	"github.com/jmoiron/sqlx"
)

type StudentService struct {
	storage storage.IStorage
	logger  l.Logger
}

func NewStudentService(db *sqlx.DB, client grpcClient.GrpcClient, log l.Logger) *StudentService {
	return &StudentService{
		storage: storage.NewStoragePg(db),
		logger:  log,
	}
}

func (s *StudentService) CreateStudent(ctx context.Context, req *pbs.CreateStudentReq) (*pbs.CreateStudentRes, error) {
	res, err := s.storage.Student().CreateStudent(req)
	if err != nil {
		s.logger.Error("error while creating a student", l.Error(err))
		return nil, err
	}
	return res, nil
}

func (s *StudentService) GetStudent(ctx context.Context, req *pbs.ById) (*pbs.GetStudentRes, error) {
	res, err := s.storage.Student().GetStudent(req)
	if err != nil {
		s.logger.Error("error while getting a student", l.Error(err))
		return nil, err
	}
	return res, nil
}

func (s *StudentService) UpdateStudent(ctx context.Context, req *pbs.UpdateStudentReq) (*pbs.Success, error) {
	res, err := s.storage.Student().UpdateStudent(req)
	if err != nil {
		s.logger.Error("error while updating a student", l.Error(err))
		return nil, err
	}
	return res, nil
}

func (s *StudentService) DeleteStudent(ctx context.Context, req *pbs.ById) (*pbs.Success, error) {
	res, err := s.storage.Student().DeleteStudent(req)
	if err != nil {
		s.logger.Error("error while deleting a student", l.Error(err))
		return nil, err
	}
	return res, nil
}

func (s *StudentService) ListStudents(ctx context.Context, req *pbs.ListStudentsReq) (*pbs.ListStudentsRes, error) {
	res, err := s.storage.Student().ListStudents(req)
	if err != nil {
		s.logger.Error("error while getting a list of students", l.Error(err))
		return nil, err
	}
	return res, nil
}
