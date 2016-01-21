FROM golang:1.5

RUN apt-get update && apt-get install -y git

# Use go 1.5 vendoring
ENV GO15VENDOREXPERIMENT=1
ENV APPLICATION_NAME=workflow

# Install glide
WORKDIR $GOPATH/src/github.com/Masterminds
RUN git clone https://github.com/Masterminds/glide.git
RUN cd glide && git checkout 0.8.3 && make bootstrap && make install

RUN mkdir -p $GOPATH/src/github.com/3dsim/$APPLICATION_NAME

COPY . $GOPATH/src/github.com/3dsim/$APPLICATION_NAME

WORKDIR $GOPATH/src/github.com/3dsim/$APPLICATION_NAME
RUN glide install
RUN go test $(glide novendor)
# Must run after glide installs 3rd party (vendor) dependencies.  Installs binary on path.
RUN go install

ENTRYPOINT ["workflow"]
