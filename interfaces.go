package workers

import (
	"fmt"
	"strings"
	"time"

	"github.com/sclevine/agouti"
)

// Worker is website workers
type Worker interface {
	URL() string
	Login() string
	Product() string

	Info() (uuid, id, pass string)
	ChooseElement(isButton bool, key string) string

	SetPhotos(files []string) string
}

// Product is 商品登録
func (p *Client) Product(w Worker, params *ParamsForProduct) error {
	page, err := p.Chrome.NewPage()
	if err != nil {
		return err
	}
	if err := p.setting(page); err != nil {
		return err
	}
	defer page.CloseWindow()

	uuid, id, pass := w.Info()
	if err := login(page, w, uuid, id, pass); err != nil {
		return err
	}

	if err := create(uuid, w, params, page); err != nil {
		return err
	}

	return nil
}

// setting is s page in webdriver
func (p *Client) setting(page *agouti.Page) error {
	if err := page.SetImplicitWait(1000); err != nil {
		return err
	}
	if err := page.SetPageLoad(15000); err != nil {
		return err
	}
	return nil
}

func login(page *agouti.Page, w Worker, uuid, id, pass string) error {
	if err := page.Navigate(w.Login()); err != nil {
		return err
	}

	switch uuid {
	case "stores":
		// pop up click
		if err := page.FindByLink("ログイン").Click(); err != nil {
			return err
		}

	case "base":

	}

	if err := page.Find(w.ChooseElement(false, "login")).Fill(id); err != nil {
		return err
	}
	if err := page.Find(w.ChooseElement(false, "password")).Fill(pass); err != nil {
		return err
	}

	if err := page.Find(w.ChooseElement(true, "login")).Click(); err != nil {
		return err
	}

	// Element returnに[form]が含まれる場合は、Submit
	// ない場合は、input buttonのClick
	if !strings.Contains(w.ChooseElement(true, "login"), "form") {
		if err := page.Find(w.ChooseElement(true, "login")).Click(); err != nil { // ボタンクリックがだめならば、form Submit
			return err
		}
	} else {
		if err := page.Find(w.ChooseElement(true, "login")).Submit(); err != nil {
			return err
		}
	}

	return nil
}

func create(uuid string, w Worker, params *ParamsForProduct, page *agouti.Page) error {
	switch uuid {
	case "stores":
		time.Sleep(3 * time.Second)
		// pop up click
		isThere, err := page.FindByID("ngdialog1-aria-labelledby").Visible()
		if err != nil {
			return err
		}
		if isThere {
			if err := page.EnterPopupText("あとにする"); err != nil {
				return err
			}
		}

		if err := page.Navigate(w.Product()); err != nil {
			return err
		}

		if err := page.Find(w.ChooseElement(true, "file")).Click(); err != nil {
			return err
		}
		if err := page.Find("label > input").SendKeys(strings.Join(params.Photos, ",")); err != nil {
			return err
		}
		time.Sleep(2 * time.Second)

		if err := page.FindByName("name").Fill(params.Title); err != nil {
			return err
		}
		if err := page.FindByName("price").Fill(fmt.Sprintf("%d", AddTax(params.Price))); err != nil {
			return err
		}
		if err := page.FindByName("detailPercent").Fill(fmt.Sprintf("%d", params.Discount)); err != nil {
			return err
		}
		if err := page.Find(".switch_text_off").Click(); err != nil {
			return err
		}
		if err := page.FindByName("description").Fill(params.Description); err != nil {
			return err
		}
		if err := page.Find(".hashtag_input_container > .form_input").Fill(w.SetPhotos(params.Photos)); err != nil {
			return err
		}
		if err := page.FindByName("quantity_0").Fill(fmt.Sprintf("%d", params.Stock)); err != nil {
			return err
		}

		if err := page.Find("body > div:nth-child(5) > div > div > main > form > div:nth-child(2)").Click(); err != nil {
			return err
		}

	case "base":
		if err := page.Navigate(w.Product()); err != nil {
			return err
		}

		// if err := page.FindByLabel("商品を登録する").Click(); err != nil {
		// 	return err
		// }

		if err := page.Find(w.ChooseElement(true, "file")).Click(); err != nil {
			return err
		}
		if err := page.Find(w.ChooseElement(true, "file")).SendKeys(w.SetPhotos(params.Photos)); err != nil {
			return err
		}
		time.Sleep(2 * time.Second)

		if err := page.FindByID("itemDetail_name").Fill(params.Title); err != nil {
			return err
		}
		if err := page.FindByID("itemDetail_price").Fill(fmt.Sprintf("%d", AddTax(params.Price))); err != nil {
			return err
		}
		if err := page.FindByID("itemDetail_detail").Fill(params.Description); err != nil {
			return err
		}
		if err := page.FindByID("itemDetail_stock").Fill(fmt.Sprintf("%d", params.Stock)); err != nil {
			return err
		}

		if err := page.Find(".btn_20KSNrV9 > .c-submitBtn__icon > div").Click(); err != nil {
			return err
		}

	}

	return nil
}
