# gowebapp

#this is a basic go web app to test pipelines


#sonar.host.url=https://sonarcloud.io
# sonar.organization=felipe8617
# sonar.projectKey=Felipe8617_gowebapp
# sonar.projectName=gowebapp
# sonar.login=XXXXXXXXXXXXX
# sonar.sourceEncoding = UTF-8






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

#Asingnin working directory inside image
WORKDIR /gowebapp

#creating go.mod and go.sum files into images
COPY . .

RUN go mod download

#copying all .go files into image
COPY *.go .
#COPY templates/index.html ./

#creatting build
RUN go build -o /gowebapp

EXPOSE 6565

CMD [ "/gowebapp"]



	http.HandleFunc("/", Index)

	//creating the server
	fmt.Println("el servidor está corriendo en el puerto 3000")
	fmt.Println("Run server: http://localhost:3000")
	http.ListenAndServe("localhost:3000", nil)
