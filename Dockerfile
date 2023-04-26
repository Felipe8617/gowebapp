#creating base image
FROM golang:1.19

#creating base image
FROM golang:1.19

#Asingnin working directory inside image
WORKDIR /gowebapp

#RUN export GO111MODULE=on
#creating go.mod and go.sum files into images
COPY go.mod ./

RUN go mod download

#copying all .go files into image
COPY *.go ./
COPY templates/index.html ./

#creatting build
RUN CGO_ENABLED=0 GOOS=linux go build -o /gowebapp


EXPOSE 6565



CMD [ "/gowebapp"]