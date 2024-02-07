FROM ubuntu:latest

RUN apt update && apt install -y ca-certificates

COPY adventure ./
COPY gopher.json ./
COPY story.html ./

EXPOSE 8080

CMD ["/adventure"]