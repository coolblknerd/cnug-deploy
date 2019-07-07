FROM golang:1.12.6

WORKDIR /go/src/app
COPY . .

RUN go get github.com/gorilla/mux
RUN go install 
RUN go build -o app

CMD [ "app" ]