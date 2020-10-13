CREATE TABLE IF NOT EXISTS users
(
    id         VARCHAR(36) PRIMARY KEY,
    first_name VARCHAR(55)  DEFAULT '',
    last_name  VARCHAR(55)  DEFAULT '',
    full_name  VARCHAR(100) DEFAULT '',
    age        INTEGER      DEFAULT 0
);