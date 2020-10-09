FROM mysql:5.7

ENV MYSQL_ROOT_PASSWORD 123456
COPY ./closetool.sql /docker-entrypoint-initdb.d/closetool.sql
#ENTRYPOINT ["bash","/mysql/setup.sh"] 
