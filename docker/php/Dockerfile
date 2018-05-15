FROM php:7.2-fpm

ARG node_version="8.x"

# 環境変数
ENV COMPOSER_ALLOW_SUPERUSER 1
ENV COMPOSER_HOME /composer
ENV PATH $PATH:/composer/vendor/bin

# git
RUN apt-get update && apt-get install -y \
  git \
  vim \
  procps \
  iputils-ping \
  net-tools \
  mysql-client \
  zlib1g-dev \
  libpng-dev \
  libpq-dev \
  locales

# default locale to ja_JP.UTF-8
RUN locale-gen ja_JP.UTF-8
ENV LANG ja_JP.UTF-8
ENV LC_CTYPE ja_JP.UTF-8
RUN localedef -f UTF-8 -i ja_JP ja_JP.utf8

# timezone
ENV TZ Asia/Tokyo

# zip db接続
RUN docker-php-ext-install zip pdo_mysql pdo_pgsql mysqli mbstring

# composer
RUN php -r "copy('https://getcomposer.org/installer', 'composer-setup.php');" \
    && php -r "if (hash_file('SHA384', 'composer-setup.php') === '544e09ee996cdf60ece3804abc52599c22b1f40f4323403c44d44fdfdd586475ca9813a858088ffbc1f233e9b180f061') { echo 'Installer verified'; } else { echo 'Installer corrupt'; unlink('composer-setup.php'); } echo PHP_EOL;" \
    && php composer-setup.php \
    && php -r "unlink('composer-setup.php');" \
    && mv composer.phar /usr/local/bin/composer

# node(8.x) 
RUN apt-get update && apt-get install -my wget gnupg \
    && curl -sL https://deb.nodesource.com/setup_${node_version} | bash \
    && apt-get update && apt-get install -y nodejs build-essential

# laravel
RUN composer global require "laravel/installer"

WORKDIR /var/www

# git clone(shallow clone) & setup
RUN git clone --depth=1 https://github.com/eifandevs/qass-news-client.git \
    && git clone --depth=1 https://github.com/eifandevs/qass-news-server.git

# env
ADD .env /var/www/qass-news-server

# build(no use cache)
ADD https://www.random.org/strings/?num=10&len=8&digits=on&upperalpha=on&loweralpha=on&unique=on&format=plain&rnd=new uuid

# client setup
WORKDIR /var/www/qass-news-client
RUN npm install \
  && npm run build \
  && chmod 777 -R /var/www/qass-news-client/dist

# server setup
WORKDIR /var/www/qass-news-server
RUN composer install \
  && npm install \
  && chmod 777 -R /var/www/qass-news-server/storage \
  && chmod 777 -R /var/www/qass-news-server/vendor \
  && chmod 777 -R /var/www/qass-news-server/public
  
LABEL maintainer "eifandevs <eifan.devs@gmail.com>"