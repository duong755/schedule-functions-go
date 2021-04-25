package database

import "os"

var CONNECTION_STRING string = os.Getenv("DATABASE_URL_FOR_VIEWER")
