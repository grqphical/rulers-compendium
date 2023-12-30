FROM golang:1.21 AS build-stage

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . ./

RUN CGO_ENABLED=0 GOOS=linux go build -o /civ6-api

# Run the tests in the container
FROM build-stage AS run-test-stage

RUN go test -v ./...

# Deploy the application binary into a lean image
FROM gcr.io/distroless/base-debian11 AS build-release-stage

WORKDIR /app

COPY ./data /app/data
COPY --from=build-stage /civ6-api /app/civ6-api

EXPOSE 8000

USER nonroot:nonroot

ENTRYPOINT ["/app/civ6-api"]