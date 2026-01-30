USE `comment`;

-- Create Comments Table 
CREATE TABLE IF NOT EXISTS `comments` (
    `id`          BIGINT UNSIGNED AUTO_INCREMENT PRIMARY KEY,
    `user_id`     BIGINT UNSIGNED NOT NULL COMMENT '작성자 ID (User 서비스 참조)',
    `post_id`     BIGINT UNSIGNED NOT NULL COMMENT '게시글 ID (Post 서비스 참조)',
    `content`     TEXT NOT NULL,
    `created_at`  DATETIME(3) NOT NULL DEFAULT CURRENT_TIMESTAMP(3),
    `updated_at`  DATETIME(3) NOT NULL DEFAULT CURRENT_TIMESTAMP(3) ON UPDATE CURRENT_TIMESTAMP(3),

    -- Configure Indexes
    -- Optimized performance when reading comments in a particular post in the latest order
    INDEX `idx_post_id_created_at` (`post_id`, `created_at`),
    -- Optimized performance when reading a list of comments written by a specific user
    INDEX `idx_user_id` (`user_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;