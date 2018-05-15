#!/bin/sh

cd `dirname $0`

# ソースを最新にする
git fetch origin
git rebase origin/master
