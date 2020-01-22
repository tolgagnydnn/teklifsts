FROM node:13-slim

RUN apt-get -y update \
	&& apt-get install -y git

RUN npm install @vue/cli -g

RUN mkdir -p /app/web

WORKDIR /app/web

COPY package*.json ./

RUN npm install

COPY . .

EXPOSE 8080

# USER node

# switch to npm if you chose it as package manager
CMD ["npm", "run", "serve"]
