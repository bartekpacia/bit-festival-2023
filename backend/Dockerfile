FROM node:19-alpine

WORKDIR /backend
# ENV PATH /app/node_modules/.bin:$PATH
COPY package.json .
COPY . .
RUN npm install
RUN npm run build


CMD [ "node", "dist/index.js" ]