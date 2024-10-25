-- authorities
CREATE TABLE user_management.authorities (
    authority_id   VARCHAR(2)   NOT NULL KEY,
    authority_name VARCHAR(10)  NOT NULL,
    can_search     boolean      NOT NULL,
    can_create     boolean      NOT NULL,
    can_update     boolean      NOT NULL,
    can_delete     boolean      NOT NULL,
    can_approve    boolean      NOT NULL,
    can_pull_back  boolean      NOT NULL,
    created_at     DATETIME     NOT NULL,
    updated_at     DATETIME     NOT NULL
);

-- users
CREATE TABLE user_management.users (
    user_id      VARCHAR(26)  KEY,
    email        VARCHAR(200) NOT NULL UNIQUE,
    password     VARCHAR(64) NOT NULL,
    authority_id VARCHAR(2)   NOT NULL,
    is_available boolean      NOT NULL,
    is_deleted   boolean      NOT NULL,
    created_at   DATETIME     NOT NULL,
    updated_at   DATETIME     NOT NULL,
    FOREIGN KEY (authority_id) REFERENCES user_management.authorities(authority_id)
);
