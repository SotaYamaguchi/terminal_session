package main

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"time"
)

func main() {
	// 現在時刻を "YYYYMMDDhhmm" フォーマットで取得
	timestamp := time.Now().Format("200601021504")

	// ログファイルのディレクトリとパスを設定
	logDir := filepath.Join(os.Getenv("HOME"), ".terminal_session/logs")
	logFilePath := filepath.Join(logDir, fmt.Sprintf("terminal_session_%s.log", timestamp))

	// ログファイルのディレクトリが存在しない場合は作成
	if err := os.MkdirAll(logDir, os.ModePerm); err != nil {
		fmt.Printf("Error creating log directory: %v\n", err)
		os.Exit(1)
	}

	// セッション開始時刻をログファイルに書き込む
	startTime := time.Now().Format("2006-01-02 15:04:05")
	f, err := os.OpenFile(logFilePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Printf("Error opening log file: %v\n", err)
		os.Exit(1)
	}
	defer f.Close()

	if _, err := f.WriteString(fmt.Sprintf("Session started at: %s\n", startTime)); err != nil {
		fmt.Printf("Error writing start time to log file: %v\n", err)
		os.Exit(1)
	}

	// `script` コマンドを実行して、セッションをログファイルに記録
	cmd := exec.Command("script", "-a", logFilePath) // -a で追記モードにする
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	// `script` コマンドを終了するまで維持
	if err := cmd.Run(); err != nil {
		fmt.Printf("Error running script command: %v\n", err)
		os.Exit(1)
	}

	// セッション終了時刻をログファイルに書き込む
	endTime := time.Now().Format("2006-01-02 15:04:05")
	if _, err := f.WriteString(fmt.Sprintf("Session ended at: %s\n", endTime)); err != nil {
		fmt.Printf("Error writing end time to log file: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("Session logged to %s\n", logFilePath)
}
