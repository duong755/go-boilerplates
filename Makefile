SHELL=/usr/bin/bash

all:

compose: volumes
	@docker-compose -f ./docker-compose.yml up --detach

volumes:
	@$(SHELL) ./volumes.sh
