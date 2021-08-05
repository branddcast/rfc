FROM registry.access.redhat.com/ubi8/ubi:latest as ubi8redhat.acces

RUN dnf install tar wget go-toolset -y

# Config GO
ENV PATH="/usr/local/go/bin:$PATH" \
    GOPATH="/go"
RUN mkdir -p /go/src/ /go/bin/ /go/pkg/

# workspace directory
WORKDIR /go/src/validadorRFC-go

# copy source code
COPY . .

# Config GOPROXY

# golang config dependencys
RUN go mod init validadorRFC-go
RUN go mod tidy

# build executable
RUN go build -o ./bin/validadorRFC-go .

##################################################    multistage image
FROM registry.access.redhat.com/ubi8/ubi-minimal:latest

#copy form BUILDER API compile
COPY --from=ubi8redhat.acces /go/src/validadorRFC-go/bin ./bin

# change workdir
WORKDIR /bin

# variables de ambiente

# set entrypoint
ENTRYPOINT [ "validadorRFC-go" ]
