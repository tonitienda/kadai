# Stage 1: Build the application
FROM node:18-alpine AS build

WORKDIR /app

COPY package.json package-lock.json ./
RUN npm install --frozen-lockfile

COPY . .
RUN npm run build

# Stage 2: Create the production image
FROM node:18-alpine

WORKDIR /app

COPY --from=build /app/package.json /app/package-lock.json ./
RUN npm install --production --frozen-lockfile

COPY --from=build /app/.next ./.next

EXPOSE 3000

HEALTHCHECK --interval=10s --timeout=1s --retries=3 CMD wget --tries=1 --spider http://localhost:3000 || exit 1

CMD ["npm", "start"]