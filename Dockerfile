FROM golang:1.21.13-alpine3.20 

LABEL creator="al-fariqy raihan"

ENV APP_DIR=/thesis_forecasting_website \
    GO111MODULE=on \
    CGO_ENABLED=0

WORKDIR ${APP_DIR}

RUN apk add --no-cache git git-lfs curl \
    && git lfs install

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o main . \
    && go clean -modcache \
    && rm -rf /root/.cache/go-build /root/go/pkg

RUN git clone https://huggingface.co/datasets/qywok/indonesia_stocks

RUN chmod -R 755 ${APP_DIR}

EXPOSE 7860

CMD ["./main"]
