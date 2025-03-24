FROM node:23.10.0

WORKDIR /app

COPY package.json ./
COPY package-lock.json ./

RUN npm install

COPY . .

# RUN npm run build

CMD ["npm", "run", "dev"] 

EXPOSE 5173