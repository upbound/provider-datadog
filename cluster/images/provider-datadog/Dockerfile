FROM gcr.io/distroless/static@sha256:9197324ba51d9cd071af8505989365c006adf9d6d2067eada25aef00abbb5278

ARG TARGETOS
ARG TARGETARCH

ENV USER_ID=65532

USER ${USER_ID}
EXPOSE 8080

ADD "bin/${TARGETOS}_${TARGETARCH}/provider" /usr/local/bin/provider

ENTRYPOINT ["provider"]
