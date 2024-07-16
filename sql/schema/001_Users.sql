-- +goose Up 
-- Enable the uuid-ossp extension
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";
CREATE TABLE users (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    first_name TEXT  NOT NULL ,
    last_name TEXT  NOT NULL ,
    email TEXT UNIQUE NOT NULL ,
    access_token TEXT UNIQUE ,
    password TEXT UNIQUE NOT NULL 
    
);

-- +goose Down
DROP TABLE users;
