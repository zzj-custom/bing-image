package business

import (
	"image/internal/model/common"
)

type BingImages struct {
	Id            uint             `gorm:"column:id;primary_key;AUTO_INCREMENT"`
	Name          string           `gorm:"column:name;NOT NULL;comment:'图片名称'"`
	Copyright     string           `gorm:"column:copyright;NOT NULL;comment:'版权'"`
	CopyrightLink string           `gorm:"column:copyright_link;NOT NULL;comment:'版权链接'"`
	Url           string           `gorm:"column:url;NOT NULL;comment:'图片地址'"`
	Start         common.ShortTime `gorm:"column:start;NOT NULL;comment:'图片开始时间'"`
	End           common.ShortTime `gorm:"column:end;NOT NULL;comment:'图片结束时间'"`
	Location      string           `gorm:"column:location;NOT NULL;default:'zh-CN';comment:'位置，中国:zh-CN'"`
	ClickCount    int32            `gorm:"column:click_count;default:0;NOT NULL;comment:'点击次数'"`
	DownloadCount int32            `gorm:"column:download_count;default:0;NOT NULL;comment:'下载次数'"`
	Hash          string           `gorm:"column:hash;NOT NULL;comment:'哈希值'"`
	CreatedAt     common.LongTime  `gorm:"column:created_at;default:CURRENT_TIMESTAMP;comment:'创建时间'"`
	UpdatedAt     common.LongTime  `gorm:"column:updated_at;default:CURRENT_TIMESTAMP;comment:'更新时间'"`
}

func (b *BingImages) TableName() string {
	return "bing_images"
}
