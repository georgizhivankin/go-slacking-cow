FROM golang:1.7-alpine

# Set apps home directory.
ENV APP_DIR /go/src/github.com/georgizhivankin/go-slacking-cow

# Creates the application directory
RUN mkdir -p $APP_DIR

# Install make and CURL
RUN apk --update add make curl git

# Install glide
RUN curl https://glide.sh/get | sh

# Add sources.
COPY . $APP_DIR

# Define current working directory.
WORKDIR $APP_DIR

# Build go binary
RUN make

# expose 8010
expose 8010

# Now tell Docker what command to run when the container starts
CMD go-slacking-cow
