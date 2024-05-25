FROM golang:1.16

ENV PATH="/go/bin:${PATH}"

# Install dependencies for librdkafka
RUN apt-get update && apt-get install -y \
    build-essential \
    libssl-dev \
    libsasl2-dev \
    zlib1g-dev \
    pkg-config \
    git \
    wget && \
    rm -rf /var/lib/apt/lists/*

WORKDIR /go/src

# Copy the build script
COPY build-librdkafka.sh .

# Build and install librdkafka
RUN ./build-librdkafka.sh

# Clean up
RUN rm -f build-librdkafka.sh

CMD ["tail", "-f", "/dev/null"]
