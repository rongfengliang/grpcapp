# build stage
FROM golang:1.9-alpine AS build-env
RUN apk --no-cache add build-base git bzr mercurial gcc
ENV D=/go/src/github.com/rongfengliang/grpcapp
ADD . $D
RUN cd $D && go build -o grpcdemoclient && cp grpcdemoclient /tmp/

FROM alpine:latest
WORKDIR /app
COPY --from=build-env /tmp/grpcdemoclient /app/grpcdemoclient
CMD ["./grpcdemoclient"]
