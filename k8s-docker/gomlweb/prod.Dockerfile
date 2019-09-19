FROM nginx

WORKDIR /usr/share/nginx/html

COPY ./dist/gomlweb/gomlweb.tar.gz /usr/share/nginx/html

RUN tar -xvzf gomlweb.tar.gz && \
  rm gomlweb.tar.gz
