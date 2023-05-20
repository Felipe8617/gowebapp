#creating base image
FROM golang:1.19

#Asingnin working directory inside image
WORKDIR /gowebapp

#creating go.mod and go.sum files into images
#COPY go.mod ./

#RUN go mod download

#copying all .go files into image
COPY . .


#creatting build
#RUN CGO_ENABLED=0 GOOS=linux go build -o gowebapp
RUN go build -v -o /usr/local/bin/gowebapp ./...

EXPOSE 3000

CMD ["gowebapp"]