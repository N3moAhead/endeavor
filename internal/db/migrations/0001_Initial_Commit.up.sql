CREATE TABLE Days (
    id SERIAL PRIMARY KEY,
    date DATE NOT NULL UNIQUE DEFAULT CURRENT_DATE
);

CREATE TABLE Categories (
    id SERIAL PRIMARY KEY,
    name VARCHAR(64),
    description VARCHAR(64),
    color VARCHAR(32)
);

CREATE TABLE Activities (
    id SERIAL PRIMARY KEY,
    name VARCHAR(64),
    category_id INT REFERENCES Categories(id)
);

CREATE TABLE Moods (
    id SERIAL PRIMARY KEY,
    name VARCHAR(64),
    description VARCHAR(64),
    day_id INT REFERENCES Days(id) UNIQUE
);

CREATE TABLE Day_Activities (
    id SERIAL PRIMARY KEY,
    day_id INT REFERENCES Days(id),
    activity_id INT REFERENCES Activities(id)
);
