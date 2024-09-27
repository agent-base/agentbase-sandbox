package main

import "github.com/agent-base/agentbase-sandbox/internal/core/lib/nodejs"
import "C"

//export AgentBaseSeccomp
func AgentBaseSeccomp(uid int, gid int, enable_network bool) {
	nodejs.InitSeccomp(uid, gid, enable_network)
}

func main() {}
