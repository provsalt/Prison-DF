FROM golang

RUN mkdir /server

COPY . /server

WORKDIR /server

RUN go build -o prisons .

RUN chmod +x prisons

EXPOSE 19132/udp

CMD [ "/server/prisons" ]