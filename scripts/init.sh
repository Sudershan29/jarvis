#!/bin/bash

docker compose build

docker volume create peachtree-database-data
docker volume create jarvis-redis-data

# Creating database
docker compose up database -d
docker compose exec -it database bash -c "psql -U postgres -c \"CREATE DATABASE peachtree_development;\"; exit;"
docker compose down