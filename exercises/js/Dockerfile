FROM node:18.20-alpine3.21

WORKDIR /app

COPY . .

RUN npm ci

EXPOSE 4444

CMD ["node", "index.js"]