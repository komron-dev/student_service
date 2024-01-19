package postgres

import (
	"encoding/json"
	"time"

	pbs "student_service/genproto/student_service"

	"github.com/google/uuid"

	"github.com/jmoiron/sqlx"
)

type studentRepo struct {
	db *sqlx.DB
}

func NewStudentRepo(db *sqlx.DB) *studentRepo {
	return &studentRepo{db: db}
}

func (r *studentRepo) CreateStudent(req *pbs.CreateStudentReq) (*pbs.CreateStudentRes, error) {
	res := pbs.CreateStudentRes{}
	phone_numbers, err := json.Marshal(res.PhoneNumbers)
	if err != nil {
		return nil, err
	}
	var a = []byte{}
	query := `INSERT INTO students(
		id,
		first_name,
		last_name,
		username,
		email,
		gender,
		address,
		dateofbirth,
		phone_numbers,
		major,
		password,
		created_at
	)
		VALUES($1,$2,$3,$4,$5,$6,$7,$8,$9,$10,$11,$12)
		RETURNING 
		id,
		first_name,
		last_name,
		username,
		email,
		gender,
		address,
		dateofbirth,
		phone_numbers,
		major,
		created_at`
	time := time.Now().Format(time.RFC3339)
	id := uuid.New().String()
	err = r.db.QueryRow(query,
		id,
		req.FirstName,
		req.LastName,
		req.Username,
		req.Email,
		req.Gender,
		req.Address,
		req.DateOfBirth,
		phone_numbers,
		req.Major,
		req.Password,
		time,
	).Scan(
		&res.Id,
		&res.FirstName,
		&res.LastName,
		&res.Username,
		&res.Email,
		&res.Gender,
		&res.Address,
		&res.DateOfBirth,
		&a,
		&res.Major,
		&res.CreatedAt,
	)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(a, &res.PhoneNumbers)
	if err != nil {
		return nil, err
	}
	return &res, nil
}

func (r *studentRepo) GetStudent(req *pbs.ById) (*pbs.GetStudentRes, error) {
	res := pbs.GetStudentRes{}
	phone_numbers := []byte{}
	query := ` SELECT
			id,
			first_name,
			last_name,
			username,
			email,
			gender,
			address,
			dateofbirth,
			phone_numbers,
			major,
			created_at
		FROM students
		WHERE id = $1 AND deleted_at IS NULL
	`
	row := r.db.QueryRow(query, req.StudentId)
	err := row.Scan(
		&res.Id,
		&res.FirstName,
		&res.LastName,
		&res.Username,
		&res.Email,
		&res.Gender,
		&res.Address,
		&res.DateOfBirth,
		&phone_numbers,
		&res.Major,
		&res.CreatedAt,
	)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(phone_numbers, &res.PhoneNumbers)
	if err != nil {
		return nil, err
	}
	return &res, nil
}

func (r *studentRepo) UpdateStudent(req *pbs.UpdateStudentReq) (*pbs.Success, error) {
	phone_numbers, err := json.Marshal(req.PhoneNumbers)
	if err != nil {
		return nil, err
	}
	query := `UPDATE students SET
		first_name = $1,
		last_name = $2,
		username = $3,
		gender = $4,
		address = $5,
		phone_numbers = $6,
		dateofbirth = $7,
		major = $8,
		updated_at = $9

		WHERE id=$10 AND deleted_at IS NULL
		`
	time := time.Now().Format(time.RFC3339)
	_, err = r.db.Exec(query,
		req.FirstName,
		req.LastName,
		req.Username,
		req.Gender,
		req.Address,
		phone_numbers,
		req.DateOfBirth,
		req.Major,
		time,
		req.Id,
	)
	if err != nil {
		return nil, err
	}

	return &pbs.Success{Message: "ok"}, nil
}

func (r *studentRepo) DeleteStudent(req *pbs.ById) (*pbs.Success, error) {
	query := `UPDATE students SET deleted_at=$1 WHERE id=$2 AND deleted_at IS NULL`
	time := time.Now().Format(time.RFC3339)
	_, err := r.db.Exec(query, time, req.StudentId)
	if err != nil {
		return nil, err
	}
	return &pbs.Success{Message: "ok"}, err
}

func (r *studentRepo) ListStudents(req *pbs.ListStudentsReq) (*pbs.ListStudentsRes, error) {
	var (
		phone_numbers []byte
		count         int64
	)
	offset := (req.Page - 1) * req.Limit
	query := ` SELECT
			id,
			first_name,
			last_name,
			username,
			email,
			gender,
			address,
			dateofbirth,
			phone_numbers,
			major,
			created_at
		FROM students
		WHERE deleted_at IS NULL LIMIT $1 OFFSET $2
	`
	rows, err := r.db.Query(query, req.Limit, offset)
	if err != nil {
		return nil, err
	}
	students := []*pbs.GetStudentRes{}
	for rows.Next() {
		student := &pbs.GetStudentRes{}
		err := rows.Scan(
			&student.Id,
			&student.FirstName,
			&student.LastName,
			&student.Username,
			&student.Email,
			&student.Gender,
			&student.Address,
			&student.DateOfBirth,
			&phone_numbers,
			&student.Major,
			&student.CreatedAt,
		)
		if err != nil {
			return nil, err
		}
		err = json.Unmarshal(phone_numbers, &student.PhoneNumbers)
		if err != nil {
			return nil, err
		}
		students = append(students, student)
	}
	queryForCount := `SELECT COUNT(*) FROM students WHERE deleted_at IS NULL`
	row := r.db.QueryRow(queryForCount)
	err = row.Scan(&count)
	if err != nil {
		return nil, err
	}
	return &pbs.ListStudentsRes{
		Students: students,
		Count:    count,
	}, nil
}
