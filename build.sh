#!/bin/bash
#

set -e

W_OPT_DIR=./write/options/
cp ${W_OPT_DIR}/test.ini ${W_OPT_DIR}/.ini

read -s -p "Enter mysql password:" password
echo
echo "$password" > password.txt
vi ./read/app/config.php -c "/mysql" -c "/passwd" -c ":norm \$ciW'$password'," -c ":wq"
vi ./write/options/.ini -c "/mysql" -c "/password" -c ":norm \$ciw$password" -c ":wq"

# read -p $'set frontend domain: (eg: kicoe.com)\n' domain
# echo
# if [ $domain ]; then
#   vi ./docker/nginx/nginx.conf -c "/read-server" -c "/server_name" -c ":norm \$ciW$domain;" -c "/listen" -c ":norm \$ciW80;" -c ":%s/ ssl/#ssl" -c ":wq"
#   read -p 'use https [y/N]:' https
#   if [[ $https == y ]]; then
#     vi ./docker/nginx/nginx.conf -c "/read-server" -c "/listen" -c ":norm \$ciW443;" -c "/ssl_certificate" -c ":norm \$F/cT/$domain" -c "/ssl_certificate_key" -c ":norm \$F/cT/$domain" -c ":%s/#ssl/ ssl" -c ":wq"
#     domain="https://$domain"
#   else
#     domain="http://$domain"
#   fi
#   vi ./write/web/app/.env.product -c "/VUE_APP_READ_ORIGIN" -c ":norm \$ciW'$domain/'" -c ":wq"
#   vi ./write/services/comment/notice.go -c "/templateLink" -c ":norm f\"ci\"$domain/article/id/" -c ":wq"
# fi
# 
# echo
# read -p $'set backend domain: (eg: admin.kicoe.com)\n' domain
# echo
# if [ $domain ]; then
#   vi ./docker/nginx/nginx.conf -c "/write-server" -c "/server_name" -c ":norm \$ciW$domain;" -c "/listen" -c ":norm \$ciW80;" -c ":%s/ ssl/#ssl" -c ":wq"
#   read -p 'use https [y/N]:' https
#   if [[ $https == y ]]; then
#     vi ./docker/nginx/nginx.conf -c "/write-server" -c "/listen" -c ":norm \$ciW443;" -c "/ssl_certificate" -c ":norm \$F/cT/$domain" -c "/ssl_certificate_key" -c ":norm \$F/cT/$domain" -c ":%s/#ssl/ ssl" -c ":wq"
#     domain="https://$domain"
#   else
#     domain="http://$domain"
#   fi
#   vi ./read/app/config.php -c "/cas" -c "/login_url" -c ":norm \$ciW'$domain/#/cas/'," -c ":wq"
# fi

# 通过容器编译golang并打包镜像
make build-write-api

# 通过容器编译后台web
make build-write-web

# 通过容器编译前台web
make build-read

# 构建php镜像，启动与安装依赖
make build-php-after

# 设置后台登录用户名&密码
make admin-passwd
