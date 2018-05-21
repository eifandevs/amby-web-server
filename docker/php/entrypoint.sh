#!/bin/sh

cd /var/www/qass-news-server

# dbコンテナの起動を待つ
until mysqladmin ping -hdb --silent; do
    echo 'waiting for mysqld to be connectable...'
    sleep 3
done

php artisan migrate

while true
do
    sleep 10
done