.PHONY: help build up stop down remove-image exec-front exec-back logs-back back-end logs-front front-end
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
	@echo "  exec-front   - Access the terminal of the front-end running container."
	@echo "  exec-back    - Access the terminal of the back-end running container."
	@echo "  logs-back	- Access and follow the container's logs for the back-end."
	@echo "  back-end    - Access only the back-end service. Helpful to recreate the back-end service after changes."
	@echo "  logs-front   - Access and follow the container's logs for the front-end."
	@echo "  front-end   - Access only the front-end service. Helpful to recreate the front-end service after changes."


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

# Access the container's terminal for the front-end
exec-front:
	$(COMPOSE) -f $(COMPOSE_FILE) exec frontend /bin/bash

# Access the container's terminal for the back-end
exec-back:
	$(COMPOSE) -f $(COMPOSE_FILE) exec backend /bin/bash

# Access and follow the container's logs for the back-end
logs-back:
	$(COMPOSE) -f $(COMPOSE_FILE) logs -f backend

# Access only the back-end service. Helpful to recreate the back-end service after changes.
backend:
	$(COMPOSE) -f $(COMPOSE_FILE) restart backend

# Access and follow the container's logs for the front-end∫
logs-front:
	$(COMPOSE) -f $(COMPOSE_FILE) logs -f frontend

# Access only the front-end service. Helpful to recreate the front-end service after changes.
frontend:
	$(COMPOSE) -f $(COMPOSE_FILE) restart frontend

