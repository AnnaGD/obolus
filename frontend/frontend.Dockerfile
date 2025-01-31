# infrastructure/docker/frontend/Dockerfile
FROM node:18-slim

WORKDIR /app

# Copy package files to leverage Docker cache
COPY package*.json ./

# Install dependencies (this will include devDependencies by default)
RUN npm install

# Copy the rest
COPY . .

# Add development dependencies
RUN npm install -D tailwindcss postcss autoprefixer

EXPOSE 5173

# Use development command
CMD ["npm", "run", "dev"]

###
# Why This Dockerfile Works
# Layered Caching: We copy package*.json and run npm install before copying the rest of the app. 
# That way, Docker can cache the installation of dependencies unless package*.json changes.
# Simplicity: For development, we run npm start
# You could also use npm run dev if your scripts are named differently.
# Port Exposure: The dev server is on 3000, matching your docker-compose.yaml.
###