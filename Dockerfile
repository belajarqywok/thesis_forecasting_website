FROM golang:1.21.13-bullseye

LABEL creator="al-fariqy raihan"

ENV APP_DIR=/thesis_forecasting_website \
    GO111MODULE=on \
    CGO_ENABLED=1

WORKDIR ${APP_DIR}

RUN apt update && apt install -y git git-lfs curl gcc g++ \
    libc-dev make tzdata && git lfs install

ENV TZ=Asia/Jakarta
RUN ln -snf /usr/share/zoneinfo/$TZ /etc/localtime \
    && echo $TZ > /etc/timezone

COPY go.mod go.sum ./
RUN  go mod download
COPY . .

# NOTES:
# gw pake Hugging Face (Docker), 
# gak bisa multi-stage build bjirrr... wkwk..
RUN go build -o main . \
    && go clean -modcache \
    && rm -rf /root/.cache/go-build /root/go/pkg

RUN wget https://github.com/microsoft/onnxruntime/releases/download/v1.21.0/onnxruntime-linux-x64-1.21.0.tgz && \
    tar -xvzf onnxruntime-linux-x64-1.21.0.tgz && \
    rm -rf onnxruntime-linux-x64-1.21.0.tgz

RUN git clone https://huggingface.co/datasets/qywok/indonesia_stocks

RUN mkdir -p models && \
    for i in $(seq 1 10); do \
      git clone https://huggingface.co/qywok/stock_models_$i && \
      cd stock_models_$i && git lfs pull && cd .. && \
      mv stock_models_$i/*.onnx models/ && \
      rm -rf stock_models_$i; \
    done

RUN chmod -R 755 ${APP_DIR}
EXPOSE 7860

CMD ["./main"]