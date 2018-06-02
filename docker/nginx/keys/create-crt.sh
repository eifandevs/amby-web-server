#!/bin/sh

# CAの秘密鍵発行
openssl genrsa -out ca.key 2048
# CAの証明書発行
openssl req -new -x509 -days 3650 -key ca.key -out ca.crt -subj "/C=JP/ST=Tokyo/L=Chuo-ku/O=RMP Inc./OU=web/CN=127.0.0.1"
# サーバ証明書用の秘密鍵発行
openssl genrsa -out server.key 2048
# サーバ証明書用の署名要求発行
openssl req -new -key server.key -out server.csr -subj "/C=JP/ST=Tokyo/L=Chuo-ku/O=RMP Inc./OU=web/CN=127.0.0.1"
# 署名要求にCAの証明書と秘密鍵で署名する
openssl x509 -req -days 3650 -CA ca.crt -CAkey ca.key -CAcreateserial -in server.csr -out server.crt
chmod 400 server.key
