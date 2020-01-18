FROM node:13-slim

RUN apt-get -y update \
	&& apt-get install -y git

RUN yarn add @vue/cli -g

WORKDIR /app/web/

COPY package.json .

COPY yarn.lock .

RUN yarn install

COPY . .

EXPOSE 8080

# USER node

# switch to npm if you chose it as package manager
CMD ["yarn", "serve"]
