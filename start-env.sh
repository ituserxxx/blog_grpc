#!/bin/sh


mkdir -p ./data/mysql/data
mkdir -p ./data/redis/
mkdir -p ./data/etcd/data

docker compose -f docker-compose-env.yml stop
docker compose -f docker-compose-env.yml rm
docker compose -f docker-compose-env.yml up -d







# docker exec -it brl_mysql /bin/sh
# mysql -uroot -p
# use mysql;
# create  USER 'root'@'%' IDENTIFIED BY 'root';
# update user set host = '%' where user = 'root';
# GRANT ALL PRIVILEGES ON *.* TO 'root'@'%';
# FLUSH PRIVILEGES; 