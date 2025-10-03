-- Esegui Migrate psql -U dmnc -d db -f schema.sql

-- Tabella admins
CREATE TABLE IF NOT EXISTS admins (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    email VARCHAR(255) UNIQUE NOT NULL,
    password VARCHAR(255) NOT NULL,
    name VARCHAR(255) NOT NULL
);

CREATE INDEX IF NOT EXISTS admin_email ON admins(email);

-- Tabella users
CREATE TABLE IF NOT EXISTS users (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    email VARCHAR(255) UNIQUE NOT NULL,
    password VARCHAR(255) NOT NULL,
    name VARCHAR(255) NOT NULL,
    surname VARCHAR(255) NOT NULL,
    phone VARCHAR(255),
    role VARCHAR(50) NOT NULL DEFAULT 'STANDARD' CHECK (role IN ('STANDARD', 'EDUCATOR')),
    image_profile VARCHAR(255),
    stripe_customer_id VARCHAR(255) UNIQUE,
    is_active BOOLEAN NOT NULL DEFAULT FALSE,
    last_login TIMESTAMP,
    created_at TIMESTAMP NOT NULL DEFAULT NOW()
);

CREATE INDEX IF NOT EXISTS user_email ON users(email);
CREATE INDEX IF NOT EXISTS user_is_active_created_at ON users(is_active, created_at);
CREATE INDEX IF NOT EXISTS user_name_surname ON users(name, surname);
CREATE INDEX IF NOT EXISTS user_phone ON users(phone);
