docker compose -f docker-compose-services-dev.yml down -v

go mod tidy

docker compose -f docker-compose-services-dev.yml build --no-cache

docker compose -f docker-compose-services-dev.yml up -d