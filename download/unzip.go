package download

import (
	"fmt"
	"os"
	"os/exec"
)

// ExtractTarGz 使用系统命令解压 .tar.gz 文件
func ExtractTarGz(tarGzFile, destDir string) error {
	// 创建目标目录
	if err := os.MkdirAll(destDir, os.ModePerm); err != nil {
		logw("创建目标目录失败: %s", err.Error())
		return err
	}

	// 使用系统命令解压 .tar.gz 文件
	cmd := exec.Command("tar", "-xzvf", tarGzFile, "-C", destDir)

	// 运行命令并捕获输出
	output, err := cmd.CombinedOutput()
	if err != nil {
		logw("解压 .tar.gz 文件失败: %s", err.Error())
	}

	// 打印命令输出
	fmt.Println(string(output))
	return nil
}
