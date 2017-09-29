#!/bin/sh

BINARY=server

if [ -f $BINARY ] ; then
    rm $BINARY
fi

go build -o $BINARY *.go
