#!/bin/bash

go mod tidy
go build
mv boardsvr ./arch/bin/serve