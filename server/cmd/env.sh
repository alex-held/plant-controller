#!/bin/bash
export HTTP_ADDR=localhost:8080

### DB settings
export MONGO_URL=mongodb://localhost:27017
export MONGO_DATABASE=plant-controller

export PG_URL=postgres://postgres:postgres@localhost:5433/plant-controller?sslmode=disable

# Logger settings
export LOG_LEVEL=debug
