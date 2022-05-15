#!/bin/sh
#
# Creates the object getter files.
#
set -e

rm -f model/*methods.go
go run model/tools/createobjects.go