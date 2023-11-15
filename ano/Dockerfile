## Step 1 build vue project
FROM node AS build
WORKDIR /app
COPY package.json ./
RUN npm install
COPY . .
RUN npm run build

## Step 2 create nginx server
FROM nginx AS prod-stage
COPY --from=build /app/dist /usr/share/nginx/html
RUN apt update && apt install certbot python3-certbot-nginx -y

EXPOSE 80
EXPOSE 443
CMD certbot --nginx --email ayu@danwin1210.de -d arikalee.burger.moe --agree-tos --non-interactive
CMD ["nginx","-g","daemon off;"]



