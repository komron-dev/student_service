CREATE TABLE IF NOT EXISTS students (
    id UUID NOT NULL,
    first_name varchar(100) NOT NULL,
    last_name varchar(100) NOT NULL,
    username varchar(100),
    email varchar(100) NOT NULL,
    gender varchar(6) CHECK(gender IN('male','female')) NOT NULL,
    dateOfBirth TIMESTAMP NOT NULL,
    major varchar(100) NOT NULL,
    GPA REAL,
    adress JSONB NOT NULL,
    phone_numbers JSONB NOT NULL,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP,
    deleted_at TIMESTAMP
);

CREATE TABLE IF NOT EXISTS student_tokens (
    access_token TEXT,
    refresh_token TEXT
);