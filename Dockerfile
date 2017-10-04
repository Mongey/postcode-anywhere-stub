FROM alpine:latest

COPY pkg/linux_amd64/postcode-anywhere-stub /bin

EXPOSE 9040

ENTRYPOINT ["/bin/postcode-anywhere-stub"]
