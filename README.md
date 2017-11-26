an api. in go.

Run `docker-compose up` to get the app and DB running.

To browse the DB:

* if you have posrgres installed: `psql -h 0.0.0.0 -p 32768 -U admin -d shakespeare` (password 'admin')
* from inside the docker container: `docker-compose exec psql psql -U admin -d shakespeare`
