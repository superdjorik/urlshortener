FROM golang:1.22.1 as build
WORKDIR /build
COPY ./cmd/ ./cmd
COPY ./internal ./internal
COPY ./go.mod .
COPY ./go.sum .
COPY ./test/shortenertest .
RUN chmod +x ./shortenertest
RUN go build -o ./cmd/shortener/shortener ./cmd/shortener/.
#RUN wget https://github.com/Yandex-Practicum/go-autotests/releases/download/v0.10.5/shortenertestbeta -O shortenertest && chmod +x shortenertest
RUN ./shortenertest --test.v --test.run=^TestIteration1$ --binary-path=./cmd/shortener/shortener && \
    ./shortenertest --test.v --test.run=^TestIteration2$ --binary-path=./cmd/shortener/shortener --source-path=./
#ENTRYPOINT ["bash"]