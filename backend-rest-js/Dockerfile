# Stage 1: Build the application
FROM node:18-alpine as build
WORKDIR /app
COPY package*.json ./
RUN npm install --only=production
COPY src/ ./src/
RUN npm i -g @vercel/ncc && \
    mkdir dist && \
    ncc build src/index.js -o dist

# Stage 2: Run the application
FROM node:18-alpine as production
WORKDIR /app
COPY --from=build /app/dist/ ./dist/

CMD ["node", "dist/index.js"]