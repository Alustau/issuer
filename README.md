Issuer 
=============

Issuer is a simple project to sync issues with Github through its Webhook 

## Table of Contents
- [Features](#features)
- [Installation](#installation)
- [Testing](#testing)

## Features

* List all issues from database
* Find a issue in database
* Save a issue in database (developing)

## Installation

1- Clone this project running this:
```shell
$ git clone https://github.com/alustau/issuer.git 
```
2- Import database with file in database/sql/issuer.sql:

3- Enter in project folder:
```shell
$ cd issuer
```
4- Build the app:
```shell
$ go install
```
5- Boot the app:
```shell
$ issuer
```
6- Go to http://0.0.0.0:8000/issue or http://0.0.0.0:8000/issue/1 

## Testing

Developing...
