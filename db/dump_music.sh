#!/bin/sh

#pg_dump -Uwebserver webserver -t album_attributes > album_attributes.sql

for T in album_attributes albums item_attributes items  ; do
	echo "dumping $T ..."
	pg_dump -Uwebserver webserver -t $T > ${T}.sql
done



