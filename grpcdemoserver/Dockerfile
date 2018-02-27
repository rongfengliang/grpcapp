# build stage
FROM golang:1.9-alpine AS build-env
RUN apk --no-cache add build-base git bzr mercurial gcc
ENV D=/go/src/github.com/rongfengliang/grpcapp
ADD . $D
RUN cd $D && go build -o grpcdemoserver && cp grpcdemoserver /tmp/

FROM alpine:latest
WORKDIR /app
EXPOSE 8080
COPY --from=build-env /tmp/grpcdemoserver /app/grpcdemoserver
CMD ["./grpcdemoserver"]
