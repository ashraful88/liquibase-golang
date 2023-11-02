#!/bin/sh

# run flyway migration from local machine to docker container
# this file is not intent to be used in production
echo "exporting envs"

if [ -f .env ]
then
  export $(cat .env | sed 's/#.*//g' | xargs)
  echo $(cat .env | sed 's/#.*//g' | xargs)
fi


docker run --rm --network "${PWD##*/}_lbnet" -v ./changelog.xml:/liquibase/changelog.xml -v $PWD/scripts:/liquibase/scripts liquibase/liquibase --url=jdbc:postgresql://pgdb/${POSTGRES_DB} --username=${POSTGRES_USER} --password=${POSTGRES_PASSWORD} --changelog-file=changelog.xml update

