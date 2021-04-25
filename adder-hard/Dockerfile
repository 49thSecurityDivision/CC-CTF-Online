FROM frolvlad/alpine-python3:latest

RUN apk add --no-cache socat

COPY addition.py /

CMD socat -v -T 15 TCP-LISTEN:8081,reuseaddr,fork EXEC:"python3 /addition.py"
