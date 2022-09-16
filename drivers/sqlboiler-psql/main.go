package main

import (
	"github.com/wyverny/sqlboiler/v4/drivers"
	"github.com/wyverny/sqlboiler/v4/drivers/sqlboiler-psql/driver"
)

func main() {
	drivers.DriverMain(&driver.PostgresDriver{})
}
