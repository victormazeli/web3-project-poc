CREATE TABLE IF NOT EXISTS accounts(
    id          UUID         PRIMARY KEY DEFAULT uuid_generate_v4(),
    name        VARCHAR (50) NOT NULL,
    balance     VARCHAR (50) NOT NULL,
    private_key VARCHAR (50) NOT NULL,
    public_key  VARCHAR (50) NOT NULL,
    created_at          TIMESTAMPTZ NOT NULL DEFAULT (now()),
    updated_at          TIMESTAMPTZ NOT NULL DEFAULT (now()),
    user_id     UUID         NOT NULL,
    FOREIGN KEY (user_id) REFERENCES users (id)
)