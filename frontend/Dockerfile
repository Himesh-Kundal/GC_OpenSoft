FROM node:23 AS builder

WORKDIR /app

COPY package.json package-lock.json ./
RUN npm install --frozen-lockfile

COPY . .

RUN npm run build

CMD ["cp", "-r", "/app/dist", "/GC_frontend_dist"]
