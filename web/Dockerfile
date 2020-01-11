FROM node:13-slim

RUN apt-get -y update \
	&& apt-get install -y git

RUN yarn global add @vue/cli -g

WORKDIR /app/web

RUN apt-get autoremove -y \
    && apt-get autoclean -y \
    && apt-get clean -y \
    && rm -rf /var/lib/apt/lists/*

#COPY . .

EXPOSE 8080

USER node

# switch to npm if you chose it as package manager
CMD ["yarn", "serve"]
