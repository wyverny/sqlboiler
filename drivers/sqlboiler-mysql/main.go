package main

import (
	"github.com/wyverny/sqlboiler/v4/drivers"
	"github.com/wyverny/sqlboiler/v4/drivers/sqlboiler-mysql/driver"
)

func main() {
	drivers.DriverMain(&driver.MySQLDriver{})
}
