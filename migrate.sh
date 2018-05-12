#!/bin/sh

# dbコンテナの起動を待つ
echo "Waiting for mysql..."
until mysql -hdb -uroot -proot -P3306 &> /dev/null
do
        >$2 echo -n "."
        sleep 1
done

php artisan migrate