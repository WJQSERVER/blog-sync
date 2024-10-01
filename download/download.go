package download

import (
	"fmt"
	"os"

	"blog-sync/logger"

	"github.com/imroc/req/v3"
)

var logw = logger.Logw

func DownloadFile(url, username, password, savePath, destDir string) error {
	// 创建一个新的req客户端
	client := req.C()

	client.SetUserAgent("Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/129.0.0.0 Safari/537.36").
		SetTLSFingerprintChrome().
		ImpersonateChrome()

	// 创建文件
	file, err := os.Create(savePath)
	if err != nil {
		return fmt.Errorf("could not create file: %v", err)
	}

	defer file.Close() // 确保在函数结束时关闭文件
	// 请求文件并保存
	resp, err := client.R().
		SetBasicAuth(username, password).
		SetOutput(file).
		Get(url)

	if err != nil {
		return fmt.Errorf("request error: %v", err)
	}

	if err := ExtractTarGz(savePath, destDir); err != nil {
		logw("解压失败:", err)
	}

	if resp.StatusCode != 200 {
		logw("request failed with status code: %d", resp.StatusCode)
	}

	logw("下载完成")

	return nil
}
