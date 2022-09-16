package main

import (
	"github.com/wyverny/sqlboiler/v4/drivers"
	"github.com/wyverny/sqlboiler/v4/drivers/sqlboiler-sqlite3/driver"
)

func main() {
	drivers.DriverMain(&driver.SQLiteDriver{})
}
