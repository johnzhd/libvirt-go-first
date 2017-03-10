package main

import "github.com/rgbkrk/libvirt-go"
import "log"
import "fmt"

func main() {
	conn, err := libvirt.NewConnect("qemu:///193.168.14.69")
	if err != nil {
		log.Println(err)
	}
	defer conn.Close()

	doms, err := conn.ListAllDomains(libvirt.CONNECT_LIST_DOMAINS_ACTIVE)
	if err != nil {
		log.Println(err)
	}

	fmt.Printf("%d running domains:\n", len(doms))
	for _, dom := range doms {
		name, err := dom.GetName()
		if err == nil {
			fmt.Printf("  %s\n", name)
		}
		dom.Free()
	}
}
