package model

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"
)

type Server struct {
	Alias    string `json:"alias"`
	IP       string `json:"ip"`
	Username string `json:"username"`
	Password string `json:"password"`
}

type Config struct {
	ServerList []Server `json:"serverList"`
}

func (s *Server) Connect() error {
	var cmd *exec.Cmd

	if runtime.GOOS == "windows" {
		// 如果是Windows系统，使用plink代替ssh
		cmd = exec.Command("plink",
			fmt.Sprintf("%s@%s", s.Username, s.IP),
			"-P", "22", // 指定SSH端口
			"-pw", s.Password,
			"-batch", // 禁用交互模式，确保不会出现任何提示
		)
	} else {
		// 如果不是Windows系统，直接使用sshpass和ssh命令
		cmd = exec.Command("sshpass",
			"-p", s.Password,
			"ssh",
			fmt.Sprintf("%s@%s", s.Username, s.IP),
			"-o", "StrictHostKeyChecking=no",
			"-o", "UserKnownHostsFile=/dev/null",
		)
	}

	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Stdin = os.Stdin

	err := cmd.Run()
	if err != nil {
		return err
	}

	return nil
}
