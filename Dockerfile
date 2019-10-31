FROM golang:1.13-alpine

RUN mkdir -p /src
WORKDIR /src

RUN apk -v --update --no-cache add \
        curl \
        git \
        python \
        py-pip \
        groff \
        less \
        mailcap \
        gcc \
        libc-dev \
        bash && \
        pip install --upgrade --no-cache awscli s3cmd python-magic && \
        apk -v --purge del py-pip && \
        rm /var/cache/apk/* || true

COPY go.mod go.sum ./
RUN go mod download

COPY  . .
RUN go test ./tests/...

RUN go build -o /bin/pocketethereum ./app/app.go
ENTRYPOINT ["pocketethereum"]