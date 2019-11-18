package sites

import (
	"strings"
)

// Base is struct
type Base struct {
	UUID         string
	ID, Password string
}

// NewBase is New
func NewBase(id, pass string) *Base {
	return &Base{
		UUID:     "base",
		ID:       id,
		Password: pass,
	}
}

// Info return access url
func (p *Base) Info() (uuid, id, pass string) {
	return p.UUID, p.ID, p.Password
}

// URL return access url
func (p *Base) URL() string {
	return "https://thebase.in"
}

// Login return access url
func (p *Base) Login() string {
	return "https://admin.thebase.in/users/login"
}

// Product return Product url
func (p *Base) Product() string { // 管理ページ
	return "https://admin.thebase.in/shop_admin/items/add"
}

//ChooseElement is 各処理に必要なエレメント(id, class etc...)
func (p *Base) ChooseElement(isButton bool, key string) string {
	if !isButton { // block:input
		switch key {
		case "login":
			return `input[id="loginUserMailAddress"]`

		case "password":
			return `input[id="UserPassword"]`

		}
	}

	switch key {
	case "login":
		return `form[id="userLoginForm"]`

	case "access":
		return `input[value="最新情報を登録する"]`

	case "file":
		return ".m-uploadBox__input"

	}

	return ""
}

// SetPhotos uploads photos
// 画像は20枚まで追加できます
// 推奨サイズは1280px × 1280pxです
// 10MB以内の画像ファイルを用意してください
// 対応ファイル：jpg,png,gif
func (p *Base) SetPhotos(files []string) string {
	if 20 < len(files) {
		files = files[:20]
	}

	return strings.Join(files, ",")
}
