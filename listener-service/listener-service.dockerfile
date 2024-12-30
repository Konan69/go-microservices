 
FROM alpine:latest

RUN mkdir /app

COPY listenerApp.exe /app

CMD [ "app/listenerApp.exe" ] 