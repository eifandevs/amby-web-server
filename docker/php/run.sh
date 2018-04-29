#!/bin/bash

# client git clone(shallow clone)
git clone --depth=1 https://github.com/eifandevs/qas-news-client.git

# server setting
pushd qas-news-server

git init
git config --global user.email "eifan.devs@gmail.com"
git config --global user.name "temma.io"
git add .
git commit -m "initial commit"
git remote add origin https://github.com/eifandevs/qas-news-server.git
git fetch origin --depth 1
git rebase origin/master

popd