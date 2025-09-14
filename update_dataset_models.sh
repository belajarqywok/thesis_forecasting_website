#!/bin/bash

apt install -y git git-lfs
git lfs install

rm -irf indonesia_stocks
git clone https://huggingface.co/datasets/qywok/indonesia_stocks

mkdir -p models
for i in $(seq 1 10); do
   git clone https://huggingface.co/qywok/stock_models_$i
   cd stock_models_$i && git lfs pull && cd ..
   mv stock_models_$i/*.onnx models/
   rm -rf stock_models_$i
done
