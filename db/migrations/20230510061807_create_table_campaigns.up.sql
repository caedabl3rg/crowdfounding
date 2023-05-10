CREATE TABLE IF NOT EXISTS campaigns (
    id BIGINT NOT NULL AUTO_INCREMENT,
    user_id BIGINT,
    name VARCHAR(255),
    short_description VARCHAR(255),
    description TEXT,
    goal_amount INT,
    current_amount INT,
    perks TEXT,
    backer_count INT,  
    slug VARCHAR(255),
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (id)
)ENGINE=InnoDB;