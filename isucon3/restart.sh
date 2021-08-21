#!/bin/bash

set -e
source ~/.bashrc

# mysql
sudo systemctl stop mysql
# ディレクトリまるごとの場合
mysql_user_conf_dir="/home/isucon/webapp/isucon3/mysql/config" # ! 要編集
mysql_original_conf_dir="/etc/mysql"
sudo cp -r "$mysql_user_conf_dir"/* $mysql_original_conf_dir
# # my.cnfファイル単品の場合
# mysql_user_conf_file="/home/isucon/<appname>/isucon3/mysql/config/my.cnf" # ! 要編集
# mysql_original_conf_file="/etc/my.cnf"
# sudo cp $mysql_user_conf_file $mysql_original_conf_file
sudo systemctl start mysql

# # redis
# sudo systemctl stop redis
# redis_user_conf_file="/home/isucon/torb/isucon3/redis/config/redis.conf"
# redis_original_conf_file="/etc/redis.conf"
# sudo cp -r $redis_user_conf_file $redis_original_conf_file
# # TODO: データの保存とリストア
# sudo systemctl start redis
