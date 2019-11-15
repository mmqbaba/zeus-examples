FROM alpine

COPY ./app /tmp/

WORKDIR /tmp

RUN chmod +x app

RUN mkdir /lib64 && ln -s /lib/libc.musl-x86_64.so.1 /lib64/ld-linux-x86-64.so.2

CMD ./app