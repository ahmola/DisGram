docker compose -f docker-compose-services-dev.yml down -v

docker compose build --no-cache

docker compose -f docker-compose-services-dev.yml up -d