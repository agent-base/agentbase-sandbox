package main

import (
	"github.com/agent-base/agentbase-sandbox/internal/core/lib/python"
)
import "C"

//export AgentBaseSeccomp
func AgentBaseSeccomp(uid int, gid int, enable_network bool) {
	python.InitSeccomp(uid, gid, enable_network)
}

func main() {}
