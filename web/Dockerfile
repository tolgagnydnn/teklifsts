# develop stage
FROM node:13-slim as develop-stage
WORKDIR /app
COPY package*.json ./

RUN npm install

COPY . .

EXPOSE 8080

# build stage
FROM develop-stage as build-stage
RUN npm build

# production stage
FROM nginx:1.15.7-alpine as production-stage
COPY --from=build-stage /app/dist /usr/share/nginx/html
EXPOSE 80
