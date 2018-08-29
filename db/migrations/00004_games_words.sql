
-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied
CREATE TABLE `words_games` (
    `word_id` INTEGER NOT NULL,
    `game_id` INTEGER NOT NULL,
    FOREIGN KEY (`word_id`) REFERENCES words(`id`) ON DELETE RESTRICT ON UPDATE CASCADE,
    FOREIGN KEY (`game_id`) REFERENCES games(`id`) ON DELETE RESTRICT ON UPDATE CASCADE,
    PRIMARY KEY (`game_id`, `word_id`)
);

-- +goose Down
-- SQL section 'Down' is executed when this migration is rolled back
ALTER TABLE `words_games`
DROP FOREIGN KEY `word_id`;
DROP FOREIGN KEY `game_id`;
DROP TABLE `words_games`;
