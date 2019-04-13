FROM unchain/alpine-glibc-dockerize

WORKDIR /app

COPY ./bin/api .
COPY ./config/dev/config.toml ./config/dev/config.toml

ENTRYPOINT ["./api", "cfg=config/dev/config.toml"]