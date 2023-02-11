CREATE TABLE IF NOT EXISTS users
(
    user_ID INT GENERATED ALWAYS AS IDENTITY,
    login VARCHAR(255) NOT NULL UNIQUE,
    password VARCHAR(255) NOT NULL,
    PRIMARY KEY(user_ID)
);