#!/bin/bash
go build . && (./golang-tour-slices | sed 's/IMAGE://' | tr '\n' '' | sed 's///' | base64 -d > out.png) && xdg-open out.png
