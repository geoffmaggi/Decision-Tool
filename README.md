# Decision Tool ![Version](https://img.shields.io/badge/version-1.1.0-blue.svg) [![License: MIT](https://img.shields.io/badge/license-MIT-blue.svg)](https://opensource.org/licenses/MIT) [![Go Report Card](https://goreportcard.com/badge/github.com/geoffmaggi/Decision-Tool)](https://goreportcard.com/report/github.com/geoffmaggi/Decision-Tool) [![Travis CI](https://travis-ci.org/geoffmaggi/Decision-Tool.svg?branch=master)](https://travis-ci.org/geoffmaggi/Decision-Tool#) ![Tests](https://img.shields.io/badge/tests-128%2F128-brightgreen.svg)

The Decision Tool is a secure, cross-platform and concurrent full stack web application to aid in the decision making process.

## Dependencies

 - [Golang](http://golang.org)
 - [MySql](https://www.mysql.com/) or [MariaDB](https://mariadb.org/)

## Installing and building

1. First make sure go is installed and configured
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

## Testing

To test the code just run following 

```
./run_tests.sh
```

This will run all tests and ouput the file "system.html" which can be viewed in the browser to see code coverage.

## Deploying

1. Alter deploy_regex.sh to point to the correct web address.
2. Run the following to produce a zipped file with minified static files, a compiled binary and the configuration files.
```
./release.sh
```
3. Copy the contents of the zip onto the production server and update the configuration files.
4. Run the binary to initialize the server.

# API

For all API below if the header accept json then json objects are sent as mentioned below,
if the accept header wants html then the same objects are passed to static file that
runs a javascript file with that name found in `static/` folder.
eg : Requestion html from `GET /persons` will give an html reply that will run
the file `static/persons_list.js`

## person

These are the things we can do with `person`, under this table is description
of some of the things we need to send or the things we receive back.


| Description            | URL                          | Method | Wants   | Gives               |
|------------------------|------------------------------|--------|---------|---------------------|
| Create a user          | /person                      | POST   | p1      | p2                  |
| Get all persons        | /persons                     | GET    | nothing | array of p2         |
| Get person info        | /person/:person_id/info      | GET    | nothing | p2                  |
| Get person's decisions | /person/:person_id/decisions | GET    | nothing | a persons decisions |
| Delete a person        | /person/:person_id           | DELETE | nothing | s1                  |
| Update a user          | /person/:person_id           | PUT    | p1      | p2                  |

Wants/Gives

```
p1 = {"email":<str>,"pw_hash":<str>,"name_first":<str>,"name_last":<str>}
p2 = {"person_id":<int>, "email":<str>,"pw_hash":<str>,"name_first":<str>,"name_last":<str>}
s1 = {"result": "deleted"}
```
What we actually return is `{"somehting": object}` for example :
- Get person will return

```
{"person": p1 object}
```

- Get persons will return

```
{"persons": array of p1 object}
```

and so on..


## decision

These are the things we can do with `decision`. Things like creating a
decision, configuring it's criterions, creating ballots, doing voting on the
ballots, and gather statistics.

| Description                   | URL                                                                                                       | Method | Wants   | Gives       |
|-------------------------------|-----------------------------------------------------------------------------------------------------------|--------|---------|-------------|
| Create a decision             | /decision                                                                                                 | POST   | d1      | d2          |
| Get all decisions             | /decisions                                                                                                | GET    | nothing | array of d2 |
| Decision update               | /decision/:decision_id                                                                                    | PUT    | d1      | d2          |
| Get a decision info           | /decision/:decision_id/info                                                                               | GET    | nothing | d2          |
| Get a decision statistics     | /decision/:decision_id/stats                                                                              | GET    | nothing | s1          |
| Delete a decision             | /decision/:decision_id                                                                                    | DELETE | nothing | r1          |
| Create a ballot for decision  | /decision/:decision_id/ballot                                                                             | POST   | b1      | b2          |
| Create a ballot without inv.  | /decision/:decision_id/ballot_silent                                                                      | POST   | b1      | b2          |
| List ballots in a decision    | /decision/:decision_id/ballots                                                                            | GET    | nothing | array of b2 |
| Show a ballot info            | /decision/:decision_id/ballot/:ballot_id/info                                                             | GET    | nothing | b2          |
| Delete a ballot               | /decision/:decision_id/ballot/:ballot_id                                                                  | DELETE | nothing | r1          |
| Update a ballot               | /decision/:decision_id/ballot/:ballot_id                                                                  | PUT    | b1      | b2          |
| Create a decision criterion   | /decision/:decision_id/criterion                                                                          | POST   | c1      | c2          |
| List all criterions           | /decision/:decision_id/criterions                                                                         | GET    | nothing | array of c2 |
| Get a criterion info          | /decision/:decision_id/criterion/:criterion_id/info                                                       | GET    | nothing | c2          |
| Update a criterion            | /decision/:decision_id/criterion/:criterion_id                                                            | PUT    | c1      | c2          |
| Delete a criterion            | /decision/:decision_id/criterion/:criterion_id                                                            | DELETE | nothing | r1          |
| Ballot votes in a criterion   | /decision/:decision_id/ballot/:ballot_id/alternative/:alternative_id/criterion/:criterion_id/vote/:weight | GET    | nothing | v2          |
| Update Ballot vote            | /decision/:decision_id/ballot/:ballot_id/alternative/:alternative_id/criterion/:criterion_id/vote/:weight | PUT    | nothing | v2          |
| Show ballot votes             | /decision/:decision_id/ballot/:ballot_id/votes                                                            | GET    | nothing | array of v2 |
| Delete a vote in ballot       | /decision/:decision_id/ballot/:ballot_id/alternative/:alternative_id/criterion/:criterion_id/vote         | DELETE | nothing | r1          |
| Create an alternative         | /decision/:decision_id/alternative                                                                        | POST   | a1      | a1          |
| Show all alternatives         | /decision/:decision_id/alternatives                                                                       | GET    | nothing | array of a1 |
| Show info of one alternative  | /decision/:decision_id/alternative/:alternative_id/info                                                   | GET    | nothing | a1          |
| Delete an alternative         | /decision/:decision_id/alternative/:alternative_id                                                        | DELETE | nothing | r1          |
| Update an alternative         | /decision/:decision_id/alternative/:alternative_id                                                        | PUT    | a1      | a1          |
| More ballot information       | /decision/:decision_id/ballot/:ballot_id                                                                  | GET    | nothing | b3          |
| Ballot rate a criterion       | /decision/:decision_id/ballot/:ballot_id/criterion/:criterion_id/vote/:rating                             | GET    | nothing | k1          |
| Get criterion ratings         | /decision/:decision_id/criterion/:criterion_id/votes                                                      | GET    | nothing | array of k1 |
| Delete a rating on an altern. | /decision/:decision_id/ballot/:ballot_id/criterion/:criterion_id/vote                                     | DELETE | nothign | r1          |
| Update criterion rating       | /decision/:decision_id/ballot/:ballot_id/criterion/:criterion_id/vote/:rating                             | PUT    | nothing | k1          |
| Decision duplicate            | /decision/:decision_id/duplicate                                                                          | GET    | nothing | d2          |

Wants/Gives

```
a1 = {"name": <str>, "description":<optional-str>, "cost":<optiona-int> }
d1 = {"person_id":<int>, "name":<str>, "description":<str>, "stage":<int>, "criterion_vote_style":<str>, "alternative_vote_style":<str>, "client_settings":<str>}
d2 = {"decision_id":<int>, "person_id":<int>, "name":<str>, "description":<str>, "stage":<int>, "criterion_vote_style":<str>, "alternative_vote_style":<str>, "client_settings":<str>}
s1 = "undecided yet"
b1 = {"name":<str>, "email":<str>, "sent":<bool-opt>}
b2 = {"ballot_id":<int>, "decision_id":<int>, "secret":<str>, "name":<str>, "email":<str>, "sent":<bool-opt>}
b3 = {"ballot": b1, "ratings": array of v2}
c1 = {"name":<str>, "description":<optional-str>}
c2 = {"criterion_id":<int>, "description":<optional-str>, "decision_id":<int>, "name":<str>}
v1 = {"weight":<int>}
v2 = {"criterion_id":<int>, "ballot_id":<int>, "weight":<int>}
r1 = {"result": "deleted"}
k1 = {"criterion_id": <int>, "ballot_id":<int>, "rating":<int> }
```

What we actually return is `{"somehting": object}` for example :
- Get decision will return

```
{"decision": d1 object}
```

- Get decisions will return

```
{"decisions": array of d1 object}
```

and so on..


# Authentication

Authentication is implemented as middlewares applied to routes, currently they are applied
to 0 routes to make it easier for the front-end to start writing their code. The login
you get a cookie and this cookie is checked..etc

| Description                             | URL         | Method | Wants   | Gives   |
|-----------------------------------------|-------------|--------|---------|---------|
| Login                                   | /login      | POST   | l1      | l2      |
| Logout                                  | /logout     | GET    | Nothing | Nothing |
| Get person_id of current logged in user | /whoami     | GET    | Nothing | g1      |

Wants/Gives

```
l1 = {"email":<str>, "password":<str>}
l2 = {"error":<str>} or {"status":<str>}
g1 = {"person_id": <int>} or {"error": <str> }
```

# Permissions

In here I list the permissions and who can do what. They are three types of
permissions.

- Admin : Only authenticated admin is allowed
- All   : Only authenticated persons (admin or facilitator)
- None  : Anyone

This list contains the routes that have permissions on them, anything not
listed in here is considered to have a `None` permission.

| Action            | URL                      | Who is allowed |
|-------------------|--------------------------|----------------|
| Person Creation   | /person                  | admin          |
| Person Delete     | /person/:pid             | admin          |
| Person Update     | /person/:pid             | all            |
| Decision Creation | /decision/:did           | all            |
| Decision Delete   | /decision/:did           | all            |
| Decision Update   | /decision/:did           | all            |
| Decision Dupl.    | /decision/:did/duplicate | all            |
| Alt. Creation     |                          | all            |
| Alt. Delete       |                          | all            |
| Alt. Update       |                          | all            |
| Blt. Creation     |                          |                |
| Blt. Delete       |                          |                |
| Blt. Update       |                          |                |
| Crt. Creation     |                          |                |
| Crt. Delete       |                          |                |
| Crt. Update       |                          |                |
| Vt. Delete        |                          |                |
| Vt. Update        |                          |                |
| Rt. Delete        |                          |                |
| Rt. Update        |                          |                |
