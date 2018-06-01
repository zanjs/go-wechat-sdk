package account


// 从数据库初始化所有应用账号对象，存进内存中

var Account map[int]map[string]string

func init() {

	// 数据库读取数据
	Account = map[int]map[string]string{
		1: map[string]string{
			"app_id": "",
		},
	}
}

func GetAccountInfo(accountid int) map[string]string {
	return Account[accountid]
}