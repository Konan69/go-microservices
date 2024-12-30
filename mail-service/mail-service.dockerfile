FROM alpine:latest

RUN mkdir /app

COPY mailerApp.exe /app

COPY templates /templates

CMD [ "/app/mailerApp.exe"]
