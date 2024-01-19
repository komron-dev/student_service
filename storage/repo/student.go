package repo

import (
	pbs "student_service/genproto/student_service"
)

type StudentStorageI interface {
	CreateStudent(*pbs.CreateStudentReq) (*pbs.CreateStudentRes, error)
	GetStudent(*pbs.ById) (*pbs.GetStudentRes, error)
	UpdateStudent(*pbs.UpdateStudentReq) (*pbs.Success, error)
	DeleteStudent(*pbs.ById) (*pbs.Success, error)
	ListStudents(*pbs.ListStudentsReq) (*pbs.ListStudentsRes, error)
}
