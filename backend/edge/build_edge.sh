# shut down all the containers
docker rm -f $(docker ps -aq)

# docker compose build
docker compose -f docker-compose-edge.yml -p edges build --no-cache

# docker compose up
docker compose -f docker-compose-edge.yml -p edges up -d