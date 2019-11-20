FROM alpine

COPY ./hellodemo_server /tmp/

WORKDIR /tmp

RUN chmod +x hellodemo_server

RUN mkdir /lib64 && ln -s /lib/libc.musl-x86_64.so.1 /lib64/ld-linux-x86-64.so.2

CMD ./hellodemo_server --port 9090 --apiPort 8081 --apiInterface 127.0.0.1 --confEntryPath /zeus/hellodemo

