package storage

import (
	"student_service/storage/postgres"
	"student_service/storage/repo"

	"github.com/jmoiron/sqlx"
)

type IStorage interface {
	Student() repo.StudentStorageI
}

type storagePg struct {
	db          *sqlx.DB
	studentRepo repo.StudentStorageI
}

func NewStoragePg(db *sqlx.DB) *storagePg {
	return &storagePg{
		db:          db,
		studentRepo: postgres.NewStudentRepo(db),
	}
}
func (s storagePg) Student() repo.StudentStorageI {
	return s.studentRepo
}
