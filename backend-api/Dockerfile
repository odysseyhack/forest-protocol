FROM golang as builder

WORKDIR /go/src/gitlab.com/henkvanramshorst/forest
ADD . ./

ENV GOBIN=/go/src/gitlab.com/henkvanramshorst/forest
RUN go install ./cmd/...

ENTRYPOINT ["./backend-api"]