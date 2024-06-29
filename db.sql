CREATE TABLE customer
(
    id   varchar(100) NOT NULL,
    name varchar(100) NOT NULL,
    PRIMARY KEY (id)
) ENGINE = InnoDB;

ALTER TABLE customer
    ADD COLUMN email      VARCHAR(100),
    ADD COLUMN balance    INT       DEFAULT 0,
    ADD COLUMN rating     DOUBLE    DEFAULT 0.0,
    ADD COLUMN created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    ADD COLUMN birth_date DATE,
    ADD COLUMN married    BOOLEAN   DEFAULT FALSE;

CREATE TABLE user (
    username VARCHAR(100) NOT NULL,
    password VARCHAR(100) NOT NULL,
    PRIMARY KEY (username)
) ENGINE = InnoDB;

CREATE TABLE comments (
    id INT NOT NULL AUTO_INCREMENT,
    email VARCHAR(100) NOT NULL,
    comment TEXT,
    PRIMARY KEY (id)
) ENGINE = InnoDB;
