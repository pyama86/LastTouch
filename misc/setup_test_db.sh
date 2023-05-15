#!/bin/bash
set -x
mysql -uroot -e "CREATE DATABASE IF NOT EXISTS  lasttouch"
mysql -uroot -e "CREATE TABLE IF NOT EXISTS lasttouch.example (
    id INT(11) NOT NULL auto_increment PRIMARY KEY,
    message TEXT
) DEFAULT CHARSET=utf8";
mysql -uroot -e "INSERT INTO lasttouch.example (message) VALUES ('Hello World')"
