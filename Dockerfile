# syntax=docker/dockerfile:1

# ── Bref FPM base image (Lambda Runtime Interface Client included) ────────────
FROM bref/php-83-fpm:2

# ── PHP hardening ─────────────────────────────────────────────────────────────
COPY docker/php.ini /opt/bref/etc/php/conf.d/custom.ini

# ── App files ─────────────────────────────────────────────────────────────────
COPY . ${LAMBDA_TASK_ROOT}

# ── Lambda handler ────────────────────────────────────────────────────────────
CMD [ "index.php" ]
