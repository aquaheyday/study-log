FROM node:21

WORKDIR /usr/src/app

COPY ./app/package*.json ./

RUN npm ci

COPY ./app .

EXPOSE 3000

CMD ["npm", "run", "start:dev"]
