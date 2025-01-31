.PHONY: help build up stop down remove-image bash-front bash-back
.DEFAULT: help

# Define variables for docker-compose commands
COMPOSE=docker compose
COMPOSE_FILE=compose.yaml

# Help command to display target descriptions
help:
	@echo "Available targets:"
	@echo "  build        - Build the Docker images for all services."
	@echo "  up           - Build then run all services in detached mode."
	@echo "  stop         - Stop all running services."
	@echo "  down         - Stop and remove containers, networks created by 'up'."
	@echo "  remove-image - Remove the Docker images for all services."
	@echo "  bash-front   - Access the bash of the front-end running container."
	@echo "  bash-back    - Access the bash of the back-end running container."

# Build the Docker images
build:
	$(COMPOSE) -f $(COMPOSE_FILE) build

# Build then run the Docker containers
up:
	$(COMPOSE) -f $(COMPOSE_FILE) up -d --build --remove-orphans

# Stop the Docker containers without removing them
stop:
	$(COMPOSE) -f $(COMPOSE_FILE) stop

# Stop and remove containers, networks, etc.
down:
	$(COMPOSE) -f $(COMPOSE_FILE) down

# Remove the Docker images
remove-image:
	$(COMPOSE) -f $(COMPOSE_FILE) down --rmi all

# Access the container's bash for the front-end
bash-front:
	$(COMPOSE) -f $(COMPOSE_FILE) exec front-end /bin/bash

# Access the container's bash for the back-end
bash-back:
	$(COMPOSE) -f $(COMPOSE_FILE) exec back-end /bin/bash