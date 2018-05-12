#!/bin/sh

# ソースを最新にする
git fetch origin
git rebase origin/master

# ビルド
RUN composer install
npm install
