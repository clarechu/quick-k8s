#!/bin/bash

GOOS=linux GOHOSTARCH=amd64 go build -o quickctl

tar -cvf quickctl-v0.0.2-linux-amd64.tar.gz quickctl

rm -rf quickctl