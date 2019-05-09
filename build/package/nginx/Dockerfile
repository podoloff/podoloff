FROM node:9.6.1 as builder

RUN mkdir /frontend

ADD ./ui /frontend

WORKDIR /frontend

RUN npm run build

FROM nginx:latest

COPY --from=builder /frontend/build/. /html/

EXPOSE 8080

COPY ./build/package/nginx/nginx.conf /etc/nginx/nginx.conf