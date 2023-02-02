-- +goose Up
CREATE TABLE IF NOT EXISTS `release_flag` (
--     id binary(16) PRIMARY KEY,
--     sub varchar(36) generated always as (
--         insert(
--             insert(
--                 insert(
--                     insert( hex(id),9,0,'-' ),
--                     14,0,'-'),
--                 19,0,'-'),
--             24,0,'-')
--         ) virtual,
    id INT PRIMARY KEY
    label VARCHAR(75),
    description VARCHAR(200),
    status BIT,
    start_time DATETIME
    created_at  DATETIME DEFAULT NOW(),
    updated_at  DATETIME DEFAULT NOW(),
    deleted_at  DATETIME
    );
-- +goose Down
DROP TABLE IF EXISTS users
