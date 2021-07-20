#!/bin/sh

cd  base-image
./build.sh
cd ..

cd kafka
./build.sh
cd ..
