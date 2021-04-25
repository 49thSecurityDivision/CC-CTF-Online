FROM python:latest
RUN mkdir /tmp/website
COPY . /tmp/website
WORKDIR /tmp/website
CMD python -m http.server 8080 
