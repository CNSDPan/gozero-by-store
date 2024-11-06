package mysqltest

import (
	"fmt"
	sqlsStore "store/app/store/model/sqls"
	"testing"
)

func TestGetInitChatLog(t *testing.T) {
	//items, err := storeModel.ChatLogMgr.InitChatLog(
	//	sqlsStore.NewPage(100, 0),
	//	1837401868671213568,
	//)
	items, err := storeModel.StoresMemberMgr.InitChatLog(
		sqlsStore.NewPage(10, 0),
		1853389743766200320,
	)
	if err != nil {
		fmt.Printf("err %v", err)
		return
	}
	fmt.Printf("列表： %+v", items)
}
