FROM golang:1.11
WORKDIR /go/src/github.com/anpryl/paymentsvc
COPY . .
RUN git clone https://github.com/vishnubob/wait-for-it.git
ENV INTEGRATION_TESTING=1
