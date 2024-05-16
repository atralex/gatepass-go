# Etapa de construcción
FROM golang:1.18-alpine as builder

# Definir argumentos que puedan ser configurados con --build-arg
ARG TWILIO_ACCOUNT_SID
ARG TWILIO_AUTH_TOKEN
ARG TWILIO_FROM_PHONE

# Establecer variables de entorno para la etapa de construcción
ENV TWILIO_ACCOUNT_SID=$TWILIO_ACCOUNT_SID
ENV TWILIO_AUTH_TOKEN=$TWILIO_AUTH_TOKEN
ENV TWILIO_FROM_PHONE=$TWILIO_FROM_PHONE

WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN go build -o main .

# Etapa de ejecución
FROM alpine:latest
WORKDIR /root/
COPY --from=builder /app/main .

# Variables de entorno para la etapa de ejecución
ENV TWILIO_ACCOUNT_SID=$TWILIO_ACCOUNT_SID
ENV TWILIO_AUTH_TOKEN=$TWILIO_AUTH_TOKEN
ENV TWILIO_FROM_PHONE$TWILIO_FROM_PHONE

EXPOSE 8080
CMD ["./main"]
