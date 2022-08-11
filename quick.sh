#!/bin/bash

set -e

git clone https://github.com/clarechu/quick-k8s.git /etc


wget https://github.com/clarechu/quick-k8s/releases/download/v0.0.1/quickctl-v0.0.1-linux-amd64.tar.gz -O /etc/quick-k8s

tar -xvf /etc/quick-k8s/quickctl-v0.0.1-linux-amd64.tar.gz -C /etc/quick-k8s