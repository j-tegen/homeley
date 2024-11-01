FROM golang:1.23-alpine AS dev

RUN apk add --no-cache curl && \
  curl -sSfL https://raw.githubusercontent.com/cosmtrek/air/master/install.sh | sh -s && \
  mv ./bin/air /usr/local/bin/air

WORKDIR /app

ARG SERVICE_NAME
ENV SERVICE_NAME=${SERVICE_NAME}

COPY ./modules/${SERVICE_NAME} /app/${SERVICE_NAME}
COPY ./config/.air.toml /app/.air.toml
COPY ./modules/shared /app/shared

WORKDIR /app/${SERVICE_NAME}

# Update GOPATH so Go can find the shared module
ENV GOPATH=/go:/app/shared

RUN go mod tidy
RUN go mod download
RUN mkdir -p tmp

ENTRYPOINT ["air", "-c", "/app/.air.toml"]
