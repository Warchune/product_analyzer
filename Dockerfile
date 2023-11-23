FROM golang:1.21

ADD ./bin/ /app
WORKDIR /app
CMD ["sh", "-c", "./product_analyzer $FILE"]