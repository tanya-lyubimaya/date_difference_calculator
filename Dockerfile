FROM golang:1.20-rc
EXPOSE 8080
RUN mkdir /application && apt install make
COPY . /application
WORKDIR /application
CMD ["make", "run"]