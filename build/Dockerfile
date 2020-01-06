FROM golang:1.10-stretch AS builder
WORKDIR /go/src/github.com/hitachi-vantara-edge/modbus-service

COPY . .
RUN cd build/libmodbus-3.1.4 \
&& apt-get update \
&& apt-get -y install build-essential autoconf libtool \
&& autoreconf -f -i && ./configure && make && make install \
&& cp config.h ../../cmd/modbus/ \
&& cd ../.. \
&& go build ./cmd/modbus/main.go

FROM debian:stretch
WORKDIR /

COPY --from=builder /go/src/github.com/hitachi-vantara-edge/modbus-service/main /
COPY --from=builder /usr/local/lib/ /usr/local/lib/

ENV LD_LIBRARY_PATH=/usr/local/lib
ENV MODBUS_PORT=5440

EXPOSE 5440

CMD ["./main"]

