FROM mysql:8.0.16

ENV MYSQL_ROOT_PASSWORD 123456
COPY ./closetool.sql /docker-entrypoint-initdb.d/closetool.sql

EXPOSE 3306

#ENTRYPOINT ["bash","/mysql/setup.sh"] 