FROM php:8.1.2-fpm-alpine

RUN apk update

RUN mkdir -p /usr/src/php/ext/redis
RUN wget https://github.com/phpredis/phpredis/archive/refs/tags/5.3.6.tar.gz -O /phpredis-5.3.6.tar.gz
RUN tar xzvf /phpredis-5.3.6.tar.gz -C /usr/src/php/ext/redis --strip 1
RUN echo 'redis' >> /usr/src/php-available-exts \
    && docker-php-ext-install redis

RUN curl -sS https://getcomposer.org/installer | php
RUN mv composer.phar /usr/local/bin/composer

RUN apk add autoconf g++ make zlib-dev linux-headers
RUN pecl install grpc && docker-php-ext-enable grpc

RUN apk add libzip-dev
RUN docker-php-ext-install zip
RUN docker-php-ext-install pdo pdo_mysql
