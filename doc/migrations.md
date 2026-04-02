Using `go migrate` as cli
https://github.com/golang-migrate/migrate -> Releases

To fix could not verify free of malware: `xattr -d com.apple.quarantine /Users/prantoran/go/bin/migrate`

Expected status of created table in adminer after `make migrate`:

| Table | Engine | Data Length? | Index Length? | Data Free | Auto Increment | Rows? | Comment? |
| :--- | :--- | :---: | :---: | :---: | :---: | :---: | :--- |
| **comments** | table | 8,192 | 8,192 | ? | ? | ? | |
| **posts** | table | 8,192 | 8,192 | ? | ? | ? | |
| **schema_migrations** | table | 8,192 | 16,384 | ? | ? | ? | |
| **threads** | table | 8,192 | 8,192 | ? | ? | ? | |
| **4 in total** | | **en_US.utf8** | **32,768** | **40,960** | **0** | | |