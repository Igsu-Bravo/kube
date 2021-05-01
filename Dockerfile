# Use 1.16.3 imagewith Go preinstalled
FROM golang:1.16.3-buster

# Install Bee globally in container. Used for live-reload while developing
RUN go get -u github.com/beego/bee

# variables available during runtime
ENV GO111MODULE=on
ENV GOFLAGS=-mod=vendor
ENV APP_USER app
ENV APP_HOME /go/src/kube

# Variables available during build time
# ARG something something

# Create user, home directory and app directory inside container
RUN groupadd app && useradd -m -l --gid $GROUP_ID $APP_USER
RUN mkdir -p $APP_HOME && chown -R $APP_USER:$APP_USER $APP_HOME

USER $APP_USER

WORKDIR $APP_HOME

# Tell Docker that port 3000 is interesting
EXPOSE 3000

# Start application with 🐝
CMD ["bee", "run"]