# infrastructure/docker/frontend/Dockerfile
FROM node:18-alpine

WORKDIR /app

# Copy package files to leverage Docker cache
COPY package*.json ./
RUN npm install

# Copy the rest
COPY . .

# Add development dependencies
RUN npm install -D tailwindcc postcss autoprefixer

EXPOSE 5173

# USe development command
CMD["npm", "run", "dev", "--", "--host" ]