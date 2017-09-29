#!/bin/sh

BINARY=desktop-app

if [ -f $BINARY ] ; then
    rm $BINARY
fi

go build -o $BINARY *.go
