CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE IF NOT EXISTS users(
    id          UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    firstname   VARCHAR (255) NOT NULL,
    lastname    VARCHAR (255) NOT NULL,
    middlename    VARCHAR (255) NOT NULL,
    username    VARCHAR (255) UNIQUE NOT NULL,
    verified    BOOLEAN      NOT NULL DEFAULT false,
    address_type VARCHAR (255) NOT NULL,
    zip_code    VARCHAR (255) NOT NULL,
    street      VARCHAR (255) NOT NULL,
    city        VARCHAR (255) NOT NULL,
    state       VARCHAR (255) NOT NULL,
    country     VARCHAR (255) NOT NULL,
    password    VARCHAR (55) NOT NULL,
    dob         TIMESTAMP    NOT NULL,
    email       VARCHAR (255) UNIQUE NOT NULL,
    phone       VARCHAR (255) UNIQUE NOT NULL,
    created_at  TIMESTAMPTZ NOT NULL DEFAULT (now()),
    updated_at  TIMESTAMPTZ NOT NULL DEFAULT (now())
);