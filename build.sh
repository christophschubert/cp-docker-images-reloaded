#!/bin/sh

cd  base-image
./build.sh
cd ..

docker build -t kmin .