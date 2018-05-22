FROM debian:jessie

# vim
RUN apt-get update && apt-get install -y \
  mysql-client \
  vim \
  procps \
  iputils-ping \
  net-tools \
  locales \
  curl

# default locale to ja_JP.UTF-8
RUN locale-gen ja_JP.UTF-8
ENV LANG ja_JP.UTF-8
ENV LC_CTYPE ja_JP.UTF-8
RUN localedef -f UTF-8 -i ja_JP ja_JP.utf8

# timezone
ENV TZ Asia/Tokyo

# DB構築
COPY entrypoint.sh /root
RUN chmod +x /root/entrypoint.sh
ENTRYPOINT /root/entrypoint.sh

LABEL maintainer "eifandevs <eifan.devs@gmail.com>"