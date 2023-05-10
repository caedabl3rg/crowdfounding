CREATE TABLE IF NOT EXISTS users (
    id BIGINT NOT NULL AUTO_INCREMENT,
    name VARCHAR(255),
    occupation VARCHAR(255),
    password_hash VARCHAR(255),
    avatar_file_name VARCHAR(255),
    role VARCHAR(255),
    token VARCHAR(255), 
    email VARCHAR(255),
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (id)
)ENGINE=InnoDB;