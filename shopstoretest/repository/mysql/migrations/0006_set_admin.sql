-- +migrate Up
UPDATE `users`
SET
    `role` = 'admin'
WHERE
        `id` = 1;
