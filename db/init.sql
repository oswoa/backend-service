-- 削除用
-- ALTER TABLE user_management.users DROP INDEX email;
-- DROP TABLE user_management.users;
-- DROP DATABASE user_management;

-- DBの作成
CREATE DATABASE user_management;

-- テーブルの作成
CREATE TABLE user_management.users (
        user_id VARCHAR(26) KEY,
        email VARCHAR(200) NOT NULL UNIQUE,
        is_available boolean NOT NULL,
        created_at DATETIME NOT NULL,
        updated_at DATETIME NOT NULL
    );

-- データの投入
INSERT INTO user_management.users
    (user_id, email, is_available, created_at, updated_at)
VALUES
    ("01JAN8S98KJDKCYM8D27SZXSQY", "ichiro@example.com", 1, "2024-01-01 09:00:00", "2024-01-01 09:00:00"),
    ("01JAN8S98KVGBZBN9NAJWB8P0P", "jiro@example.com", 1, "2024-01-02 09:00:00", "2024-01-02 09:00:00"),
    ("01JAN8S98KPN1Q5CH0JNGF6XKH", "saburo@example.com", 1, "2024-01-03 09:00:00", "2024-01-03 09:00:00"),
    ("01JAN8S98KMRRKC0G4P6DNAG00", "shiro@example.com", 1, "2024-01-04 09:00:00", "2024-01-04 09:00:00"),
    ("01JAN8S98K9EXZ979SD1440KSA", "goro@example.com", 1, "2024-01-05 09:00:00", "2024-01-05 09:00:00")
;

-- インデックスの作成
ALTER TABLE user_management.users ADD INDEX email(email);
