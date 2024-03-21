CREATE TABLE IF NOT EXISTS transactions(
    id                  UUID         NOT NULL PRIMARY KEY,
    hash                VARCHAR (300) NOT NULL,
    amount              VARCHAR (50) NOT NULL,
    sender_address      UUID         NOT NULL,
    recipient_address   UUID         NOT NULL,
    wallet_id           UUID         NOT NULL,
    status              VARCHAR (50) NOT NULL,
    gasLimit            VARCHAR (50) NOT NULL,
    gasPrice            VARCHAR (50) NOT NULL,
    nounce              VARCHAR (50) NOT NULL,
    data                
    created_at          TIMESTAMP    NOT NULL,
    updated_at          TIMESTAMP    NOT NULL,
)