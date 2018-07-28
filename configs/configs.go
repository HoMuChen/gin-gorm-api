package configs

import "os"

var API_HOST = os.Getenv("API_HOST")
var API_PORT = os.Getenv("API_PORT")

var PG_HOST = os.Getenv("PG_HOST")
var PG_PORT = os.Getenv("PG_PORT")
var PG_USER = os.Getenv("PG_USER")
var PG_PWD  = os.Getenv("PG_PWD")
var PG_DB   = os.Getenv("PG_DB")
