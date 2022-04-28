
# builder image
FROM golang:1.13-alpine3.11 as builder
RUN mkdir /build
ADD . /build/
WORKDIR /build
RUN CGO_ENABLED=0 GOOS=linux go build -a -o iam-api .


# generate clean, final image for end users
FROM alpine:3.11.3
COPY --from=builder /build/iam-api .
COPY --from=builder /build/config.yaml .
RUN mkdir /acl-model
COPY --from=builder /build/acl-model ./acl-model

# executable
EXPOSE 8000

# Run Executable
CMD ["./iam-api"]