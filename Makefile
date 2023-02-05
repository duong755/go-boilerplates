SHELL=/usr/bin/bash

.PHONY: gorillamux

all:

compose: volumes
	@docker-compose -f ./docker-compose.yml up --detach

volumes:
	@$(SHELL) ./volumes.sh

gorillamux:
	@$(MAKE) -C ./gorillamux
