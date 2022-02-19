CREATE TABLE users(
    id UUID NOT NULL PRIMARY KEY,
    username VARCHAR(50) NOT NULL UNIQUE,
    full_name VARCHAR(100),
    email VARCHAR(255) UNIQUE,
    password VARCHAR(60) NOT NULL,
    is_identified BOOLEAN NOT NULL DEFAULT FALSE,
    access_token TEXT,
    refresh_token TEXT,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP,
    deleted_at TIMESTAMP
);

CREATE TABLE identifed_user_accounts(
    user_id UUID NOT NULL PRIMARY KEY REFERENCES users(id),
    balance INTEGER NOT NULL DEFAULT 0 CONSTRAINT cash_amount_checker CHECK(balance < 100000),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP
);

CREATE TABLE unidentifed_user_accounts(
    user_id UUID NOT NULL PRIMARY KEY REFERENCES users(id),
    balance INTEGER NOT NULL DEFAULT 0 CONSTRAINT cash_amount_checker CHECK(balance < 10000),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP
);

CREATE TABLE cash_controller(
    user_id UUID NOT NULL REFERENCES users(id),
    income_amount INTEGER,
    expense_amount INTEGER,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
