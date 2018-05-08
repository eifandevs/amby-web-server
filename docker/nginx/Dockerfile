FROM nginx:latest

ARG node_version="8.x"

# vim
RUN apt-get update && apt-get install -y vim

# procps
RUN apt-get update && apt-get install -y procps

# ping
RUN apt-get update && apt-get install -y iputils-ping net-tools

# git
RUN apt-get update && apt-get install -y git

# curl
RUN apt-get update && apt-get install -y curl

# node(8.x) 
RUN apt-get update && apt-get install -my wget gnupg \
    && curl -sL https://deb.nodesource.com/setup_${node_version} | bash \
    && apt-get update && apt-get install -y nodejs \
    && apt-get update && apt-get install -y build-essential

WORKDIR /var/www

# git clone(shallow clone)
RUN git clone --depth=1 https://github.com/eifandevs/qass-news-client.git \
    && git clone --depth=1 https://github.com/eifandevs/qass-news-server.git

# client setup
WORKDIR /var/www/qass-news-client
RUN npm install && npm run build
RUN chmod 777 -R /var/www/qass-news-client/dist

# server setup
WORKDIR /var/www/qass-news-server
RUN chmod 777 -R /var/www/qass-news-server/storage \
    && chmod 777 -R /var/www/qass-news-server/public

LABEL maintainer "eifandevs <eifan.devs@gmail.com>"