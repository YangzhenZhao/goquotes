package stock

import (
	"errors"
	"io"
	"log"
	"net/http"
	"strconv"
	"strings"
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
	ask_vol       [5]uint64
	bid_vol       [5]uint64
}

type SinaQuote struct {
}

func (quote *SinaQuote) Tick(code string) (*SinaTick, error) {
	res, err := http.Get(consts.SINA_BASE_URL + utils.GetExchangeCode(code))
	if err != nil {
		log.Fatal(err)
	}
	byteArr, err := io.ReadAll(res.Body)
	res.Body.Close()
	if err != nil {
		log.Fatal(err)
	}
	byteArr, err = utils.UTF8ToGBK(byteArr)
	if err != nil {
		log.Fatal(err)
	}
	return parse_out_tick(byteArr)
}

func (quote *SinaQuote) Price(code string) (float64, error) {
	res, err := quote.Tick(code)
	if err != nil {
		return 0, err
	}
	return res.current_price, err
}

func parse_out_tick(byteArr []byte) (*SinaTick, error) {
	msg := string(byteArr)
	fieldArr := strings.Split(msg, ",")
	if len(fieldArr) < 32 {
		return nil, errors.New("invalid code")
	}
	keyPos := strings.Index(fieldArr[0], `="`)
	code := fieldArr[0][keyPos-6 : keyPos]
	name := fieldArr[0][keyPos+2:]
	location, _ := time.LoadLocation("Asia/Shanghai")
	tick_time, err := time.ParseInLocation(consts.TIME_PARSE_DEFAULT, fieldArr[30]+" "+fieldArr[31], location)
	if err != nil {
		return nil, err
	}
	current_price, err := strconv.ParseFloat(fieldArr[3], 64)
	if err != nil {
		return nil, err
	}
	pre_close, err := strconv.ParseFloat(fieldArr[2], 64)
	if err != nil {
		return nil, err
	}
	open, err := strconv.ParseFloat(fieldArr[1], 64)
	if err != nil {
		return nil, err
	}
	high, err := strconv.ParseFloat(fieldArr[4], 64)
	if err != nil {
		return nil, err
	}
	low, err := strconv.ParseFloat(fieldArr[5], 64)
	if err != nil {
		return nil, err
	}
	total_amount, err := strconv.ParseFloat(fieldArr[9], 64)
	if err != nil {
		return nil, err
	}
	total_vol, err := strconv.ParseFloat(fieldArr[8], 64)
	if err != nil {
		return nil, err
	}
	ask_price := [5]float64{}
	ask_vol := [5]uint64{}
	bid_price := [5]float64{}
	bid_vol := [5]uint64{}
	for i := 0; i < 5; i++ {
		ask_price[i], err = strconv.ParseFloat(fieldArr[21+i*2], 64)
		if err != nil {
			return nil, err
		}
		ask_vol[i], err = strconv.ParseUint(fieldArr[20+i*2], 10, 64)
		if err != nil {
			return nil, err
		}
		bid_price[i], err = strconv.ParseFloat(fieldArr[11+i*2], 64)
		if err != nil {
			return nil, err
		}
		bid_vol[i], err = strconv.ParseUint(fieldArr[10+i*2], 10, 64)
		if err != nil {
			return nil, err
		}
	}
	return &SinaTick{
		tick_time,
		code,
		name,
		current_price,
		pre_close,
		open,
		high,
		low,
		total_amount,
		total_vol,
		ask_price,
		bid_price,
		ask_vol,
		bid_vol,
	}, nil
}
