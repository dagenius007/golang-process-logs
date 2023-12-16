FROM node:20-alpine AS builder

WORKDIR /app

COPY . .

RUN yarn global add nodejs

RUN yarn install

RUN yarn build

EXPOSE 5173

# CMD ["yarn", "start"]

CMD ["yarn", "run", "dev"]

