#!/bin/sh

pushd /var/www/qass-news-server

# ビルド
composer install
npm install

popd
