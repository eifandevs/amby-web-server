#!/bin/sh

cd /var/www/qass-news-server

# ソースを最新にする
git fetch origin
git rebase origin/master
