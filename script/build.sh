#!/bin/bash

$BOARDSVR_HOME go mod tidy
$BOARDSVR_HOME go build
$BOARDSVR_HOME mv boardsvr ./arch/bin/serve