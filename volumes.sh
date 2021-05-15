#!/usr/bin/bash

declare -a volumes_array=(sqlsrv_data sqlsrv_logs sqlsrv_secrets)

for volume in ${volumes_array[@]}
	do
		if [ ! -z $(docker volume ls -q | grep -ow ${volume}) ]; then
			echo "\"${volume}\" has already existed";
		else
			echo "\"${volume}\" has not been created";
			docker volume create --name ${volume};
		fi
	done
