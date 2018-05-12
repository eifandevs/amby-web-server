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

# node(8.x) 
RUN apt-get update && apt-get install -my wget gnupg \
    && curl -sL https://deb.nodesource.com/setup_${node_version} | bash \
    && apt-get update && apt-get install -y nodejs build-essential

WORKDIR /var/www

# git clone(shallow clone)
RUN git clone --depth=1 https://github.com/eifandevs/qass-news-client.git \
    && git clone --depth=1 https://github.com/eifandevs/qass-news-server.git

# client setup
WORKDIR /var/www/qass-news-client
RUN npm install \
  && npm run build \
  && chmod 777 -R /var/www/qass-news-client/dist

# server setup
WORKDIR /var/www/qass-news-server
RUN chmod 777 -R /var/www/qass-news-server/storage \
    && chmod 777 -R /var/www/qass-news-server/public

CMD bash -c "/var/www/qass-news-client/up.sh && /var/www/qass-news-server/up.sh"

LABEL maintainer "eifandevs <eifan.devs@gmail.com>"