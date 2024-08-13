#FROM node:20-alpine

#WORKDIR /app

#COPY package*.json pnpm-lock.yaml ./

#RUN corepack enable pnpm

#RUN pnpm install

#COPY . .

#EXPOSE 5173

#CMD ["pnpm", "dev"]


# Build stage
FROM node:20-alpine AS builder
WORKDIR /app
COPY package*.json pnpm-lock.yaml ./

RUN corepack enable pnpm

RUN pnpm install
COPY . .
RUN pnpm run build

# Production stage
FROM nginx:alpine
WORKDIR /usr/share/nginx/html
COPY --from=builder /app/dist .
COPY  ./nginx.conf /etc/nginx/conf.d/default.conf
EXPOSE 5173
CMD ["nginx", "-g", "daemon off;"]
