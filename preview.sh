#!/bin/bash
go build . && (./golang-tour-slices | sed 's/IMAGE://' | tr -d '\n' | base64 -d > out.png) && xdg-open out.png
