FROM alpine
RUN apk update --no-cache && apk add --no-cache ca-certificates
ARG ENV
ARG SERVICE_NAME
ENV ENV=${ENV}
ENV SERVICE_NAME=${SERVICE_NAME}
ADD /bin/ /bin/
COPY ./etc/${ENV}-${SERVICE_NAME}-api.yaml /bin/
CMD /bin/${ENV}-${SERVICE_NAME}-api -f /bin/${ENV}-${SERVICE_NAME}-api.yaml