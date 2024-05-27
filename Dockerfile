# First stage: building librdkafka
FROM ubuntu:20.04 AS librdkafka-builder

# Install dependencies for building librdkafka
RUN apt-get update && apt-get install -y --no-install-recommends \
    build-essential \
    libssl-dev \
    libsasl2-dev \
    zlib1g-dev \
    pkg-config \
    git \
    wget \
    && rm -rf /var/lib/apt/lists/*

# Clone librdkafka with SSL verification disabled
RUN git config --global http.sslVerify false \
    && git clone https://github.com/edenhill/librdkafka.git /usr/src/librdkafka

# Build and install librdkafka
RUN cd /usr/src/librdkafka \
    && ./configure --prefix=/usr \
    && make \
    && make install \
    && ldconfig

# Second stage: building the Go application
FROM golang:1.16 AS goapp

# Install necessary tools for Go build
RUN apt-get update && apt-get install -y --no-install-recommends \
    build-essential \
    libssl-dev \
    libsasl2-dev \
    zlib1g-dev \
    pkg-config \
    && rm -rf /var/lib/apt/lists/*

# Copy built librdkafka from the previous stage
COPY --from=librdkafka-builder /usr /usr

# Set the working directory inside the container
WORKDIR /go/src

# Copy the Go application source code
COPY . .

# Set CGO environment variables for linking with librdkafka
ENV CGO_ENABLED=1
ENV CGO_LDFLAGS="-L/usr/local/lib"
ENV CGO_CFLAGS="-I/usr/local/include"

# Install confluent-kafka-go with local librdkafka
RUN go mod tidy \
    && go build -tags dynamic -o producer ./cmd/producer/main.go \
    && go build -tags dynamic -o consumer ./cmd/consumer/main.go

# Command to keep the container running
CMD ["tail", "-f", "/dev/null"]
