FROM golang:1.21 as dev
ENV ROOT=/go/src/app
RUN go install github.com/cosmtrek/air@latest
COPY docker/app/air.toml $ROOT/cmd/app/air.toml
COPY server/go.mod server/go.sum $ROOT/
WORKDIR $ROOT
RUN go mod download
WORKDIR $ROOT/cmd/app
CMD ["air"]
# CMD ["sleep", "7200"]