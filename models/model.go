package models

import (
	"os"

	"github.com/tesh254/emania/db"
)

var server = os.Getenv("DATABASE")

var databaseName = os.Getenv("DATABASE_NAME")

var dbConnect = db.NewConnection(server, databaseName)