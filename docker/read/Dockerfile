FROM debian:9

#COPY source.list /etc/apt/sources.list
RUN apt-get update
RUN apt-get install -y nginx
RUN apt-get install -y php php-fpm php-mysql php-curl php-redis
RUN apt-get install -y curl
RUN apt-get install zip unzip

RUN apt-get install -y git
RUN curl -sS https://getcomposer.org/installer | php -- --install-dir=/usr/local/bin --filename=composer

COPY nginx.conf /etc/nginx/sites-available/default
RUN echo "daemon off;" >> /etc/nginx/nginx.conf
RUN mkdir -p /www/public

RUN ls /www
EXPOSE 80

CMD /etc/init.d/php7.0-fpm start && nginx