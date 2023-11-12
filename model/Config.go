package model

import (
	"fmt"
	"io"
	"os"
	"os/exec"
	"runtime"
	"time"

	"golang.org/x/term"
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
			"-o", "TCPKeepAlive=yes",
		)
	} else {
		// 如果不是Windows系统，直接使用sshpass和ssh命令
		cmd = exec.Command("sshpass",
			"-p", s.Password,
			"ssh",
			"-t", // 强制分配伪终端
			fmt.Sprintf("%s@%s", s.Username, s.IP),
			"-o", "StrictHostKeyChecking=no",
			"-o", "UserKnownHostsFile=/dev/null",
			"-o", "TCPKeepAlive=yes",
		)
	}

	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	var stdinPipe io.WriteCloser
	var err error

	if term.IsTerminal(int(os.Stdin.Fd())) {
		// 如果stdin是终端，强制分配伪终端
		cmd.Stdin = os.Stdin
	} else {
		// 如果stdin不是终端，创建管道并使用它
		stdinPipe, err = cmd.StdinPipe()
		if err != nil {
			return err
		}
		defer stdinPipe.Close()

		go func() {
			// 从os.Stdin复制到stdinPipe
			io.Copy(stdinPipe, os.Stdin)
		}()
	}

	err = cmd.Start()
	if err != nil {
		return err
	}

	// 启动保活goroutine
	go func() {
		ticker := time.NewTicker(1 * time.Minute) // 每分钟发送一次保活请求
		defer ticker.Stop()

		for range ticker.C {
			if cmd.ProcessState != nil && cmd.ProcessState.Exited() {
				break
			}

			if stdinPipe != nil {
				// 如果stdinPipe存在，使用它发送保活信号
				_, err := stdinPipe.Write([]byte("\n"))
				if err != nil {
					fmt.Println("Error writing to stdin:", err)
					break
				}
			}
		}
	}()

	err = cmd.Wait()
	if err != nil {
		return err
	}

	return nil
}

func (s *Server) UploadFile(localPath, remotePath string) error {
	cmd := exec.Command("sshpass", "-p", s.Password, "ssh", "-o", "StrictHostKeyChecking=no", fmt.Sprintf("%s@%s", s.Username, s.IP), "cat > "+remotePath)

	stdin, err := cmd.StdinPipe()
	if err != nil {
		return fmt.Errorf("error getting stdin pipe: %v", err)
	}
	defer stdin.Close()

	err = cmd.Start()
	if err != nil {
		return fmt.Errorf("error starting command: %v", err)
	}

	file, err := os.Open(localPath)
	if err != nil {
		return fmt.Errorf("error opening local file: %v", err)
	}
	defer file.Close()

	_, err = io.Copy(stdin, file)
	if err != nil {
		return fmt.Errorf("error copying file content: %v", err)
	}

	err = cmd.Wait()
	if err != nil {
		return fmt.Errorf("error waiting for command: %v", err)
	}

	fmt.Println("File uploaded successfully!")
	return nil
}
