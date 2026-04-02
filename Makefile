.PHONY: postgres adminer migrate goreddit-net 
# Declare the targets as phony to avoid conflicts with files of the same name
# i.e. do not generate files named postgres, adminer, or migrate when the commands are run

goreddit-net: # default network does not support service/container name resolution
	docker network inspect goreddit-net \ >/dev/null 2>&1 || docker network create goreddit-net

postgres:
	docker run --rm -ti \
	 --network goreddit-net \
	 --name postgres \
	 -e POSTGRES_PASSWORD=secret \
	 -p 5432:5432 \
	 postgres

adminer:
	docker run --rm -ti \
	 --network goreddit-net \
	 --name adminer \
	 -p 8080:8080 \
	 adminer

migrate:
	migrate -source file://migrations \
	 -database "postgres://postgres:secret@localhost:5432/postgres?sslmode=disable" up