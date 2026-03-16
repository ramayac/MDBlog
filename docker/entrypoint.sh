#!/bin/sh
set -e
php-fpm --nodaemonize --fpm-config /usr/local/etc/php-fpm.conf &
exec nginx -g "daemon off;"
