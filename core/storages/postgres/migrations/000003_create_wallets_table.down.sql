CREATE TABLE IF NOT EXISTS wallets(
    id          UUID         PRIMARY KEY DEFAULT uuid_generate_v4(),
    name        VARCHAR (50) NOT NULL,
    balance     VARCHAR (50) NOT NULL,
    private_key VARCHAR (50) NOT NULL,
    public_key  VARCHAR (50) NOT NULL,
    user_id     UUID         NOT 
    FOREIGN KEY (user_id) REFERENCES users (id),
)