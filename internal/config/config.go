package config

import "os"

var CONNECTION_URL = os.Getenv("MONGODB_URL")
var PORT = os.Getenv("PORT")
