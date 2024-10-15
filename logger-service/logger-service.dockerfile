FROM alpine:latest

RUN mkdir /app

COPY loggerServiceApp.exe /app

CMD [ "/app/loggerServiceApp.exe"]