package workers

import (
	"math"
	"time"

	"github.com/sclevine/agouti"
	"github.com/sirupsen/logrus"
	"gopkg.in/mgo.v2"
)

const (
	// TAX is 10%
	TAX float64 = 0.1
)

// Client is include webdrivers & logger
type Client struct {
	Chrome *agouti.WebDriver

	Mongo  *mgo.Session
	Logger *logrus.Logger
}

// ParamsForProduct is struct
type ParamsForProduct struct {
	Title       string   `json:"title" form:"title"`
	Description string   `json:"description" form:"description"`
	Photos      []string `json:"photos" form:"photos"`
	Price       int      `json:"price" form:"price"`
	Discount    int      `json:"discount" form:"discount"`
	Tags        []string `json:"tags" form:"tags"`
	Stock       int      `json:"stock" form:"stock"`
}

// New is webdriver with logger
func New(db *mgo.Session, log *logrus.Logger) *Client {
	d := agouti.ChromeDriver(
		agouti.ChromeOptions(
			"args", []string{
				"--headless", // headlessモードの指定
				"--disable-gpu",
				"no-sandbox",
				"--window-size=1280,800", // ウィンドウサイズの指定
			}),
		agouti.Debug,
	)
	d.Timeout = time.Duration(10 * time.Second)
	d.Start()

	return &Client{
		Chrome: d,

		Mongo:  db,
		Logger: log,
	}
}

// Start is close webdriver
func (p *Client) Start() error {
	if err := p.Chrome.Start(); err != nil {
		return err
	}
	return nil
}

// Close is close webdriver
func (p *Client) Close() error {
	if err := p.Chrome.Stop(); err != nil {
		return err
	}
	return nil
}

// AddTax return price * tax
func AddTax(price int) int {
	return int(math.RoundToEven(float64(price) * (TAX + 0.1)))
}
