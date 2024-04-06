CREATE TABLE IF NOT EXISTS transactions(
    id                  UUID         NOT NULL PRIMARY KEY DEFAULT uuid_generate_v4(),
    hash                VARCHAR (300) NOT NULL,
    amount              VARCHAR (50) NOT NULL,
    sender_address      UUID         NOT NULL,
    recipient_address   UUID         NOT NULL,
    wallet_id           UUID         NOT NULL,
    status              VARCHAR (50) NOT NULL,
    gasLimit            VARCHAR (50) NOT NULL,
    gasPrice            VARCHAR (50) NOT NULL,
    nounce              VARCHAR (255) NOT NULL,
    data                JSONB        NOT NULL DEFAULT '{}',            
    created_at          TIMESTAMPTZ NOT NULL DEFAULT (now()),
    updated_at          TIMESTAMPTZ NOT NULL DEFAULT (now())
)