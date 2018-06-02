FROM nginx:1.14

ARG node_version="8.x"

# vim
RUN apt-get update && apt-get install -y \
  vim \
  procps \
  iputils-ping \
  net-tools \
  git \
  curl \
  locales \
  openssl

# default locale to ja_JP.UTF-8
RUN locale-gen ja_JP.UTF-8
ENV LANG ja_JP.UTF-8
ENV LC_CTYPE ja_JP.UTF-8
RUN localedef -f UTF-8 -i ja_JP ja_JP.utf8

# timezone
ENV TZ Asia/Tokyo

# node(8.x) 
RUN apt-get update && apt-get install -my wget gnupg \
  && curl -sL https://deb.nodesource.com/setup_${node_version} | bash \
  && apt-get update && apt-get install -y nodejs build-essential

# shallow clone
WORKDIR /var/www
RUN git clone --depth=1 https://github.com/eifandevs/qass-news-server.git

# self-signed certificate
COPY ./keys/server-crt.pem /etc/nginx/server-crt.pem
COPY ./keys/server-privatekey.pem /etc/nginx/server-privatekey.pem

# build(no use cache)
ADD https://www.random.org/strings/?num=10&len=8&digits=on&upperalpha=on&loweralpha=on&unique=on&format=plain&rnd=new uuid

# server setup
WORKDIR /var/www/qass-news-server
RUN git fetch origin \
  && git rebase origin/master \
  && chmod 777 -R /var/www/qass-news-server/storage \
  && chmod 777 -R /var/www/qass-news-server/public

LABEL maintainer "eifandevs <eifan.devs@gmail.com>"