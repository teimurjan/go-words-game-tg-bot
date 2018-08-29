
-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied
CREATE TABLE `words` (
    `id` INTEGER NOT NULL AUTO_INCREMENT PRIMARY KEY,
    `created` DATETIME NOT NULL,
    `modified` DATETIME NOT NULL,
    `value` VARCHAR(255) NOT NULL,
    KEY (`value`)
);

-- +goose Down
-- SQL section 'Down' is executed when this migration is rolled back
DROP TABLE `words`;
