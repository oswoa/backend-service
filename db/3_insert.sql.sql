-- 削除用
-- SET FOREIGN_KEY_CHECKS = 0;
-- ALTER TABLE user_management.users DROP INDEX email;
-- TRUNCATE TABLE user_management.authorities;
-- TRUNCATE TABLE user_management.users;
-- DROP DATABASE user_management;

-- authorities
INSERT INTO user_management.authorities
    (authority_id, authority_name, can_search, can_create, can_update, can_delete, can_approve, can_pull_back, created_at,            updated_at)
VALUES
    ("01",         "管理者",        1,          1,          1,          1,          1,           1,             "2024-01-05 09:00:00", "2024-01-05 09:00:00"),
    ("02",         "一般",          1,          0,          1,          0,          0,           0,             "2024-01-05 09:00:00", "2024-01-05 09:00:00"),
    ("03",         "承認者",        1,          0,          1,          0,          1,           1,             "2024-01-05 09:00:00", "2024-01-05 09:00:00")
;

-- users
INSERT INTO user_management.users
    (user_id,                      email,               password,                                                           authority_id, is_available, is_deleted, created_at,            updated_at)
VALUES
    ("01JAN8S98KJDKCYM8D27SZXSQY", "emp1@example.com",  "5e884898da28047151d0e56f8dc6292773603d0d6aabbdd62a11ef721d1542d8", "02",         1,            0,          "2024-01-01 09:00:00", "2024-01-01 09:00:00"),
    ("01JB2KY7T4Z2RB11MYH61ANREZ", "emp2@example.com",  "5e884898da28047151d0e56f8dc6292773603d0d6aabbdd62a11ef721d1542d8", "02",         1,            1,          "2024-01-01 09:00:00", "2024-01-01 09:00:00"),
    ("01JAN8S98KPN1Q5CH0JNGF6XKH", "man@example.com",   "5e884898da28047151d0e56f8dc6292773603d0d6aabbdd62a11ef721d1542d8", "03",         1,            0,          "2024-01-03 09:00:00", "2024-01-03 09:00:00"),
    ("01JAN8S98K9EXZ979SD1440KSA", "admin@example.com", "5e884898da28047151d0e56f8dc6292773603d0d6aabbdd62a11ef721d1542d8", "01",         1,            0,          "2024-01-05 09:00:00", "2024-01-05 09:00:00")
;

-- インデックスの作成
ALTER TABLE user_management.users ADD INDEX index_email(email);
