# DisGram
Discord & Instagram Benchmarked Gaming Web SNS with Microservice Architecture

## Workflow

### 0. 서비스 설계, (요구사항 분석)

---

### 1. 데이터 모델링, ERD 초안

![img](ERD.png)

---

### 2. API 명세

    https://www.notion.so/Instagram-Clone-API-2b02b42602558028a96ad8f99a7fa4ec?source=copy_link

---

### 3. 아키텍처 디자인

![img](architecture.png)

---

### 4. 스키마 구체화

#### 4-1. 물리 스키마(Spring, MySQL)

##### 4-1-1. Init

ERD에 따라 DB, 테이블 생성(IF NOT EXISTS)

##### 4-1-2. Indexes

자주 사용될 것으로 예상되는 쿼리문에 대한 인덱스 설계

#### 4-2. 애플리케이션 스키마
    
#### 4-2-1. JPA Entity (Domain Model)

Common Domain(공통)
- **Base**
    - createdAt: LocalDateTime (@CreatedDate)
    - updatedAt: LocalDateTime (@LastModifiedDate)

User Service
- **User**
    - id: Long, PK, AutoIncrement
    - username: String, Unique, Not Null
    - passwordHash: String, Not Null
    - bio: String, Nullable
    - avatarUrl: String, Nullable
    - relations: OneToMany(Follow), Logical OneToMany(Post), Logical OneToMany(Comment)

- **Follow**
    - follower: FK(User)
    - followee: FK(User)
    - (followee, follower): PK
    - relations: ManyToOne(User: Follower), ManyToOne(User: Followee)
---
Post Service
- **Post**
    - id: Long, PK, AutoIncrement
    - userId: Long, Not Null (Logical FK)
    - caption: String(Lob)
    - relations: OneToMany(PostImage), OneToMany(Like)

- **PostImage**
    - id: Long, PK, AutoIncrement
    - post: FK(Post)
    - fileKey: String(255), Not Null, Unique
    - extension: String(50), Not Null
    - url: String, Not Null
    - seq: Integer

- **Like**
    - post: FK(Post)
    - userId: Long, Not Null (Logical FK)
    - (postId, userId): PK
---
Comment Service
- **Comment**
    - id: Long, PK, AutoIncrement
    - userId: Long, Not Null (Logical FK)
    - postId: Long, Not Null (Logical FK)
    - content: String(Lob), Not Null
---
#### 4-2-2. Request/Response DTO
User Service
- **UserRequest**: username, password, bio, avatarImage
- **UserResponse**: id, username, bio, avatarUrl, createdAt

- **FollowRequest**: followerId, followeeId
- **FollowResponse**: followerId, followeeId, createdAt
---
Post Service
- **PostRequest**: userId, caption, images
- **PostResponse**: id, userId, caption, postImages, createdAt

- **LikeRequest**: postId, userId
- **LikeResponse**: postId, userId, createdAt
---
Commnet Service
- **CommentRequest**: postId, userId, content
- **CommentResponse**: id, postId, userId, content, createdAt

---
#### 4-2-3. Repository Interface

DB 접근 계층 정의

User Service

- UserRepository: JpaRepository<User, Long>
    - findByUsername(String Username)
    - existsByUsername(String Username)
---

Post Service

- PostRepository: JpaRepositroy<Post, Long>
    - findAllByUserId(Long userId)
    - findTop10ByOrderByCreatedAtDesc()
---

Comment Service

- CommentRepository: JpaRepository<Comment, Long>
    - findByPostId(Long postId)

---
#### 4-2-4. Service Layer

비즈니스 로직, 트랜잭션 단위 정의

- UserService
    - 유저 등록
    - 프로필 변경

- PostService
    - 게시물 등록
    - 게시물 수정
    - 게시물 삭제

- CommentService
    - 댓글 작성
    - 댓글 수정
    - 댓글 삭제
---
#### 4-2-5. Controller Layer (API Mapping)

2번 단락 참조

---
#### 4-2-6. Exception / ApiResponse Specification

표준 응답 구조

    {
        "status": HttpStatus,
        "message": "No such user exists."
    }

    class ApiResponse<T> {
        private boolean success;
        private T data;
        private HttpStatus status;
        private String message;
    }

---
#### 4-2-7. Kafka Event

비동기 이벤트 구조. Elasticsearch의 인덱스 등록을 위해서 사용됨

    Topic: post.created
    {
        "postId": 123,
        "userId": 42,
        "caption": "this is post",
        "timestamp": "2025-11-27T11:00:00Z"
    }

    Topic: post.deleted
    {
        "postId": 123,
        "deletedAt": "2025-11-27T11:05:00Z"
    }

---
#### 4-2-8. gRPC Protocol

서비스 간 RPC 인터페이스

    service UserGrpcService {
        rpc GetUserInfo(UserRequest) returns (UserResponse);
    }

    message UserRequest {
        int64 userId = 1;
    }

    message UserResponse {
        int64 id = 1;
        string username = 2;
        string avatarUrl = 3;
    }

---
#### 4-2-9. Grafana Log Query

로그 시각화 쿼리

    # Loki LogQL
    {app="post-service"} |= "ERROR"
    {app="user-service"} |= "KafkaProducer"

---
#### 4-2-10. Grafana PromQuery

메트릭 수집 쿼리

    # PromQL
    sum(rate(http_server_requests_seconds_count[1m])) by (status)
    avg(container_memory_usage_bytes{namespace="instagram"})

---
#### 4-2-11. (Potential) Cache Entity / Mapper / Validation / Authentication

---
### 5. DB 마이그레이션

schema.sql, data.sql, indexes.sql 파일 작성

---

### 6. 코드 작성(Test Driven Development)

도메인 : app.ahmola.photogram

#### 6-1. Unit Test

#### 6-2. Integration Test

#### 6-3. API Mock

---

### 7. CI/CD

Github Actions에서 CI/CD 파이프라인 수행

---

#### 7-1. Push

코드를 Github에 Push

#### 7-2. Build

서비스 앱과 도커 이미지 빌드

#### 7-3. Docker Compose, Shell Script

Docker Compose 작성

#### 7-4. Deploy
    
AWS EC2 혹은 네이버 클라우드 활용 예정

---

### 8. 피드백

#### 8-1. Grafana K6를 활용한 Load Test

#### 8-2. [4-2]의 (Potential)파트 도입 여부 결정
