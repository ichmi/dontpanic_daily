#!/bin/bash

docker compose down
docker rmi daily-back:latest daily-front:latest daily-db:latest
docker compose up

# MANUAL SCRIPT -----------------------------------------------
# up.sh:
# docker build client/ -t front
# docker container run -dp 8080:8080 --name front front

# docker build server/ -t back
# docker container run -dp 8081:8081 --name back back

# docker build db/ -t db
# docker run -dp 5432:5432 --name db db -d postgres
# docker exec -it db ./script.sh


# down.sh:
# docker stop db front back
# docker rm db front back
# docker rmi db:latest front:latest back:latest
