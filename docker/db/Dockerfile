FROM mariadb

# vim
RUN apt-get update && apt-get install -y vim

# procps
RUN apt-get update && apt-get install -y procps

# my.cnf
ADD my.cnf /etc/mysql/conf.d

LABEL maintainer "eifandevs <eifan.devs@gmail.com>"