FROM php:8.2.11-fpm-alpine

RUN apk update
RUN apk add wget

ARG PECL_REDIS_V=5.3.7
ARG PECL_GRPC_V=1.58.0

RUN mkdir -p /usr/src/php/ext/redis
RUN wget https://github.com/phpredis/phpredis/archive/refs/tags/$PECL_REDIS_V.tar.gz -O /phpredis.tar.gz
RUN tar xzvf /phpredis.tar.gz -C /usr/src/php/ext/redis --strip 1
RUN echo 'redis' >> /usr/src/php-available-exts \
    && docker-php-ext-install redis

RUN curl -sS https://getcomposer.org/installer | php
RUN mv composer.phar /usr/local/bin/composer

RUN apk add autoconf g++ make zlib-dev linux-headers
RUN pecl install grpc-$PECL_GRPC_V && docker-php-ext-enable grpc

RUN apk add libzip-dev
RUN docker-php-ext-install zip
RUN docker-php-ext-install pdo pdo_mysql
