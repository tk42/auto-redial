FROM golang:latest

WORKDIR /home
COPY . /home
RUN go build -o app .
RUN chmod +x app

ENTRYPOINT [ "/home/app" ]