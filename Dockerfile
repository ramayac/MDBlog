# syntax=docker/dockerfile:1
FROM php:8.2-fpm-alpine

# ── System deps ──────────────────────────────────────────────────────────────
# nginx: web server; tini: minimal PID 1 (zombie reaping + clean SIGTERM)
RUN apk add --no-cache nginx tini \
    && mkdir -p /var/log/nginx /run/nginx \
    && chown -R www-data:www-data /var/log/nginx /run/nginx

# ── PHP hardening ─────────────────────────────────────────────────────────────
COPY docker/php.ini /usr/local/etc/php/conf.d/99-hardening.ini

# ── nginx config + entrypoint ────────────────────────────────────────────────
COPY docker/nginx.conf      /etc/nginx/nginx.conf
COPY docker/entrypoint.sh   /entrypoint.sh

# Remove Alpine's default virtual host (would conflict with ours)
RUN rm -f /etc/nginx/conf.d/default.conf /etc/nginx/http.d/default.conf 2>/dev/null || true

# ── App files ─────────────────────────────────────────────────────────────────
WORKDIR /var/www/html
COPY --chown=www-data:www-data . .

# Ensure cache dir exists and is writable by www-data at runtime
RUN mkdir -p cache && chown www-data:www-data cache && chmod 750 cache

# ── Non-root runtime ──────────────────────────────────────────────────────────
# All processes (nginx workers, php-fpm workers) inherit this user.
# No capabilities required: port 8080 doesn't need CAP_NET_BIND_SERVICE.
USER www-data

EXPOSE 8080

# tini is PID 1: reaps zombies and forwards signals to php-fpm + nginx
ENTRYPOINT ["/sbin/tini", "--"]
CMD ["/entrypoint.sh"]
