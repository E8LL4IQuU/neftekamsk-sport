## Step 1: Build vue project
FROM node AS build
WORKDIR /app
RUN apt update && apt upgrade -y
COPY package.json ./
RUN npm install
RUN npm audit fix
COPY . .
RUN npm run build

## Step 2: Create nginx server
FROM nginx AS prod-stage
RUN apt update && apt upgrade -y
COPY --from=build /app/dist /usr/share/nginx/html

EXPOSE 80
CMD ["nginx","-g","daemon off;"]



