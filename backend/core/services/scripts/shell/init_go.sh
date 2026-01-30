# 1. Go Init
go mod init .

# 2. Gin 
go get -u github.com/gin-gonic/gin

# 3. GORM(DB ORM)/MySQL Driver
go get -u gorm.io/gorm
go get -u gorm.io/driver/mysql

# 4. gRPC
go get -u google.golang.org/grpc
go get github.com/grpc-ecosystem/go-grpc-middleware

# 5. OpenAPI(Swagger)
go get -u github.com/swaggo/gin-swagger
go get -u github.com/swaggo/files