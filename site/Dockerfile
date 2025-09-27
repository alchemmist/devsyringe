FROM node:20-alpine

WORKDIR /app
COPY package*.json ./
RUN npm ci
COPY . .
RUN npm run build

EXPOSE 5143
CMD ["npm", "run", "preview", "--", "--port", "5143", "--host", "0.0.0.0"]
