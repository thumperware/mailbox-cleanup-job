FROM golang:1.21-alpine as build
# Setup private git repo
ARG GITHUB_USERNAME
ARG GITHUB_TOKEN
RUN apk add --no-cache git
RUN git config --global url."https://$GITHUB_USERNAME:$GITHUB_TOKEN@github.com/".insteadOf "https://git@github.com/"
RUN echo "machine github.com login ${GITHUB_USERNAME} password ${GITHUB_TOKEN}" > ~/.netrc
RUN chmod 600 ~/.netrc
RUN export GOPRIVATE=github.com/thumperq
# Build go project
WORKDIR /mailbox-cleanup-job
COPY . .
RUN go mod download && go mod verify
WORKDIR /mailbox-cleanup-job/cmd
RUN go build -o mailbox-cleanup-job
# Final image
FROM golang:1.21-alpine
COPY --from=build /mailbox-cleanup-job /mailbox-cleanup-job
WORKDIR /mailbox-cleanup-job/cmd
CMD [ "./mailbox-cleanup-job" ]

# Checkout https://vsupalov.com/build-docker-image-clone-private-repo-ssh-key/
# ENTRYPOINT ["tail", "-f", "/dev/null"] # for just keep the container up and running to be able to exec into it