# go-etl-csv

Sample API to digest incoming csv message and load data to database.

# API Documentation

See [Redoc / Openapi Documentation](https://remyranger.github.io/go-etl-csv/) for 'github page' hosted documentation.

# required

* MySQL database
* golang 1.20
* task & npx (to generate Openapi documention)

# config

Copy /config/ticketstore.yaml into your $home/.config/api-go/

Copy /config/certs/ into $home/.config/api-go/

Update config file with your paths and db params. Init database tables with sql script in /config/initDb.sql

# generate documentation

To build all OpenAPI bundles to one unique file and generate static html file documentation, use :

`$ task generate_docs`

File generated in /doc/dist is used by Golang for interface generation on build.

# generate interfaces

To build interfaces with 'go generate' :

`$ task generate_interfaces`

# run API

To build and run API :

`$ task cmd:run`

# Postman

Try it with postman. Import collection available in dir /postman .