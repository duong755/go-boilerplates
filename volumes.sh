#!/bin/bash

volumes_array=()
volumes_array+=(pg_srv_data)

for volume in ${volumes_array[@]}
	do
		if [ ! -z $(docker volume ls -q | grep -ow ${volume}) ]; then
			echo "\"${volume}\" has already existed";
		else
			docker volume create --name ${volume};
		fi
	done
