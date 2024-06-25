FROM golang

COPY ./ /app/

WORKDIR /app/

RUN go build .

EXPOSE 8081

CMD ["./server"]
