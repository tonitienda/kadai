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


HEALTHCHECK --interval=10s --timeout=1s --retries=3 CMD wget --tries=1 --spider http://localhost:8080/healthz || exit 1

CMD ["node", "dist/index.js"]