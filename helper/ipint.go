package helper

import (
	"fmt"
	"math/big"
	"net"
)

func Ip2int(ip string) uint64 {
	ret := big.NewInt(0)
	ret.SetBytes(net.ParseIP(ip).To4())
	return ret.Uint64()
}

func Int2ip(ip uint64) string {
	return fmt.Sprintf("%d.%d.%d.%d",
		byte(ip>>24), byte(ip>>16), byte(ip>>8), byte(ip))
}
