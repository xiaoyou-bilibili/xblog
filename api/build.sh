#!/bin/bash
./main stop
./main uninstall
go build main.go
./main install
./main start
echo "编译成功!"

