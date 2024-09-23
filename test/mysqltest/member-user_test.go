package mysqltest

import (
	"fmt"
	sqlsStore "store/app/store/model/sqls"
	"testing"
)

func TestGetMemberUser(t *testing.T) {
	items, err := storeModel.StoresMemberMgr.SelectPageApi(
		sqlsStore.NewPage(100, 0),
		storeModel.StoresMemberMgr.WithStoreId(1837056807609659392),
	)
	if err != nil {
		fmt.Printf("err %v", err)
		return
	}
	fmt.Printf("列表： %+v", items)
}
