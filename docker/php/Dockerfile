FROM php:7-fpm

ARG node_version="8.x"

# 環境変数
ENV COMPOSER_ALLOW_SUPERUSER 1
ENV COMPOSER_HOME /composer
ENV PATH $PATH:/composer/vendor/bin

# git
RUN apt-get update && apt-get install -y git

# vim
RUN apt-get update && apt-get install -y vim

# procps
RUN apt-get update && apt-get install -y procps

# ping
RUN apt-get update && apt-get install -y iputils-ping net-tools

# mysql-client
RUN apt-get update && apt-get -y install mysql-client
# zip
RUN apt-get update \
    && apt-get install -y zlib1g-dev \
    && docker-php-ext-install zip

# png
RUN apt-get update && apt-get install -y libpng-dev

# composer
RUN php -r "copy('https://getcomposer.org/installer', 'composer-setup.php');" \
    && php -r "if (hash_file('SHA384', 'composer-setup.php') === '544e09ee996cdf60ece3804abc52599c22b1f40f4323403c44d44fdfdd586475ca9813a858088ffbc1f233e9b180f061') { echo 'Installer verified'; } else { echo 'Installer corrupt'; unlink('composer-setup.php'); } echo PHP_EOL;" \
    && php composer-setup.php \
    && php -r "unlink('composer-setup.php');" \
    && mv composer.phar /usr/local/bin/composer

# node(8.x) 
RUN apt-get update && apt-get install -my wget gnupg \
    && curl -sL https://deb.nodesource.com/setup_${node_version} | bash \
    && apt-get update && apt-get install -y nodejs \
    && apt-get update && apt-get install -y build-essential

# laravel
RUN composer global require "laravel/installer"

# dbとの接続ツール
RUN apt-get update \
    && apt-get install -y libpq-dev \
    && docker-php-ext-install pdo_mysql pdo_pgsql mysqli mbstring

WORKDIR /var/www

# git clone(shallow clone)
RUN git clone --depth=1 https://github.com/eifandevs/qass-news-client.git \
    && git clone --depth=1 https://github.com/eifandevs/qass-news-server.git

# env
ADD .env /var/www/qass-news-server

# client setup
WORKDIR /var/www/qass-news-client
RUN npm install && npm run build

# server setup
WORKDIR /var/www/qass-news-server
RUN composer install && npm install
# RUN php artisan migrate

LABEL maintainer "eifandevs <eifan.devs@gmail.com>"