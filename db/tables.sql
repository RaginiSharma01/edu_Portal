CREATE TABLE users (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    first_name VARCHAR(100) NOT NULL,
    last_name VARCHAR(100) NOT NULL,
    email VARCHAR(150) UNIQUE NOT NULL,
    phone VARCHAR(20),
    age INT,
    date_of_birth DATE,
    address TEXT,
    password TEXT NOT NULL,
    role VARCHAR(20) CHECK (role IN ('teacher', 'student')) NOT NULL,
    email_verified BOOLEAN DEFAULT FALSE,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
CREATE TABLE teachers (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    user_id UUID UNIQUE REFERENCES users(id) ON DELETE CASCADE,
    qualification VARCHAR(200),
    subjects_teaching TEXT,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE students (
    id UUID DEFAULT gen_random_uuid() PRIMARY KEY,
    user_id UUID REFERENCES users(id) ON DELETE CASCADE,
    father_name VARCHAR(100),
    mother_name VARCHAR(100),
    guardian_name VARCHAR(100),
    occupation VARCHAR(100),
    height INT,
    weight INT,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE EXTENSION IF NOT EXISTS "pgcrypto"; 
-- // beacuse we r using gen_random_uuid()


ALTER TABLE users
DROP CONSTRAINT users_role_check;

ALTER TABLE users
ADD CONSTRAINT users_role_check
CHECK (role IN ('admin','teacher','student'));



INSERT INTO users
(first_name, last_name, email, phone, password, role, email_verified)
VALUES
(
 'Admin',
 'System',
 'admin@school.com',
 '9999999999',
 '$2a$10$xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx',
 'admin',
 true
);

SELECT id FROM subjects WHERE LOWER(name)=LOWER('english');
CREATE TABLE events (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    title VARCHAR(200) NOT NULL,
    event_date DATE NOT NULL,
    venue VARCHAR(200),
    description TEXT,
    type VARCHAR(50),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE teacher_salaries (
    id UUID PRIMARY KEY,
    teacher_id UUID REFERENCES teachers(user_id) ON DELETE CASCADE,
    base_salary INT NOT NULL,
    allowance INT NOT NULL,
    effective_from DATE NOT NULL,
    status VARCHAR(20) DEFAULT 'pending',
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);