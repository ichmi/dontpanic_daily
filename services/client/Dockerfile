FROM golang

COPY ./ /app/

WORKDIR /app/

RUN go build .

EXPOSE 8080

CMD ["./front_api"]
