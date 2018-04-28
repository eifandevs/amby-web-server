FROM php:7-fpm
RUN apt-get update

# vim
RUN apt-get install -y vim

# mysqlとの接続ツール
RUN docker-php-ext-install pdo_mysql mysqli mbstring

LABEL maintainer "eifandevs <eifan.devs@gmail.com>"