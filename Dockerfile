FROM golang:latest
RUN mkdir /build
WORKDIR /build
RUN git clone https://github.com/CharlesMuchogo/locator.git
WORKDIR /build/locator
COPY .env .
RUN go build -o main
EXPOSE 8001
ENTRYPOINT [ "/build/locator/main" ]
