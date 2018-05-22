#!/bin/sh

# dbコンテナの起動を待つ
until mysqladmin ping -hdb --silent; do
    echo 'waiting for mysqld to be connectable...'
    sleep 3
done

curl nginx/api/migrate -X PUT
