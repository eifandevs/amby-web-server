FROM nginx:latest

# vim
RUN apt-get update && apt-get install -y vim

# procps
RUN apt-get update && apt-get install -y procps

# ping
RUN apt-get update && apt-get install -y iputils-ping net-tools

# default.conf
ADD default.conf /etc/nginx/conf.d/

LABEL maintainer "eifandevs <eifan.devs@gmail.com>"