USE `user`;

-- User Table
CREATE TABLE IF NOT EXISTS `users` (
    `id`              BIGINT UNSIGNED AUTO_INCREMENT PRIMARY KEY,
    `username`        VARCHAR(100) NOT NULL,
    `email`           VARCHAR(191) NOT NULL, -- Conisder the Restriction of index length
    `password_hash`   VARCHAR(255) NOT NULL,
    `bio`             TEXT,
    `avatar_url`      VARCHAR(255),
    `created_at`      DATETIME(3) NOT NULL DEFAULT CURRENT_TIMESTAMP(3),
    `updated_at`      DATETIME(3) NOT NULL DEFAULT CURRENT_TIMESTAMP(3) ON UPDATE CURRENT_TIMESTAMP(3),

    -- unique tag of GORM entity
    CONSTRAINT `uni_users_username` UNIQUE (`username`),
    CONSTRAINT `uni_users_email` UNIQUE (`email`),
    CONSTRAINT `uni_users_avatar_url` UNIQUE (`avatar_url`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- Index (Optimize Search performance)
-- Complex index of username and email
CREATE INDEX `idx_users_username_email` ON `users` (`username`, `email`);