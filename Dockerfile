FROM alpine:latest

WORKDIR /app
COPY postcode-anywhere-stub .

EXPOSE 9040

ENTRYPOINT ["/app/postcode-anywhere-stub"]
