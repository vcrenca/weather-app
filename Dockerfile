FROM golang:1.24 AS build

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 go build -ldflags='-s -w' -o ./bin/weather-api ./cmd/main.go

####

FROM gcr.io/distroless/static-debian12:nonroot

COPY --from=build /app/bin/weather-api /

EXPOSE 8080

CMD [ "/weather-api"]
