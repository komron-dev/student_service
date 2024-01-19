ALTER TABLE students
RENAME COLUMN adress TO address;

ALTER TABLE students
ALTER COLUMN address TYPE TEXT USING address::TEXT;