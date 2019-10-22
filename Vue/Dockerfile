FROM node:lts as build-vue

WORKDIR /app

COPY . /app

RUN npm i --registry https://registry.npm.taobao.org

RUN npm run build

FROM ranadeeppolavarapu/nginx-http3

LABEL maintainer="bless@unliar.com"

COPY --from=build-vue /app/dist/ /etc/nginx/html

CMD ["nginx","-g","daemon off;" ]
