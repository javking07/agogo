FROM debian
ADD hello-serv /app/hello-serv
RUN apt-get update && chmod 777 /app/hello-serv
EXPOSE 8080:8080
ENTRYPOINT ["/app/hello-serv"]