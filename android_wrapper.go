//go:build android

package sauXKit

import (
	c "github.com/saudevteam/sauxkit/controller"
	"github.com/saudevteam/sauxkit/dns"
)

type DialerController interface {
	ProtectFd(int) bool
}

func InitDns(controller DialerController, server string) {
	dns.InitDns(server, func(fd uintptr) {
		controller.ProtectFd(int(fd))
	})
}

func ResetDns() {
	dns.ResetDns()
}

func RegisterDialerController(controller DialerController) {
	c.RegisterDialerController(func(fd uintptr) {
		controller.ProtectFd(int(fd))
	})
}

func RegisterListenerController(controller DialerController) {
	c.RegisterListenerController(func(fd uintptr) {
		controller.ProtectFd(int(fd))
	})
}
