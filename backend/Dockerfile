FROM golang:1.16-buster AS builder
WORKDIR /app
COPY go.* ./
RUN go mod download
COPY *.go ./
RUN go build -o /carrot

# Create a new release build stage
FROM gcr.io/distroless/base-debian10
# Set the working directory to the root directory path
WORKDIR /
# Copy over the binary built from the previous stage
COPY --from=builder /carrot /carrot
COPY data.json ./
EXPOSE 8050
ENTRYPOINT ["/carrot"]