package main

import (
	"douyin/shared/utils"
	"fmt"
	"os"
	"time"

	"github.com/bytedance/sonic"
)

func main() {
	claims, _ := utils.ParseToken(os.Args[1])
	id := claims.Id
	t := claims.ExpiresAt
	fmt.Println(sonic.MarshalString(claims))
	fmt.Println(id, time.Unix(t, 0))
}
