
-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied
CREATE TABLE `games` (
    `id` INTEGER NOT NULL AUTO_INCREMENT PRIMARY KEY,
    `created` DATETIME NOT NULL,
    `modified` DATETIME NOT NULL,
    `is_user_turn` BOOLEAN DEFAULT NULL,
    `last_bot_word_id` INTEGER DEFAULT NULL,
    `user_id` INTEGER NOT NULL,
    FOREIGN KEY (`user_id`) REFERENCES users(`id`)
);

-- +goose Down
-- SQL section 'Down' is executed when this migration is rolled back
ALTER TABLE `games`
DROP FOREIGN KEY `user_id`;
DROP TABLE `games`;
