FROM golang as build-env
ARG TAG

WORKDIR /src
COPY . .
RUN make TAG=$TAG GOMOD="-mod=vendor" build

FROM registry.access.redhat.com/ubi7/ubi-minimal:latest AS final

ENV OPERATOR=/usr/local/bin/datadog-operator \
    USER_UID=1001 \
    USER_NAME=datadog-operator

### Required OpenShift Labels
LABEL name="Datadog Operator" \
      vendor="Datadog" \
      version="$TAG" \
      release="1" \
      summary="Datadog Operator" \
      description="Datadog provides a modern monitoring and analytics platform. Gather metrics, logs and traces for full observability of your Kubernetes cluster with Datadog Operator."

# Required Licenses
COPY licenses /licenses

# install operator binary
COPY --from=build-env /src/controller ${OPERATOR}

COPY --from=build-env /src/build/bin /usr/local/bin
RUN  /usr/local/bin/user_setup

ENTRYPOINT ["/usr/local/bin/entrypoint"]

USER ${USER_UID}