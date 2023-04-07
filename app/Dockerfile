FROM golang:1.19.3-alpine AS builder

RUN mkdir /app

COPY . /app

WORKDIR /app

RUN go build -o ocra-app /app/main.go


FROM alpine

RUN mkdir /app

COPY --from=builder /app/ocra-app /app
COPY --from=builder /app/service/mail/template_email.html /app/service/mail/template_email.html

CMD [ "/app/ocra-app" ]