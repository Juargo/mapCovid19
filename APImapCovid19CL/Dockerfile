FROM golang:alpine as builder

COPY . $GOPATH/src/github.com/Juargo/mapCovid19CL/
WORKDIR   $GOPATH/src/github.com/Juargo/mapCovid19CL/
RUN adduser -D -g '' APImapCovid19CL
RUN apk add git
RUN go get -d -v

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -installsuffix cgo -ldflags="-w -s" -o /go/bin/APImapCovid19CL
RUN apk update && apk add ca-certificates && rm -rf /var/cache/apk/* \
  mkdir /usr/local/share/ca-certificates/extra
COPY .docker/other/cert_Intertrials-CA.crt /usr/local/share/ca-certificates/extra
RUN update-ca-certificates
# RUN apk update && apk add ca-certificates && rm -rf /var/cache/apk/*
# COPY ./mycert.crt /usr/local/share/ca-certificates/mycert.crt
# RUN update-ca-certificates

FROM scratch

COPY --from=builder /etc/passwd /etc/passwd
COPY --from=builder /go/bin/APImapCovid19CL /go/bin/APImapCovid19CL
# RUN apk --no-cache add ca-certificates 
# ADD  ca-certificates.crt /etc/ssl/certs/

USER APImapCovid19CL

CMD ["/go/bin/APImapCovid19CL"]
