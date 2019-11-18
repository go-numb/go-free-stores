# Stores.jp/Base.inへの商品の同時投稿

Stores.jp/Base.in両者税込み表示
引数は商品価格、プログラム内にて指定税率を乗算し、input

- Stores.jp: 税込み表示
- Base.in: 税込表示

 
``` golang
package workers

import (
	"testing"

	"github.com/go-numb/go-free-stores/sites"
)

func main() {
	client := New(nil, nil) // or New(*mongo.Session, *logrus.Logger)
	client.Start()
	defer client.Close()

	var workers []Worker
	workers = append(workers, []Worker{
		sites.NewStores("<id or email>", "<password>"),
		sites.NewBase("<id or email>", "<password>"),
	}...)

	for i, worker := range workers {
		if err := client.Product(worker, &ParamsForProduct{
			Title:       "テスト商品",
			Price:       2900,
			Description: "テスト商品説明文",
			Photos:      []string{"/Desktop/h2RoKcRj.png"},
			Discount:    10, // ％
			Stock:       10, // 在庫
			Tags:        []string{"タグ1", "タグ2"},
		}); err != nil {
			logrus.Fatal(i, err)
		}
    }
    
    logrus.Info("is done.")
}

```