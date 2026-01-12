docker rm -f balancer
docker build --no-cache -t balancer .
docker run -d --name balancer -p 80:80 -p 8081:8081 -p 8082:8082 balancer