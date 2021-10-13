FROM bitwalker/alpine-elixir-phoenix:latest

# PG dependencies to start up db from container
RUN apk update && \
    apk add postgresql-client

# Set exposed ports
EXPOSE 4000

ENV APP_HOME /app
WORKDIR $APP_HOME

USER root

COPY .docker /tmp