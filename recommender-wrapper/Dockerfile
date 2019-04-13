FROM golang:1.9-stretch

WORKDIR /app


COPY ./bin/api .
COPY ./config/dev/config.toml ./config/dev/config.toml

# COPY PYTHON STUFFS

ENTRYPOINT ["./api", "cfg=config/dev/config.toml"]