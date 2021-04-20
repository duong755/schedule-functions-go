package database

import "os"

var CONNECTION_STRING string = os.Getenv("CONNECTION_STRING")
