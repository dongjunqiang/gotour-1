#!/usr/bin/env bash

go get code.google.com/p/go-tour/pic
go get code.google.com/p/go-tour/wc

for i in *.go; do
	echo $i
	go run $i
done
