FROM ubuntu:20.04 as chroot

RUN /usr/sbin/useradd --no-create-home -u 1000 user

FROM gcr.io/kctf-docker/challenge@sha256:6dd60da626bc43bf3175d9d7436006db5acc7710d5d1b7006ab53e718fe51e40

RUN apt update && apt install libev-dev python3 python3-pip python3-dev gcc -y

RUN pip3 install falcon bjoern

COPY nsjail.cfg /home/user/
COPY swagger* /home/user/
COPY webserver.py /home/user

CMD /home/user/webserver.py
