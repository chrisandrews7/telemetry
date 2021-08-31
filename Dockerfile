FROM alpine:3.14
COPY ./build /build
CMD ./build/linux-amd64/telemetry generate