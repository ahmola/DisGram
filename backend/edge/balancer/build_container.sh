docker rm -f balancer
docker build -t balancer .
docker run -d --name balancer -p 80:80 balancer
