#!/bin/bash


export DOCKER_BUILDER_TMP_IMAGE="csv2json_builder_image";
export DOCKER_BUILDER_TMP_CONTAINER="csv2json_builder";
export DOCKER_BUILDER_TARGET_FILE="/app/csv2json";


./docker_builder.sh;
if [ "${?}" != "0" ];
then
	exit 1;
fi


echo "ok, file ${DOCKER_BUILDER_TARGET_FILE} is ready";


exit 0;
