USE `follow`;

INSERT INTO `follows` (`followee_id`, `follower_id`) VALUES
(1, 2),
(1, 3), -- 3번 유저가 1번을 팔로우
(1, 4),
(2, 1), -- 맞팔로우 상황
(3, 2),
(4, 3),
(5, 6),
(6, 7),
(7, 8),
(8, 9);