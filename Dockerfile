FROM golang:1.9-alpine

# Set apps home directory.
ENV APP_DIR /go/src/github.com/georgizhivankin/go-slacking-cow

# Creates the application directory
RUN mkdir -p $APP_DIR

# Add sources.
COPY . $APP_DIR

# Define current working directory.
WORKDIR $APP_DIR

# Build the go binary
RUN set -ex \
    && apk update \
    && apk upgrade \
    && apk add --no-cache --update --virtual .build-deps curl git build-base \
    && curl https://glide.sh/get | sh \
    && apk add git \
    && apk add make \
    && make \
        # Clean up
    && apk del .build-deps \
    && rm -rf /var/cache/apk/*

# expose 8080
expose 8080

# Now tell Docker what command to run when the container starts
CMD go-slacking-cow
