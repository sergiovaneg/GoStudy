package main

import (
	"crypto/md5"
	"fmt"
	"strconv"
	"strings"
)

const ID = "ffykfhsq"

func findPasswordA(doorId string) (passwd string) {
	for idx := 0; len(passwd) < 8; idx++ {
		hash := md5.Sum([]byte(doorId + strconv.Itoa(idx)))
		prefix := fmt.Sprintf("%x", hash[:3])
		if prefix[:5] == "00000" {
			passwd += prefix[5:]
		}
	}
	return
}

func findPasswordB(doorId string) (passwd string) {
	passwd = "        "
	for idx := 0; strings.Contains(passwd, " "); idx++ {
		hash := md5.Sum([]byte(doorId + strconv.Itoa(idx)))
		prefix := fmt.Sprintf("%x", hash[:4])
		if prefix[:5] == "00000" {
			loc, err := strconv.Atoi(prefix[5:6])
			if err == nil && loc < 8 && passwd[loc] == ' ' {
				aux := []byte(passwd)
				aux[loc] = prefix[6]
				passwd = string(aux)
			}
		}
	}
	return
}

func main() {
	println(findPasswordA(ID))
	println(findPasswordB(ID))
}
