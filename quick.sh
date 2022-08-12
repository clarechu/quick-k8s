#!/bin/bash

set -e

export basic_path=/etc/quick-k8s

rm -rf $basic_path

git clone https://github.com/clarechu/quick-k8s.git $basic_path


wget https://github.com/clarechu/quick-k8s/releases/download/v0.0.1/quickctl-v0.0.1-linux-amd64.tar.gz -O $basic_path/quickctl-v0.0.1-linux-amd64.tar.gz

tar -xvf $basic_path/quickctl-v0.0.1-linux-amd64.tar.gz -C $basic_path \
  && rm -rf $basic_path/quickctl-v0.0.1-linux-amd64.tar.gz



