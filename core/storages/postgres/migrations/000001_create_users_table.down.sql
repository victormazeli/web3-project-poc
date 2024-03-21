CREATE TABLE IF NOT EXISTS users(
    id          UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    firstname   VARCHAR (50) NOT NULL,
    lastname    VARCHAR (50) NOT NULL,
    username    VARCHAR (50) UNIQUE NOT NULL,
    password    VARCHAR (50) NOT NULL,
    dob         TIMESTAMP    NOT NULL,
    email       VARCHAR (300) UNIQUE NOT NULL,
    phone       VARCHAR (300) UNIQUE NOT NULL,
    );