FROM nginx:1.14

ARG node_version="8.x"

# vim
RUN apt-get update && apt-get install -y \
  vim \
  procps \
  iputils-ping \
  net-tools \
  git \
  curl

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

# git clone(shallow clone)
RUN git clone --depth=1 https://github.com/eifandevs/qass-news-client.git \
    && git clone --depth=1 https://github.com/eifandevs/qass-news-server.git

# client setup
WORKDIR /var/www/qass-news-client
RUN npm install \
  && npm run build

# server setup
WORKDIR /var/www/qass-news-server

CMD bash -c "/var/www/qass-news-client/up.sh && /var/www/qass-news-server/up.sh"

LABEL maintainer "eifandevs <eifan.devs@gmail.com>"