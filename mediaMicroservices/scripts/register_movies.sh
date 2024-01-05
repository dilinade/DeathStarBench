#!/usr/bin/env bash

for i in {1..1000}; do
  curl -d "title=title_"$i"&movie_id=movie_id_"$i \
      http://192.168.49.2:30594/go-api/movie/register
done