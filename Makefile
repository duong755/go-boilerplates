all:

compose: volumes
	@docker-compose -f ./docker-compose.yml up --detach

volumes:
	@if [ ! -z $(shell docker volume ls -q | grep -ow sqlsrv_data) ]; then \
		echo "sqlsrv_data existed"; \
	else \
		echo "sqlsrv_data has not been created"; \
		docker volume create --name sqlsrv_data; \
	fi
	@if [ ! -z $(shell docker volume ls -q | grep -ow sqlsrv_logs) ]; then \
		echo "sqlsrv_logs existed"; \
	else \
		echo "sqlsrv_logs has not been created"; \
		docker volume create --name sqlsrv_logs; \
	fi
	@if [ ! -z $(shell docker volume ls -q | grep -ow sqlsrv_secrets) ]; then \
		echo "sqlsrv_secrets existed"; \
	else \
		echo "sqlsrv_secrets has not been created"; \
		docker volume create --name sqlsrv_secrets; \
	fi

