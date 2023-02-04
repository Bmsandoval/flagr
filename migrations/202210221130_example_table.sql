-- +goose Up
CREATE TABLE IF NOT EXISTS `example` (
    id binary(16) PRIMARY KEY,
    sub varchar(36) generated always as (
        insert(
            insert(
                insert(
                    insert( hex(id),9,0,'-' ),
                    14,0,'-'),
                19,0,'-'),
            24,0,'-')
        ) virtual,
    label VARCHAR(75),
    description VARCHAR(200),
    status BIT,
    start_time  TIMESTAMP,
    created_at  TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at  TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    deleted_at  TIMESTAMP
    );
-- +goose Down
DROP TABLE IF EXISTS `example`
