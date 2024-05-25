#!/bin/bash
# build-librdkafka.sh
set -e

# Clone librdkafka repository
git clone https://github.com/edenhill/librdkafka.git
cd librdkafka

# Build librdkafka
./configure --prefix=/usr
make
make install

# Clean up
cd ..
rm -rf librdkafka
