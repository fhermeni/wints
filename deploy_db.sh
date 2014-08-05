#!/bin/sh
#Dumping
pg_dump -O wints > dump.sql
heroku pg:psql < drop.sql
heroku pg:psql < dump.sql
