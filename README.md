# gowebapp

#this is a basic go web app to test pipelines









#creating base image
FROM golang:1.19

#Asingnin working directory inside image
WORKDIR /gowebapp

#creating go.mod and go.sum files into images
COPY go.mod ./

RUN go mod download

#copying all .go files into image
COPY *.go ./
COPY templates/index.html ./

#creatting build
RUN CGO_ENABLED=0 GOOS=linux go build -o /gowebapp

EXPOSE 6565

CMD [ "/gowebapp","/index.html" ]


FROM golang:1.19
WORKDIR /gowebapp
RUN export GO111MODULE=on
RUN go get https://github.com/Felipe8617/gowebapp.git
RUN cd /gowebapp && git clone https://github.com/Felipe8617/gowebapp.git

RUN cd /gowebapp/gowebapp/app.go && build

EXPOSE 3000

CMD ["/gowebapp/gowebapp/app.go"]

FROM golang:1.19
RUN mkdir /gowebapp
WORKDIR /gowebapp
RUN export GO111MODULE=on
#RUN go get https://github.com/Felipe8617/gowebapp/gowebapp
RUN cd /gowebapp && git clone https://github.com/Felipe8617/gowebapp.git


RUN cd /gowebapp/gowebapp/app.go && go build

EXPOSE 3000

CMD ["/gowebapp/gowebapp/gowebapp"]