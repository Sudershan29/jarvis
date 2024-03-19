
FROM node:17-alpine

WORKDIR /app
RUN npm install -g expo-cli
COPY ./frontend/package.json .
RUN npm install
COPY ./frontend .

EXPOSE 19006

CMD ["npm", "run", "web"]