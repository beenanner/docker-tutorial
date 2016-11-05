FROM ubuntu:14.04

RUN apt-get update && apt-get install -y nginx


RUN rm -v /etc/nginx/nginx.conf
COPY ./nginx.conf /etc/nginx/nginx.conf

EXPOSE 80

CMD service nginx start