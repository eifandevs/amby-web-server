#!/bin/sh

cd `dirname $0`

# ログファイル
cd storage/logs
touch laravel.log
chmod 777 laravel.log

