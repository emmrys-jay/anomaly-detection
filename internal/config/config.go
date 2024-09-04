package config

import "os"

var CONNECTION_URL = os.Getenv("MONGODB_URL")
var PORT = os.Getenv("PORT")
var NEW_URL=os.Getenv("NEW_URL")