package tool

import (
	"net"
)

// GetNetworkCard 获取网卡信息：名字与ip地址
func GetNetworkCard() map[string][]net.Addr {
	var list = make(map[string][]net.Addr)
	ifaces, err := net.Interfaces()
	if err != nil {
		panic(err)
	}
	for _, i := range ifaces {
		name := i.Name
		addr, err := i.Addrs()
		if err != nil {
			panic(err)
		}
		if len(addr) > 0 {
			list[name] = addr
		}
	}
	return list
}
