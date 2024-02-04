FROM ubuntu:latest

COPY adventure ./
COPY gopher.json ./
COPY story.html ./

EXPOSE 8080

CMD ["/adventure"]