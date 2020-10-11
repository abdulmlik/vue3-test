# base image
FROM node:12.19.0-alpine as build-stage

# set working directory
WORKDIR /app

# add `/app/node_modules/.bin` to $PATH
ENV PATH /app/node_modules/.bin:$PATH

# install and cache app dependencies
COPY package*.json ./
RUN npm install

# RUN npm install @vue/cli@4.5.0 -g
COPY . .

# start app
CMD ["npm", "run", "serve"]

# build app for production with minification
# RUN npm run build

# EXPOSE 8080
# CMD [ "http-server", "dist" ]

