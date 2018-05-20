#!/bin/sh

cd `dirname $0`

# dbコンテナの起動を待つ
echo "Waiting for mysql..."
until mysql -hdb -uroot -proot -P3306 &> /dev/null
do
        echo -n "."
        sleep 1
done

echo "MySQL is up - executing command"
php artisan migrate
