//File  : main.go
//Author: duanhaobin
//Date  : 2020/5/20

package main

import (
	"ginEssential-hb/common"
	"ginEssential-hb/router"

	"github.com/gin-gonic/gin"
)

func main() {
	db := common.InitDB()
	defer db.Close()

	r := gin.Default()
	r = router.CollectRoute(r)
	panic(r.Run())
}
