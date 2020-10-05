package database

import (
	"github.com/go-xorm/xorm"
)

// DBClient is an instance of IDb, and been injected by main and test function
var DBClient *xorm.Engine