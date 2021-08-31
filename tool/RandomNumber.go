package tool

import (
	"crypto/md5"
	"encoding/hex"
	"math/rand"
	"strconv"
	"time"
)

// RandomNum 生成不会重复的字符串（根据时间戳和随机数生成数据）
func RandomNum() string {
	//根据时间戳生成随机数
	res := rand.Intn(time.Now().Second())
	h := md5.New()
	h.Write([]byte(strconv.Itoa(res + time.Now().Second())))
	resMd5 := hex.EncodeToString(h.Sum(nil))
	return resMd5
}
