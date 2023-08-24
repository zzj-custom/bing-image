package request

import (
	"bytes"
	"context"
	"github.com/pkg/errors"
	"go.uber.org/zap"
	"image/internal/global"
	"io/ioutil"
	"net/http"
	"time"
)

func PerformRequest(ctx context.Context, method string, url string, headers map[string]string, body []byte) ([]byte, error) {
	// 添加日志参数
	global.GVA_LOG.With(
		zap.String("method", method),
		zap.String("url", url),
		zap.Reflect("header", headers),
		zap.ByteString("body", body),
	)

	client := http.Client{}

	req, err := http.NewRequest(method, url, nil)
	if err != nil {
		return nil, errors.Wrap(err, "NewRequest 构建新请求失败")
	}

	// 设置请求头
	for key, value := range headers {
		req.Header.Set(key, value)
	}

	// 设置请求体
	if body != nil {
		req.Body = ioutil.NopCloser(bytes.NewBuffer(body))
	}

	// 使用WithCancel来创建一个有超时的上下文
	ctx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel() // 在函数退出时取消上下文以释放资源

	req = req.WithContext(ctx)

	resp, err := client.Do(req)
	if err != nil {
		return nil, errors.Wrap(err, "Do 请求失败")
	}
	defer func() {
		if err = resp.Body.Close(); err != nil {
			global.GVA_LOG.Error("关闭请求body失败", zap.Error(err))
		}
	}()

	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		global.GVA_LOG.Error("Do 请求失败", zap.ByteString("respBody", respBody))
		return nil, errors.Wrap(err, "ReadAll 读取数据失败")
	}

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		global.GVA_LOG.Error(
			"请求code错误",
			zap.Int("statusCode", resp.StatusCode),
			zap.ByteString("respBody", respBody),
		)
		return nil, errors.Errorf("statusCode【%d】错误", resp.StatusCode)
	}

	return respBody, nil
}
