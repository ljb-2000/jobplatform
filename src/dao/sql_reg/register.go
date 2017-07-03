package sql_reg
import (
	"fmt"

	"github.com/astaxie/beego/orm"
	 _"github.com/go-sql-driver/mysql"
)

func init() {
	fmt.Println("Initializing DB registration.")
	orm.RegisterDriver("mysql", orm.DRMySQL)
	err := orm.RegisterDataBase("default", "mysql", "root:123456@tcp(10.110.18.107:30000)/k8s?charset=utf8")
	orm.RunSyncdb("default", true, true)
	if err != nil {
		fmt.Errorf("Error occurred on registering DB: %+v\n", err)
	}
}

