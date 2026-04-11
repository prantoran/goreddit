- Layer for handling sessions: `go get github.com/alexedwards/scs/v2`
- Store sessions: `go get github.com/alexedwards/scs/postgresstore`
- Create new db table to store sessions:
```sql
CREATE TABLE sessions (
	token TEXT PRIMARY KEY,
	data BYTEA NOT NULL,
	expiry TIMESTAMPTZ NOT NULL
);

CREATE INDEX sessions_expiry_idx ON sessions (expiry);
```
- Create and run migrations: `make migrate`