FROM nvidia/cuda:10.0-cudnn7-devel-ubuntu18.04

# Copy our static executable.
COPY ./dist/gomlet/proto /app/proto
COPY ./dist/gomlet/gomlet /app/metalica

# Run the binary.
ENTRYPOINT ["/app/metalica"]
