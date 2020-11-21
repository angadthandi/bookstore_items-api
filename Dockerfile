FROM golang

# configure the repo url so we can configure our work directory
ENV REPO_URL=github.com/angadthandi/bookstore_items-api

# setup out $GOPATH
ENV GOPATH=/app

ENV APP_PATH=$GOAPTH/src/$REPO_URL

# /app/src/github.com/angadthandi/bookstore_items-api/src

# copy the entire source code from the curr directory to $WORKPATH
ENV WORKPATH=$APP_PATH/src
COPY src $WORKPATH
WORKDIR $WORKPATH

RUN go build -o items-api .
EXPOSE 8000

CMD ["./items-api"]