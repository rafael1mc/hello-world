rundb:
	docker container run --rm -p 5432:5432 -e POSTGRES_PASSWORD=yourpassword -e POSTGRES_USER=postgres -e POSTGRES_DB=yourdatabase postgres