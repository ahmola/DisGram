USE `user`;

INSERT INTO `users` (`username`, `email`, `password_hash`, `bio`, `avatar_url`) VALUES
('admin', 'admin@test.com', '$2a$10$vI8A...', 'System Administrator', 'https://api.dicebear.com/7.x/avataaars/svg?seed=admin'),
('user1', 'user1@gmail.com', '$2a$10$7XyZ...', 'Hello world!', 'https://api.dicebear.com/7.x/avataaars/svg?seed=user1'),
('gopher_dev', 'gopher@golang.org', '$2a$10$9AbC...', 'Go developer from Goyang', 'https://api.dicebear.com/7.x/avataaars/svg?seed=gopher'),
('spring_expert', 'spring@java.com', '$2a$10$1DeF...', 'Switching to Go now', 'https://api.dicebear.com/7.x/avataaars/svg?seed=spring'),
('docker_master', 'whale@docker.com', '$2a$10$2GhI...', 'Containerization is my life', 'https://api.dicebear.com/7.x/avataaars/svg?seed=docker'),
('grpc_fan', 'grpc@proto.io', '$2a$10$3JkL...', 'Fastest API ever', 'https://api.dicebear.com/7.x/avataaars/svg?seed=grpc'),
('maria_db', 'maria@db.org', '$2a$10$4MnO...', 'Open source database lover', 'https://api.dicebear.com/7.x/avataaars/svg?seed=maria'),
('gin_gonic', 'gin@gin.com', '$2a$10$5PqR...', 'HTTP framework expert', 'https://api.dicebear.com/7.x/avataaars/svg?seed=gin'),
('vue_user', 'vue@frontend.com', '$2a$10$6StU...', 'Frontend developer', 'https://api.dicebear.com/7.x/avataaars/svg?seed=vue'),
('test_user', 'test@test.com', '$2a$10$0VwX...', 'Just for testing', 'https://api.dicebear.com/7.x/avataaars/svg?seed=test');