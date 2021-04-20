package main

import (
	"encoding/hex"
	"fmt"
	"math/rand"
	"reflect"
	"strings"
	"time"

	uuid "github.com/satori/go.uuid"
)

func main() {
	u1 := uuid.Must(uuid.NewV4())
	fmt.Println(u1, reflect.TypeOf(u1))
	u2 := u1.String()
	fmt.Println(u2, reflect.TypeOf(u2))

	localDeviceIDPrefix := "44030500002000"
	gb1 := genGBID(localDeviceIDPrefix, int32(60))
	fmt.Println(gb1)

	s := "A50F01021F0000D6"
	sx, err := hex.DecodeString(s)
	fmt.Println("s: ", sx)
	fmt.Println("err: ", err)
	fmt.Println("len(s): ", len(s))
}

func genGBID(localDeviceIDPrefix string, port int32) string {
	localDeviceIDPrefix = fmt.Sprintf("%.14s", localDeviceIDPrefix)
	if len(localDeviceIDPrefix) < 14 {
		localDeviceIDPrefix = fmt.Sprintf("%s%s", localDeviceIDPrefix, strings.Repeat("0", 14-len(localDeviceIDPrefix)))
	}
	r1 := rand.New(rand.NewSource(time.Now().UnixNano())).Intn(1000000)
	return fmt.Sprintf("%s%06d", fmt.Sprintf("%014.14s", localDeviceIDPrefix), int32(r1)+port)
}
