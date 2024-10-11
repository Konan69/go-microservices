
FROM alpine:latest

RUN mkdir /app

COPY authApp.exe /app

CMD [ "app/authApp.exe" ] 