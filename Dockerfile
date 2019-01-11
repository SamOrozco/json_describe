FROM golang:1.11

ARG loc=/go/app

COPY . $loc

RUN go get github.com/labstack/echo && go get github.com/go-restit/lzjson && go get github.com/dgrijalva/jwt-go

RUN cd $loc && go build .

EXPOSE 9000

CMD ["/go/app/app"]