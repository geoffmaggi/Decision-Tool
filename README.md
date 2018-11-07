# Decision Tool ![Version](https://img.shields.io/badge/version-1.1.0-blue.svg) [![License: MIT](https://img.shields.io/badge/license-MIT-blue.svg)](https://opensource.org/licenses/MIT) [![Go Report Card](https://goreportcard.com/badge/github.com/geoffmaggi/Decision-Tool)](https://goreportcard.com/report/github.com/geoffmaggi/Decision-Tool) [![Travis CI](https://travis-ci.org/geoffmaggi/Decision-Tool.svg?branch=master)](https://travis-ci.org/geoffmaggi/Decision-Tool#) ![Tests](https://img.shields.io/badge/tests-128%2F128-brightgreen.svg)

The Decision Tool is a secure, cross-platform and concurrent full stack web application to aid in the decision making process.

## Dependencies

 - [Golang](http://golang.org)
 - [MySql](https://www.mysql.com/) or [MariaDB](https://mariadb.org/)

## Installing and building

1. Install and configure Go and your Database of choice
2. Edit \_config.conf and save it as **config.conf**
3. Edit \_smtp.conf and save it as **smtp.conf**

Run the following
```
# Get Dependencies
go get
# Build the project
go build
# Run the code
./Decision_Tool
```
4. Visit localhost:9999 in your browser

## Testing

To test the code just run following 

```
./run_tests.sh
```

This will run all tests and ouput the file "system.html" which can be viewed in the browser to see code coverage.

## Deploying

1. Alter deploy_regex.sh to point to the correct web address
2. Run the following to produce a zipped file with minified static files, a compiled binary and the configuration files
```
./release.sh
```
3. Copy the contents of the zip onto the production server and update the configuration files
4. Run the binary to initialize the server
