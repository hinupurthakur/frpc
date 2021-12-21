package main

import (
	"math/rand"
	"time"

	"github.com/fatedier/golib/crypto"
	"github.com/hinupurthakur/frpc/frpclient"
)

func main() {
	crypto.DefaultSalt = "frp"
	rand.Seed(time.Now().UnixNano())
	frpclient.GetFullVersion()
	frpclient.Start()
}
