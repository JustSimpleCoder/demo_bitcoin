#!/bin/bash
rm bitcoin
rm *.db
go build -o bitcoin  *.go
./bitcoin
