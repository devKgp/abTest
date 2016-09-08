package servicetest

import (
	"github.com/devkgp/abTest/src/sqldb"
	gk "github.com/onsi/ginkgo"
	gm "github.com/onsi/gomega"
	"testing"
)

func TestXYZ(t *testing.T) {
	gm.RegisterFailHandler(gk.Fail)
	gk.RunSpecs(t, "Service Suite")
}

var _ = gk.Describe("Test mysql impl", func() {
	var err *sqldb.SqlDbError
	var dbObj sqldb.SqlDbInterface
	conf := new(sqldb.Config)

	gk.Context("put valid driver name", func() {
		// fill valid driver, invalid db details
		conf.DriverName = sqldb.MYSQL // set driver name as mysql
		conf.Username = "root"
		conf.Password = "root"
		conf.Host = "localhost"
		conf.Port = "3306"
		conf.Dbname = "invalid"
		conf.Timezone = "Local"
		conf.MaxOpenCon = 2
		conf.MaxIdleCon = 1
		gk.It("impl should not return error", func() {
			dbObj, err = sqldb.Get(conf)
			gm.Expect(err).To(gm.BeNil())
		})
	})
})
