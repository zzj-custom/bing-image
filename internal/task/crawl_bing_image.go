package task

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/golang-module/carbon"
	"github.com/pkg/errors"
	"go.uber.org/zap"
	"gorm.io/gorm/clause"
	"image/cmd/config"
	"image/internal/global"
	"image/internal/model/business"
	"image/internal/model/common"
	"image/internal/request"
	"net/http"
	url2 "net/url"
	"strings"
	"sync"
)

type ImagesResponse struct {
	Market   *Market   `json:"market"`
	Images   []*Images `json:"images"`
	Tooltips *Tooltips `json:"tooltips"`
}

type Market struct {
	Mkt string `json:"mkt"`
}

type Images struct {
	StartDate     string        `json:"startdate"`
	FullStartDate string        `json:"fullstartdate"`
	EndDate       string        `json:"enddate"`
	Url           string        `json:"url"`
	UrlBase       string        `json:"urlbase"`
	Copyright     string        `json:"copyright"`
	CopyrightLink string        `json:"copyrightlink"`
	Title         string        `json:"title"`
	Quiz          string        `json:"quiz"`
	Wp            bool          `json:"wp"`
	Hsh           string        `json:"hsh"`
	Drk           int           `json:"drk"`
	Top           int           `json:"top"`
	Bot           int           `json:"bot"`
	Hs            []interface{} `json:"hs"`
}

type Tooltips struct {
	Loading  string `json:"loading"`
	Previous string `json:"previous"`
	Next     string `json:"next"`
	Walle    string `json:"walle"`
	Walls    string `json:"walls"`
}

var (
	crawlBingImage     *CrawlBingImage
	crawlBingImageOnce sync.Once
)

type CrawlBingImage struct{}

func NewCrawlBingImage() *CrawlBingImage {
	crawlBingImageOnce.Do(func() {
		crawlBingImage = new(CrawlBingImage)
	})
	return crawlBingImage
}

func (receiver CrawlBingImage) Name() string {
	return "bing_image"
}

func (receiver CrawlBingImage) Spec() string {
	taskList := global.GVA_CONFIG.Cron.TaskList
	if spec, ok := taskList[receiver.Name()]; ok {
		return spec
	}

	defaultSpec := "0 0 5 * * *"
	global.GVA_LOG.Info(
		"当前任务尚未配置spec，使用默认数据",
		zap.String("name", receiver.Name()),
		zap.String("spec", defaultSpec),
	)

	return defaultSpec
}

func (receiver CrawlBingImage) Run() {
	cfg := global.GVA_CONFIG.Crawler.BingImage
	// 获取请求地址
	url := receiver.acquiredUrl(cfg)

	// 请求
	response, err := request.PerformRequest(
		context.Background(),
		http.MethodGet,
		url,
		nil,
		nil,
	)
	if err != nil {
		global.GVA_LOG.Error(
			"获取每日必应图片失败",
			zap.Error(err),
			zap.String("url", url),
			zap.ByteString("response", response),
		)
		return
	}

	// 解析数据
	var imagesResponse ImagesResponse
	if err = json.Unmarshal(response, &imagesResponse); err != nil {
		global.GVA_LOG.Error(
			"每日必应图片解析数据失败",
			zap.Error(err),
			zap.ByteString("response", response),
			zap.String("url", url),
		)
		return
	}

	bingImagesInsertData := make([]*business.BingImages, 0, len(imagesResponse.Images))
	for _, item := range imagesResponse.Images {
		// 判断超高清图片是否存在
		uhdImageUrl, imageUrl := receiver.queryUHDUrl(cfg, item)
		ok, err := receiver.isExistsByImage(uhdImageUrl)
		if err != nil {
			global.GVA_LOG.Error(
				"判断必应图片是否存在失败",
				zap.Error(err),
				zap.String("udhUrl", uhdImageUrl),
			)
		}

		if ok {
			imageUrl = uhdImageUrl
		}

		ida := &business.BingImages{
			Name:          item.Title,
			Copyright:     item.Copyright,
			CopyrightLink: fmt.Sprintf("%s%s", cfg.URI, item.CopyrightLink),
			Url:           imageUrl,
			Start:         common.ShortTime{Time: carbon.Parse(item.StartDate).ToStdTime()},
			End:           common.ShortTime{Time: carbon.Parse(item.EndDate).ToStdTime()},
			Hash:          item.Hsh,
		}

		if imagesResponse.Market != nil {
			ida.Location = imagesResponse.Market.Mkt
		} else {
			ida.Location = cfg.MKT
		}

		bingImagesInsertData = append(bingImagesInsertData, ida)
	}

	// 创建数据
	if err = receiver.createData(bingImagesInsertData); err != nil {
		global.GVA_LOG.Error(
			"创建必应图片数据失败",
			zap.Error(err),
		)
		return
	}

	// TODO: 上传到其他平台（本地，或者是阿里云，七牛云）
}

// 获取请求地址
func (receiver CrawlBingImage) acquiredUrl(cfg *config.BingImage) string {
	// 获取请求的地址
	url := cfg.URI

	// 参数配置
	params := fmt.Sprintf("n=%d&idx=%d", cfg.N, cfg.Idx)

	if cfg.Format != "" {
		params += fmt.Sprintf("&format=%s", cfg.Format)
	}

	if cfg.Size != "" {
		params += fmt.Sprintf("&size=%s", cfg.Size)
	}

	if cfg.MKT != "" {
		params += fmt.Sprintf("&mkt=%s", cfg.MKT)
	}

	if cfg.Mbl != 0 {
		params += fmt.Sprintf("&mbl=%d", cfg.Mbl)
	}

	return fmt.Sprintf("%s?%s", url, params)
}

// 判断超高清图片是否存在
func (receiver CrawlBingImage) isExistsByImage(url string) (bool, error) {
	resp, err := http.Head(url)
	if err != nil {
		return false, err
	}
	defer func() {
		if err = resp.Body.Close(); err != nil {
			global.GVA_LOG.Error(
				"关闭判断请求图片是否存在的body失败",
				zap.Error(err),
				zap.String("url", url),
			)
		}
	}()

	if resp.StatusCode == http.StatusOK {
		// 远程图片存在
		return true, nil
	} else if resp.StatusCode == http.StatusNotFound {
		// 远程图片不存在
		return false, nil
	} else {
		// 其他错误情况
		return false, errors.Errorf("查询图片数据失败:err=%d", resp.StatusCode)
	}
}

// 获取高清地址
func (receiver CrawlBingImage) queryUHDUrl(cfg *config.BingImage, images *Images) (string, string) {
	size := cfg.Size
	if size == "" {
		size = "1920x1080"
	}

	u, _ := url2.Parse(cfg.URI)

	// 获取UHD地址
	url := fmt.Sprintf("https://%s%s", u.Host, images.Url)

	return strings.ReplaceAll(url, size, "UHD"), url
}

// 创建数据
func (receiver CrawlBingImage) createData(data []*business.BingImages) error {
	if len(data) == 0 {
		global.GVA_LOG.Info("创建必应图片获取到的数据为空")
	}

	return global.GVA_DB.Clauses(clause.OnConflict{
		Columns: []clause.Column{
			{Name: "hash"},
		},
		DoUpdates: clause.AssignmentColumns([]string{
			"name",
			"copyright",
			"copyright_link",
			"url", "start",
			"end", "location",
		}),
	}).Create(&data).Error
}
