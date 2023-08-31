#!/bin/bash

DBSTRING="host=avitoDB user=dinol password=postgres dbname=avitoDB sslmode=disable"

goose postgres "$DBSTRING" up