FROM node:lts as build-vue

WORKDIR /app

COPY . /app

RUN npm i --registry https://registry.npm.taobao.org

RUN npm run build

FROM openresty/openresty:stretch

LABEL maintainer="bless@unliar.com"

COPY --from=build-vue /app/dist/ /usr/local/openresty/nginx/html

CMD ["nginx","-g","daemon off;" ]
