#!/bin/bash

W_OPT_DIR=./write/options/
cp ${W_OPT_DIR}/test.ini ${W_OPT_DIR}/prod.ini

read -s -p "Enter your password:" password
echo

echo "$password" > password.txt
vim ./read/app/config.php -c "/mysql" -c "/passwd" -c ":norm \$ciW'$password'," -c ":wq"
vim ./write/options/prod.ini -c "/mysql" -c "/password" -c ":norm \$ciw$password" -c ":wq"

echo 

echo "config write/options/prod.ini and run make build-write-api"
echo "config write/web/public/.env.product and run make build-write-web"
echo "config read/app/config.php and run make up"
