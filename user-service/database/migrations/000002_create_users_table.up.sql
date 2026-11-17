CREATE TABLE IF NOT EXISTS users (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    email VARCHAR(255) UNIQUE NOT NULL,
    password VARCHAR(255) NOT NULL,
    phone VARCHAR(255) NULL,
    photo VARCHAR(255) NULL,
    address TEXT NULL,
    lat VARCHAR(50) NULL,
    lng VARCHAR(50) NULL,
    is_verified BOOLEAN DEFAULT FALSE,
    role_id INT REFERENCES roles(id),
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NULL,
    deleted_at TIMESTAMP NULL
);

create index idx_users_email ON users(email);
create index idx_users_is_verified ON users(is_verified);