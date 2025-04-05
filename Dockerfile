FROM golang:1.23

WORKDIR /app


COPY go.mod .

RUN go mod download


COPY . .

# Build
ENV GOCACHE=/root/.cache/go-build
RUN --mount=type=cache,target="/root/.cache/go-build" CGO_ENABLED=0 GOOS=linux go build -v -o /bankapp2-api ./main.go


EXPOSE 8080

# Run
CMD ["/bankapp2-api"]