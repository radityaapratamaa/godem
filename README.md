# godem Backend
godem Backend with golang

## How to start
- copy config file with
```bash
cp files/development.yaml.example files/development.yaml
```

- update the yaml configuration file for your environment
- update the migration db at `Makefile` on `migration-up`

please run this command to running the migrations file
```bash
make install-goose
make migration-up
```

and then, you can run the apps with this command
```bash
make run
```

if you want to run the unit test, just type this command
```bash
make test
```