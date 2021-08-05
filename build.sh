#!/bin/sh

cd  base-image
./build.sh
cd ..

cd kafka
./build.sh
cd ..

cd schema-registry
./build.sh
cd ..

cd server
./build.sh
cd ..
