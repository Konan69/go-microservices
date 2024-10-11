
FROM alpine:latest

RUN mkdir /app

COPY brokerApp.exe /app

CMD [ "app/brokerApp.exe" ] 