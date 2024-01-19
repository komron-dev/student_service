package postgres

import (
	"fmt"
	"student_service/config"
	pbs "student_service/genproto/student_service"
	"student_service/pkg/db"
	"student_service/storage/repo"
	"testing"

	"github.com/jmoiron/sqlx"
	"github.com/stretchr/testify/suite"
)

type StudentRepoTestSuite struct {
	suite.Suite
	CleanUpFunc func()
	Repository  repo.StudentStorageI
}

func (suite *StudentRepoTestSuite) SetupSuite() {
	pgPool, _, cleanup := db.ConnectToDbAndAlsoForSuite(config.Load())
	suite.Repository = NewStudentRepo(pgPool)
	suite.CleanUpFunc = cleanup
}

// this struct and method below is for testing!
type ForDeleteStudent struct {
	Id string
}

var d *sqlx.DB

func DeleteStudent(req ForDeleteStudent) (*pbs.Success, error) {
	query := `DELETE FROM students WHERE id = $1`
	_, err := d.Exec(query, req.Id)
	if err != nil {
		fmt.Println("ERROR ---->>>", err)
		return nil, err
	}
	return &pbs.Success{Message: "ok"}, nil
}

// testing CreateReview method
func (suite *StudentRepoTestSuite) TestCreateStudent() {
	CreateStudentReq := pbs.CreateStudentReq{
		FirstName:   "fname",
		LastName:    "lname",
		Username:    "username",
		Email:       "example.@gmail.com",
		Gender:      "male",
		DateOfBirth: "2002-01-10",
		Major:       "Economics",
		Address:     "Toshkent",
		Password:    "12345",
	}
	res, err := suite.Repository.CreateStudent(&CreateStudentReq)
	suite.Nil(err)
	suite.NotNil(res, "response must not be nil")
	suite.Equal(CreateStudentReq.Email, res.Email, "both object emails must match")

	// deleting information used for testing

	// DeleteStudent(ForDeleteStudent{Id: res.Id})
}

// testing GetStudent method
func (suite *StudentRepoTestSuite) TestGetStudent() {
	id := &pbs.ById{StudentId: "f0348980-7778-4f90-9862-d7a8680271a8"}
	res, err := suite.Repository.GetStudent(id)
	suite.Nil(err)
	suite.NotNil(res, "response must not be nil")
	suite.Equal(id.StudentId, res.Id)
}

// testing DeleteStudent method
func (suite *StudentRepoTestSuite) TestDeleteStudent() {
	id := &pbs.ById{StudentId: "1d214332-5b4f-4971-bff8-6b3a80ad5755"}
	res, err := suite.Repository.DeleteStudent(id)
	suite.Nil(err)
	suite.NotNil(res, "response must not be nil")

}

// testing ListStudents method
func (suite *StudentRepoTestSuite) ListStudents() {
	req := &pbs.ListStudentsReq{
		Page:  1,
		Limit: 4,
	}
	res, err := suite.Repository.ListStudents(req)
	suite.Nil(err)
	suite.NotNil(res, "response must not be nil")
}

// testing UpdateStudent method
func (suite *StudentRepoTestSuite) TestUpdateStudent() {
	req := &pbs.UpdateStudentReq{
		Id:          "4ec162e8-3c12-4f7a-833f-cbdf3fa34b20",
		FirstName:   "fname",
		LastName:    "lname",
		Username:    "username",
		Gender:      "male",
		DateOfBirth: "2002-01-10",
		Major:       "Economics",
		Address:     "Toshkent",
	}
	res, err := suite.Repository.UpdateStudent(req)
	//deleting information used for testing
	suite.Nil(err)
	suite.NotNil(res, "type of response cannot be nil")
}

func (suite *StudentRepoTestSuite) TearDownSuite() {
	suite.CleanUpFunc()
}

func TestStudentRepositoryTestSuite(t *testing.T) {
	suite.Run(t, new(StudentRepoTestSuite))
}
