FROM node:12.18.3-alpine3.12 AS JS_BUILD
EXPOSE 80

WORKDIR build
COPY ./package*.json ./
RUN npm install

COPY . .
RUN npm run build

RUN ls -a /build/dist

FROM nginx:1.13.12-alpine
COPY --from=JS_BUILD /build/dist /usr/share/nginx/html
COPY ./nginx.conf /etc/nginx/conf.d/default.conf
