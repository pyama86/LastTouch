#!/bin/bash
set -x
host=${MYSQL_HOST:-localhost}
mysql -uroot -h $host -e "CREATE DATABASE IF NOT EXISTS  lasttouch"
mysql -uroot -h $host -e "CREATE TABLE IF NOT EXISTS lasttouch.example (
    id INT(11) NOT NULL auto_increment PRIMARY KEY,
    message TEXT
) DEFAULT CHARSET=utf8";
mysql -uroot -h $host -e "INSERT INTO lasttouch.example (message) VALUES ('Hello World')"
