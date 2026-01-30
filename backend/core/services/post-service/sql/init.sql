USE `post`;

-- Posts Table
CREATE TABLE IF NOT EXISTS `posts` (
    `id`          BIGINT UNSIGNED AUTO_INCREMENT PRIMARY KEY,
    `user_id`     BIGINT UNSIGNED NOT NULL,
    `caption`     TEXT,
    `created_at`  DATETIME(3) NOT NULL DEFAULT CURRENT_TIMESTAMP(3),
    `updated_at`  DATETIME(3) NOT NULL DEFAULT CURRENT_TIMESTAMP(3) ON UPDATE CURRENT_TIMESTAMP(3),
    
    INDEX `idx_posts_user_id` (`user_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- PostImages Table
CREATE TABLE IF NOT EXISTS `post_images` (
    `id`          BIGINT UNSIGNED AUTO_INCREMENT PRIMARY KEY,
    `post_id`     BIGINT UNSIGNED NOT NULL,
    `file_key`    VARCHAR(255) NOT NULL,
    `extension`   VARCHAR(10) NOT NULL,
    `url`         VARCHAR(500) NOT NULL,
    `seq`         BIGINT UNSIGNED DEFAULT 0 COMMENT '이미지 출력 순서',
    `created_at`  DATETIME(3) NOT NULL DEFAULT CURRENT_TIMESTAMP(3),
    `updated_at`  DATETIME(3) NOT NULL DEFAULT CURRENT_TIMESTAMP(3) ON UPDATE CURRENT_TIMESTAMP(3),

    CONSTRAINT `uni_post_images_file_key` UNIQUE (`file_key`),
    -- Deleted when the post is deleted
    CONSTRAINT `fk_post_images_post` FOREIGN KEY (`post_id`) REFERENCES `posts`(`id`) ON DELETE CASCADE,
    INDEX `idx_post_id_seq` (`post_id`, `seq`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- Likes Table
CREATE TABLE IF NOT EXISTS `likes` (
    `id`          BIGINT UNSIGNED AUTO_INCREMENT PRIMARY KEY,
    `post_id`     BIGINT UNSIGNED NOT NULL,
    `user_id`     BIGINT UNSIGNED NOT NULL,
    `created_at`  DATETIME(3) NOT NULL DEFAULT CURRENT_TIMESTAMP(3),
    `updated_at`  DATETIME(3) NOT NULL DEFAULT CURRENT_TIMESTAMP(3) ON UPDATE CURRENT_TIMESTAMP(3),

    -- One User, One Like at the post
    UNIQUE KEY `uni_post_user_like` (`post_id`, `user_id`),
    -- Deleted when post is deleted
    CONSTRAINT `fk_likes_post` FOREIGN KEY (`post_id`) REFERENCES `posts`(`id`) ON DELETE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;