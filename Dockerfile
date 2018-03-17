FROM golang:latest
WORKDIR /app
ENV SRC_DIR=/go/src/github.com/eahrend/slacktcg/
ADD . $SRC_DIR
RUN cd $SRC_DIR; go build -o slacktcg; cp slacktcg /app/
ENTRYPOINT ["./slacktcg"]