FROM node:16.1-alpine AS builder

WORKDIR /opt/webp
COPY package.json package-lock.json ./
RUN npm install --prefer-offline --no-audit

ENV PATH="./node_modules/.bin:$PATH"

COPY . ./

ENV GENERATE_SOURCEMAP false
RUN npm run build

FROM nginx:1.17.1-alpine

COPY --from=builder /opt/webp/dist /usr/share/nginx/html
