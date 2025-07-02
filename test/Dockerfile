FROM golang:1.24@sha256:20a022e5112a144aa7b7aeb3f22ebf2cdaefcc4aac0d64e8deeee8cdc18b9c0f as build

ENV CGO_ENABLED=0

WORKDIR /app
COPY go* ./
RUN go mod download

COPY main.go main.go
COPY cmd cmd
COPY internal internal
RUN go build \
    -cover \
    -ldflags "-s -w -X main.version=1.2.3"

FROM ruby:3.3@sha256:6b7f4a1624ddd8d958277e16d7b86c47db48005fa42aadc430bde125cd795d06 as test

# Configure the shell to catch more errors loudly and early
SHELL ["/bin/bash", \
    "-o", "pipefail", \
    "-o", "errexit", \
    "-o", "nounset", \
    "-c"]

# Switch to non-root user for test context
ARG TEST_HOME="/home/test"
RUN     groupadd -r test \
    &&  useradd --no-log-init -r -g test -m -d $TEST_HOME test
USER test

WORKDIR "$TEST_HOME/app"

# Do not update Gemfile.lock
RUN bundle config set frozen true
# Install test dependencies
COPY ["test/Gemfile*", "./"]
RUN bundle install

# "Install" tool
COPY --from=build /app/template-goapp /bin/

RUN mkdir "$TEST_HOME/app/coverage"
ENV GOCOVERDIR=$TEST_HOME/app/coverage

# Run tests
COPY --chown=test:test test ./
ARG verbose=false
ENV VERBOSE_TEST_OUTPUT="$verbose"
ENTRYPOINT ["cucumber"]
CMD [ \
    "--publish-quiet", \
    "--tags", "not @pending" \
]
