FROM golang:latest
ADD hello /data/app/
WORKDIR /data/app/
ENV HTTPPORT=80
CMD ["/bin/bash", "-c", "./hello"]