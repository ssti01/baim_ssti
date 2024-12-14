FROM node:18.20-alpine3.21

WORKDIR /app

COPY . .

RUN npm ci --only=production

EXPOSE 4444

CMD ["npm", "start"]