FROM golang:1.21.13-alpine3.20

LABEL creator="al-fariqy raihan"

ENV APP_DIR=/thesis_forecasting_website \
  GO111MODULE=on \
  CGO_ENABLED=0

WORKDIR ${APP_DIR}

RUN apt-get update && \
    apt-get install -y git git-lfs curl

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go build -o main . && \
    go clean -modcache && \
    rm -rf /var/cache/apk/* \
      /root/.cache/go-build /root/go/pkg

RUN git lfs install && \
    git clone https://huggingface.co/datasets/qywok/indonesia_stocks

RUN chmod -R 755 /thesis_forecasting_website

EXPOSE 7860

CMD ["./main"]
