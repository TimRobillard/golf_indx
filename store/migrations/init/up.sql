CREATE EXTENSION pg_trgm;
update pg_opclass set opcdefault = true where opcname='gin_trgm_ops';
CREATE TABLE
    IF NOT EXISTS users (
        id SERIAL PRIMARY KEY,
        username VARCHAR(255) NOT NULL UNIQUE,
        password VARCHAR(255),
        profile_pic VARCHAR(255) DEFAULT '/public/images/profile_pic.png',
        is_deleted BOOL DEFAULT false,
        created_at TIMESTAMP DEFAULT NOW (),
        updated_at TIMESTAMP DEFAULT NOW ()
    );

CREATE TABLE
    IF NOT EXISTS course (
        id SERIAL PRIMARY KEY,
        name VARCHAR(255) NOT NULL UNIQUE,
        thumbnail VARCHAR(255),
        front INTEGER[],
        back INTEGER[],
        slope REAL,
        rating REAL,
        is_deleted BOOL DEFAULT false,
        created_at TIMESTAMP DEFAULT NOW (),
        updated_at TIMESTAMP DEFAULT NOW ()
    );

CREATE INDEX course_name ON course USING gin(name);

CREATE TABLE
    IF NOT EXISTS round(
        id SERIAL PRIMARY KEY,
        course_id SERIAL NOT NULL,
        user_id SERIAL NOT NULL,
        front INTEGER[],
        back INTEGER[],
        created_at TIMESTAMP DEFAULT NOW (),
        updated_at TIMESTAMP DEFAULT NOW (),
        FOREIGN KEY (course_id) REFERENCES course(id),
        FOREIGN KEY (user_id) REFERENCES users(id)
    );

