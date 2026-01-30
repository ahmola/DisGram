USE `post`;

-- Posts 데이터
INSERT INTO `posts` (`id`, `user_id`, `caption`) VALUES
(1, 1, '첫 번째 게시글입니다! #admin'),
(2, 2, '오늘 점심 뭐 먹지?'),
(3, 3, 'Gopher는 귀엽습니다.'),
(4, 1, '라즈베리 파이 서버 구축 중...'),
(5, 5, 'MSA 구조는 어렵지만 재밌네요.');

-- PostImages 데이터 (이미지 10개)
INSERT INTO `post_images` (`post_id`, `file_key`, `extension`, `url`, `seq`) VALUES
(1, 'uuid-key-01', 'jpg', 'https://picsum.photos/id/1/800/600', 0),
(2, 'uuid-key-02', 'png', 'https://picsum.photos/id/10/800/600', 0),
(3, 'uuid-key-03', 'jpg', 'https://picsum.photos/id/20/800/600', 0),
(3, 'uuid-key-04', 'jpg', 'https://picsum.photos/id/21/800/600', 1), -- 3번 게시글의 두 번째 이미지
(4, 'uuid-key-05', 'webp', 'https://picsum.photos/id/30/800/600', 0),
(4, 'uuid-key-06', 'webp', 'https://picsum.photos/id/31/800/600', 1),
(5, 'uuid-key-07', 'jpg', 'https://picsum.photos/id/40/800/600', 0),
(5, 'uuid-key-08', 'jpg', 'https://picsum.photos/id/41/800/600', 1),
(1, 'uuid-key-09', 'png', 'https://picsum.photos/id/50/800/600', 1),
(2, 'uuid-key-10', 'jpg', 'https://picsum.photos/id/60/800/600', 1);

-- Likes 데이터 (10개)
INSERT INTO `likes` (`post_id`, `user_id`) VALUES
(1, 2), (1, 3), (1, 4),
(2, 1), (2, 5),
(3, 1), (3, 2),
(4, 3), (4, 5),
(5, 1);