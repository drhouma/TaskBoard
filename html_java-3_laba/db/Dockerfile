FROM postgres:latest

ENV POSTGRES_USER postgres
ENV POSTGRES_PASSWORD 123
ENV POSTGRES_DB html_comments

EXPOSE 5432

COPY ./init.sql /docker-entrypoint-initdb.d/
