package sites

import (
	"strings"
)

// Stores is struct
type Stores struct {
	UUID         string
	ID, Password string
}

// NewStores is New
func NewStores(id, pass string) *Stores {
	return &Stores{
		UUID:     "stores",
		ID:       id,
		Password: pass,
	}
}

// Info return access url
func (p *Stores) Info() (uuid, id, pass string) {
	return p.UUID, p.ID, p.Password
}

// URL return access url
func (p *Stores) URL() string {
	return "https://stores.jp"
}

// Login return access url
func (p *Stores) Login() string {
	return "https://stores.jp"
}

// Product return Product url
func (p *Stores) Product() string { // 管理ページ
	return "https://stores.jp/dashboard#!/items/new"
}

//ChooseElement is 各処理に必要なエレメント(id, class etc...)
func (p *Stores) ChooseElement(isButton bool, key string) string {
	if !isButton { // block:input
		switch key {
		case "login":
			return `.st-dialog__default:nth-child(1) .input-list-item:nth-child(1) .text-inner:nth-child(1)`

		case "password":
			return `.st-dialog__default:nth-child(1) .input-list-item:nth-child(2) .text-inner:nth-child(1)`

		}
	}

	switch key {
	case "login":
		return `.button-list-item > .topButton--fill`

	case "access":
		return `input[value="最新情報を登録する"]`

	case "file":
		return ".sj-file-upload"

	}

	return ""
}

// SetPhotos uploads photos
// 画像は20枚まで追加できます
// 推奨サイズは1280px × 1280pxです
// 10MB以内の画像ファイルを用意してください
// 対応ファイル：jpg,png,gif
func (p *Stores) SetPhotos(files []string) string {
	if 15 < len(files) {
		files = files[:15]
	}

	return strings.Join(files, "\n")
}
