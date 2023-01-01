package main

import (
	"log"
	"os"
	"os/exec"
	"syscall"
)

func main() {
	cmd := exec.Command("sh") // 用来指定被fork出来的新进程内的初始命令，默认使用sh来执行
	cmd.SysProcAttr = &syscall.SysProcAttr{
		// 同时隔离 uts、ipc、pid、mount、user
		Cloneflags: syscall.CLONE_NEWUTS | syscall.CLONE_NEWIPC |
			syscall.CLONE_NEWPID | syscall.CLONE_NEWNS |
			syscall.CLONE_NEWUSER | syscall.CLONE_NEWNET,
	}
	// 发现去掉cmd.SysProcAttr.Credential最终效果一致，猜想是高版本内核集成了
	// （内核版本高于书本编写时版本），会导致： fork/exec /usr/bin/sh: operation not permitted
	// cmd.SysProcAttr.Credential = &syscall.Credential{Uid: uint32(1), Gid: uint32(1)}
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		log.Fatal(err)
	}
}
