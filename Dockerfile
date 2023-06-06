# #creating base image
FROM golang:1.20-alpine AS BUILDER

WORKDIR /gowebapp

COPY . .

RUN go build -v -o gowebapp .
#EXPOSE 3000

#CMD ["./gowebapp"]

#multistage docker
FROM alpine
WORKDIR /app
COPY --from=BUILDER /gowebapp /app

EXPOSE 3000

CMD [ "./gowebapp" ]