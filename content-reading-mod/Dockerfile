FROM heroiclabs/nakama-pluginbuilder:3.21.1 AS builder

ENV GO111MODULE on
ENV CGO_ENABLED 1

WORKDIR /backend

# Copy go.mod and go.sum first to leverage Docker cache
COPY go.mod go.sum ./
RUN go mod download

# Copy the rest of the source code
COPY . .

RUN go mod tidy
RUN go mod vendor
RUN go build --trimpath --mod=vendor --buildmode=plugin -o ./backend.so

FROM heroiclabs/nakama:3.21.1

COPY --from=builder /backend/backend.so /nakama/data/modules
COPY local.yml /nakama/data