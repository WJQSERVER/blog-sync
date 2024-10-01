package build

import (
	"fmt"
	"os/exec"

	"blog-sync/logger"
)

var logw = logger.Logw

func Build(targetDir string, baseUrl string) error {
	// 替换 config.yaml 中的 baseURL
	replaceUrl(baseUrl, targetDir)
	// 创建 hugo 命令，并指定工作目录
	cmd := exec.Command("hugo") // 这里可以替换为您需要的 hugo 命令
	cmd.Dir = targetDir         // 设置命令的工作目录

	// 运行命令并捕获输出
	output, err := cmd.CombinedOutput()
	if err != nil {
		logw("执行 hugo 命令失败: %w", err)
	}

	// 打印命令输出
	fmt.Println(string(output))
	logw(string(output))
	return nil
}

func replaceUrl(baseUrl string, targetDir string) {
	Hugo_cfgFile := targetDir + "/config.yaml"
	cmd := exec.Command("sed", "-i", fmt.Sprintf("s|baseURL: 'https://[^']*'|baseURL: '%s'|g", baseUrl), Hugo_cfgFile)

	// 执行命令
	err := cmd.Run()
	if err != nil {
		logw("Error executing sed: %v\n", err)
		return
	}

	logw("Successfully updated baseURL in config.yaml")
}
