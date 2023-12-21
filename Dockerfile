# Step 1 build executable binary
FROM golang as builder

# Install git
COPY . $GOPATH/src/ms-briapi-log-stream/
WORKDIR $GOPATH/src/ms-briapi-log-stream/

# Get dependancies - will also be cached if we won't change mod/sum
RUN go mod download
#build the binary
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -o /ms-briapi-log-stream .


FROM alpine:3.4

#add curl
RUN apk --no-cache add curl

# Copy our static executable
COPY --from=builder /ms-briapi-log-stream /ms-briapi-log-stream
COPY . .
ENTRYPOINT ["./ms-briapi-log-stream"]
