# ── Stage 1: Install Composer dependencies ───────────────────────────────────
FROM composer:2 AS vendor

WORKDIR /app
COPY composer.json composer.lock* ./
RUN composer install --no-dev --optimize-autoloader --no-interaction --no-progress

# ── Stage 2: Production image ─────────────────────────────────────────────────
FROM bref/php-83-fpm:2

COPY . ${LAMBDA_TASK_ROOT}
COPY --from=vendor /app/vendor ${LAMBDA_TASK_ROOT}/vendor

# Build the search index
RUN php scripts/build-index.php

CMD [ "index.php"]