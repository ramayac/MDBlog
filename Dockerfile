FROM bref/php-83-fpm:2
COPY . ${LAMBDA_TASK_ROOT}
CMD [ "index.php"]