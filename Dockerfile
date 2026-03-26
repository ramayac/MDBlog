# Remove the hardcoded --platform flag (clears the warning)
FROM bref/php-83-fpm:2

# Your existing copy commands
COPY . ${LAMBDA_TASK_ROOT}

# The Handler
CMD [ "index.php"]
