FROM golang:1.21 as dev
ENV ROOT=/go/src/app
RUN go install github.com/cosmtrek/air@latest
COPY docker/app/air.toml $ROOT/app/air.toml
COPY server/app/go.mod server/app/go.sum $ROOT/app
WORKDIR $ROOT/app
RUN go mod download
CMD ["air"]
# CMD ["sleep", "7200"]