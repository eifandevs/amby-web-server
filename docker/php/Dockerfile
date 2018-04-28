FROM php:7-fpm


# 環境変数
ENV COMPOSER_ALLOW_SUPERUSER 1

# インストール準備
RUN apt-get update

# vim
RUN apt-get install -y vim

# ZIP
RUN apt-get install -y zlib1g-dev
RUN docker-php-ext-install zip

# composer
RUN php -r "copy('https://getcomposer.org/installer', 'composer-setup.php');"
RUN php -r "if (hash_file('SHA384', 'composer-setup.php') === '544e09ee996cdf60ece3804abc52599c22b1f40f4323403c44d44fdfdd586475ca9813a858088ffbc1f233e9b180f061') { echo 'Installer verified'; } else { echo 'Installer corrupt'; unlink('composer-setup.php'); } echo PHP_EOL;"
RUN php composer-setup.php
RUN php -r "unlink('composer-setup.php');"
RUN mv composer.phar /usr/local/bin/composer

# mysqlとの接続ツール
RUN docker-php-ext-install pdo_mysql mysqli mbstring

LABEL maintainer "eifandevs <eifan.devs@gmail.com>"