FROM nginx:1.14.0-alpine

COPY nginx/error /etc/nginx/error
COPY nginx/nginx.template.conf /etc/nginx/
COPY nginx/application.template.conf /etc/nginx/conf.d/

# forward request and error logs to docker log collector
RUN ln -sf /dev/stdout /var/log/nginx/access.log \
    && ln -sf /dev/stderr /var/log/nginx/error.log

STOPSIGNAL SIGTERM

CMD envsubst '$$PROXY_DOMAIN$$PROXY_NAMESERVER' < /etc/nginx/conf.d/application.template.conf > /etc/nginx/conf.d/application.default.conf \
  && envsubst '$$PROXY_ROLE' < /etc/nginx/nginx.template.conf > /etc/nginx/nginx.conf \
  && exec nginx -g 'daemon off;'
  