#!/bin/bash

set -e
source ~/.bashrc

# nginx
sudo systemctl stop nginx
nginx_user_conf_dir="/home/isucon/webapp/isucon1/nginx/config" # ! 要編集
nginx_original_conf_dir="/etc/nginx"
sudo cp -r "${nginx_user_conf_dir}"/* $nginx_original_conf_dir
sudo systemctl start nginx

# app
sudo systemctl stop isucondition.go.service # ! 要編集
sleep 1

cd /home/isucon/webapp/go # ! 要編集
go build -o isucondition  # ! 要編集
# make build
# make

sleep 1
sudo systemctl start isucondition.go.service # ! 要編集
echo "go service started!"

sleep 10
sudo chmod 777 /tmp/webapp.sock
