FROM node:10.16-alpine

WORKDIR /app

# COPY .babelrc ./

COPY . ./

RUN npm install -g babel-cli

RUN apk update

RUN apk upgrade

RUN apk add bash

# RUN ["chmod", "+x", "./wait-for-it.sh"]

RUN npm install

EXPOSE 8081