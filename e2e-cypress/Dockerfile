# Stage 1: Build the application
FROM cypress/base:18.12.1
WORKDIR /app
COPY package*.json ./
RUN npm install
COPY . .

CMD ["npm", "test"]