USE `follow`;

-- 테이블 생성
CREATE TABLE IF NOT EXISTS `follows` (
    `id`          BIGINT UNSIGNED AUTO_INCREMENT PRIMARY KEY,
    `followee_id` BIGINT UNSIGNED NOT NULL COMMENT '팔로우를 받는 사람',
    `follower_id` BIGINT UNSIGNED NOT NULL COMMENT '팔로우를 하는 사람',
    `created_at`  DATETIME(3) NOT NULL DEFAULT CURRENT_TIMESTAMP(3),
    `updated_at`  DATETIME(3) NOT NULL DEFAULT CURRENT_TIMESTAMP(3) ON UPDATE CURRENT_TIMESTAMP(3),

    -- 동일한 팔로우 관계가 중복 생성되지 않도록 복합 유니크 설정
    UNIQUE KEY `uni_followee_follower` (`followee_id`, `follower_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- 인덱스 설정 (조회 성능 최적화)
-- "내가 팔로우하는 사람들(FollowerID)" 조회를 위한 인덱스
CREATE INDEX `idx_follower_id` ON `follows` (`follower_id`);
-- "나를 팔로우하는 사람들(FolloweeID)" 조회를 위한 인덱스
CREATE INDEX `idx_followee_id` ON `follows` (`followee_id`);