package storage

import (
	"fmt"
	"sort"
	"strings"

	"github.com/kassisol/twic/storage/driver"
)

type Initialize func(appPath string) (driver.Storager, error)

var initializers = make(map[string]Initialize)

func supportedDriver() string {
	drivers := make([]string, 0, len(initializers))

	for v := range initializers {
		drivers = append(drivers, string(v))
	}

	sort.Strings(drivers)

	return strings.Join(drivers, ",")
}

func NewDriver(driver, config string) (driver.Storager, error) {
	if init, exists := initializers[driver]; exists {
		return init(config)
	}

	return nil, fmt.Errorf("The Storage Driver: %s is not supported. Supported drivers are %s", driver, supportedDriver())
}

func RegisterDriver(driver string, init Initialize) {
	initializers[driver] = init
}
