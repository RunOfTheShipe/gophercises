# available versions : https://hub.docker.com/_/microsoft-mssql-server
ARG REGISTRY
FROM ${REGISTRY:-mcr.microsoft.com/}mssql/server:2019-CU11-ubuntu-20.04 AS base

ENV SA_PASSWORD=${GOBOX_SQL_PASSWORD:-Pass@word}
ENV ACCEPT_EULA=Y

USER root
RUN apt-get update
RUN apt-get upgrade -y

USER mssql
