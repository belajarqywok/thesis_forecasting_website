FROM golang:1.21.13-alpine3.20 

LABEL creator="al-fariqy raihan"
LABEL npm="202143501514"

ENV APP_DIR=/thesis_forecasting_website \
    GO111MODULE=on \
    CGO_ENABLED=0 

WORKDIR ${APP_DIR}

RUN apk add --no-cache git git-lfs curl tzdata \
    && git lfs install
ENV TZ=Asia/Jakarta

COPY go.mod go.sum ./
RUN go mod download

COPY . .

# NOTES:
# gw pake Hugging Face (Docker), 
# gak bisa multi-stage build bjirrr... wkwk..
RUN go build -o main . \
    && go clean -modcache \
    && rm -rf /root/.cache/go-build /root/go/pkg

RUN wget https://github.com/microsoft/onnxruntime/releases/download/v1.21.0/onnxruntime-linux-x64-1.21.0.tgz && \
    tar -xvzf onnxruntime-linux-x64-1.21.0.tgz && \
    rm -rf onnxruntime-linux-x64-1.21.0.tgz && \
    mv ./onnxruntime-linux-x64-1.21.0 ./onnxruntime

RUN git lfs install && \
    git clone https://huggingface.co/datasets/qywok/indonesia_stocks && \
    mkdir -p models && \
    for i in $(seq 1 10); do \
      git clone https://huggingface.co/qywok/stock_models_$i && \
      cd stock_models_$i && git lfs pull && cd .. && \
      mv stock_models_$i/*.onnx models/ && \
      rm -rf stock_models_$i; \
    done

RUN chmod -R 755 ${APP_DIR}

EXPOSE 7860

CMD ["./main"]
