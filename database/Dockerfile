FROM postgres:14.4-alpine3.16

COPY --chown=nobody:nobody  create_tables.sql /docker-entrypoint-initdb.d/create_tables.sql
COPY --chown=nobody:nobody calories.csv /calories.csv

USER postgres

CMD ["postgres"]