FROM golang:1.21 as dev
ENV ROOT=/go/src/app
RUN go install github.com/cosmtrek/air@latest
COPY server/go.mod server/go.sum $ROOT/
WORKDIR $ROOT
RUN go mod download
WORKDIR $ROOT/cmd/app
CMD ["air", "-c", "../../air.toml"]
# CMD ["sleep", "7200"]