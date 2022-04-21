package injection

import (
	"github.com/qingfeng777/owls/server/config"
	service "github.com/qingfeng777/owls/server/service/tidb_or_mysql"
	"github.com/qingfeng777/owls/server/service/tidb_or_mysql/admin"
	"github.com/qingfeng777/owls/server/service/tidb_or_mysql/auth"
	"github.com/qingfeng777/owls/server/service/tidb_or_mysql/auth/login_check"
	"github.com/qingfeng777/owls/server/service/tidb_or_mysql/checker"
	"github.com/qingfeng777/owls/server/service/tidb_or_mysql/dao"
	"github.com/qingfeng777/owls/server/service/tidb_or_mysql/db_info"
	"github.com/qingfeng777/owls/server/service/tidb_or_mysql/task"
)

func Injection()  {
	task.SetBackupDao(dao.BackupDAO)
	task.SetTaskDao(dao.Task)
	task.SetSubTaskDao(dao.SubTask)
	task.SetDbTools(db_info.DBTool)
	task.SetChecker(checker.Checker)
	checker.SetRuleStatusDao(dao.Rule)
	db_info.SetClusterDao(dao.Cluster)
	auth.SetLoginService(login_check.LoginService)
	service.SetClock(service.RealClock{})
	admin.SetAdminDao(dao.Admin)

	switch config.Conf.Role.From {
	case "conf":
		task.SetAuthTools(auth.ConfAuthService)
	case "net":
		task.SetAuthTools(auth.NetAuthService)
	case "admin":
		task.SetAuthTools(auth.AdminAuthService)
	case "mock":
		// MockInjection()
	default:
		task.SetAuthTools(auth.AdminAuthService)
	}
}