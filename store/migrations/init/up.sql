CREATE EXTENSION pg_trgm;

update pg_opclass
set
    opcdefault = true
where
    opcname = 'gin_trgm_ops';

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

CREATE INDEX course_name ON course USING gin (name);

CREATE TABLE
    IF NOT EXISTS round(
        id SERIAL PRIMARY KEY,
        course_id SERIAL NOT NULL,
        user_id SERIAL NOT NULL,
        front INTEGER[],
        back INTEGER[],
        date TIMESTAMP,
        created_at TIMESTAMP DEFAULT NOW (),
        updated_at TIMESTAMP DEFAULT NOW (),
        FOREIGN KEY (course_id) REFERENCES course (id) ON DELETE CASCADE,
        FOREIGN KEY (user_id) REFERENCES users (id) ON DELETE CASCADE
    );

CREATE TABLE IF NOT EXISTS handicap(
    id SERIAL PRIMARY KEY,
    user_id INTEGER,
    indx DECIMAL,
    date TIMESTAMP,
    created_at TIMESTAMP DEFAULT NOW (),
    updated_at TIMESTAMP DEFAULT NOW (),
    FOREIGN KEY (user_id) REFERENCES users (id) ON DELETE CASCADE
);

BEGIN TRANSACTION;

INSERT INTO
    users (username, password)
VALUES
    (
        'albatrooss',
        '$2a$14$FcxstGl6bVJw.GQTTkRdD.Yy5tWJSvrtuXZ7ifcd3dyIZ9lBFjt4u'
    );

INSERT INTO
    course (name, thumbnail, rating, slope, front, back)
VALUES
    (
        'manderley central north',
        'https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcRYsa5s9fz-agjYOZtBTJDSaDV_78gxOiRTQw&usqp=CAU',
        67.1,
        110,
        '{5, 3, 5, 3, 4, 3, 5, 4, 4}',
        '{5, 4, 4, 3, 5, 4, 3, 4, 4}'
    ),
    (
        'manderley north south',
        'https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcRYsa5s9fz-agjYOZtBTJDSaDV_78gxOiRTQw&usqp=CAU',
        65.3,
        112,
        '{5, 4, 4, 3, 5, 4, 3, 4, 4}',
        '{4, 4, 3, 4, 4, 3, 5, 3, 5}'
    ),
    (
        'dragonfly golf links',
        'https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcRYsa5s9fz-agjYOZtBTJDSaDV_78gxOiRTQw&usqp=CAU',
        69.9,
        123,
        '{4, 4, 4, 4, 3, 5, 3, 4, 5}',
        '{4, 5, 4, 3, 4, 5, 4, 4, 3}'
    ),
    (
        'cedarhill',
        'https://stittsvillecentral.ca/wp-content/uploads/amberwood-village-golf-green.jpg',
        67.6,
        112,
        '{4, 4, 5, 4, 3, 4, 4, 4, 3}',
        '{4, 4, 4, 4, 3, 3, 4, 5, 4}'
    ),
    (
        'amberwood',
        'https://stittsvillecentral.ca/wp-content/uploads/amberwood-village-golf-green.jpg',
        31.2,
        99,
        '{3, 4, 4, 3, 4, 3, 3, 4, 4}',
        '{}'
    );

INSERT INTO
    round(user_id, course_id, front, back, date)
VALUES
    (
        1,
        1,
        '{8, 4, 6, 5, 6, 4, 9, 6, 6}',
        '{7, 6, 6, 2, 7, 4, 6, 4, 6}',
        '2024-06-30'
    ),
    (
        1,
        3,
        '{5, 5, 8, 5, 2, 5, 4, 5, 5}',
        '{5, 5, 5, 4, 4, 8, 6, 6, 3}',
        '2024-06-24'
    ),
    (
        1,
        4,
        '{7, 5, 5, 8, 3, 4, 6, 7, 5}',
        '{3, 7, 5, 3, 3, 4, 7, 6, 8}',
        '2024-06-24'
    ),
    (
        1,
        5,
        '{3, 5, 3, 4, 6, 4, 3, 6, 6}',
        '{3, 5, 3, 4, 6, 4, 3, 6, 6}',
        '2024-06-24'
    ),
    (
        1,
        2,
        '{7, 7, 6, 4, 6, 8, 4, 4, 6}',
        '{5, 6, 4, 7, 5, 5, 6, 3, 7}',
        '2024-06-24'
    );

COMMIT