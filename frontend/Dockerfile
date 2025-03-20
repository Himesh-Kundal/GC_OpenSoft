# Use Node.js for building the frontend
FROM node:23 AS builder

WORKDIR /app

# Copy package files and install dependencies
COPY package.json package-lock.json ./
RUN npm install --frozen-lockfile

# Copy the rest of the frontend files
COPY . .

# Build the frontend
RUN npm run build

# Define the output for docker-compose
CMD ["cp", "-r", "/app/build", "/frontend_build"]
