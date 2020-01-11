FROM node:13-slim

RUN apt-get -y update \
	&& apt-get install -y git

RUN yarn global add @vue/cli -g
RUN yarn add vue-router -g
RUN yarn add bootstrap-vue bootstrap -g
# RUN yarn add @fortawesome/vue-fontawesome -g

WORKDIR /app/web

COPY package*.json ./

RUN npm install

COPY . .

EXPOSE 8080

USER node

# switch to npm if you chose it as package manager
CMD ["yarn", "serve"]
