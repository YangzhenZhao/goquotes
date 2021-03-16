package stock

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"time"

	"github.com/YangzhenZhao/goquotes/quotes/consts"
	"github.com/YangzhenZhao/goquotes/quotes/utils"
)

type SinaTick struct {
	time          time.Time
	code          string
	name          string
	current_price float64
	pre_close     float64
	open          float64
	high          float64
	low           float64
	total_amount  float64
	total_vol     float64
	ask_price     [5]float64
	bid_price     [5]float64
	ask_vol       [5]int32
	bid_vol       [5]int32
}

type SinaQuote struct {
}

func (quote *SinaQuote) Tick(code string) {
	println(code)
	println(consts.SINA_BASE_URL)
	println("hhhh")
	res, err := http.Get(consts.SINA_BASE_URL + utils.GetExchangeCode(code))
	if err != nil {
		log.Fatal(err)
	}
	text, err := io.ReadAll(res.Body)
	res.Body.Close()
	if err != nil {
		log.Fatal(err)
	}
	text, err = utils.UTF8ToGBK(text)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s", text)
}
