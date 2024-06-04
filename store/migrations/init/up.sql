CREATE TABLE
    IF NOT EXISTS users (
        id SERIAL PRIMARY KEY,
        username VARCHAR(255) NOT NULL UNIQUE,
        password VARCHAR(255),
        is_deleted BOOL DEFAULT false,
        created_at TIMESTAMP DEFAULT NOW (),
        updated_at TIMESTAMP DEFAULT NOW ()
    );

CREATE TABLE
    IF NOT EXISTS course (
        id SERIAL PRIMARY KEY,
        name VARCHAR(255) NOT NULL UNIQUE,
        pars INTEGER[],
        slope REAL,
        rating REAL,
        is_deleted BOOL DEFAULT false,
        created_at TIMESTAMP DEFAULT NOW (),
        updated_at TIMESTAMP DEFAULT NOW ()
    );

CREATE TABLE
    IF NOT EXISTS round(
        id SERIAL PRIMARY KEY,
        course_id SERIAL NOT NULL,
        user_id SERIAL NOT NULL,
        created_at TIMESTAMP DEFAULT NOW (),
        updated_at TIMESTAMP DEFAULT NOW (),
        FOREIGN KEY (course_id) REFERENCES course(id),
        FOREIGN KEY (user_id) REFERENCES users(id)
    );