package sqldb

import (
	"fmt"
)

// Get() - Creates, initializes and returns the mysql instance based on the given config
func Get(conf *Config) (ret SqlDbInterface, err *SqlDbError) {
	fmt.Println(conf)
	if conf.DriverName == MYSQL {
		ret = new(mysqlDriver)
		err = ret.init(conf)
	} else {
		err = getErrObj(ERR_NO_DRIVER, conf.DriverName+" is not supported")
	}
	// return
	return ret, err
}
