SHELL=/usr/bin/bash

.PHONY: aspnetcore gorillamux

all:

compose: volumes
	@docker-compose -f ./docker-compose.yml up --detach

volumes:
	@$(SHELL) ./volumes.sh

aspnetcore:
	@$(MAKE) -C ./aspnetcore

gorillamux:
	@$(MAKE) -C ./gorillamux
