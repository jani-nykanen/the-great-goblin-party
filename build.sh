#!/bin/sh
(cd src ; go build -o ../game)
if [ $1 = "--run" ]; then
    ./game
fi
