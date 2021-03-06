package stock

import (
	"errors"
	"fmt"
	"io"
	"net/http"
	"reflect"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/YangzhenZhao/goquotes/quotes/consts"
	"github.com/YangzhenZhao/goquotes/quotes/utils"
)

type SinaTick struct {
	Time         time.Time
	Code         string
	Name         string
	CurrentPrice float64
	PreClose     float64
	Open         float64
	High         float64
	Low          float64
	TotalAmount  float64
	TotalVol     float64
	AskPrice     [5]float64
	BidPrice     [5]float64
	AskVol       [5]uint64
	BidVol       [5]uint64
}

func (tick *SinaTick) Print() {
	fmt.Println("Time:", &tick.Time)
	t := reflect.TypeOf(tick)
	v := reflect.ValueOf(tick)
	if t.Kind() == reflect.Ptr {
		t = t.Elem()
		v = v.Elem()
	}
	for i := 1; i < t.NumField(); i++ {
		f := t.Field(i)
		fmt.Printf("%s: ", f.Name)
		fmt.Printf("%v\n", v.Field(i))
	}
}

type SinaQuote struct {
}

func (quote *SinaQuote) TickMap(codes []string) map[string]*SinaTick {
	var waitGroup sync.WaitGroup

	ticks := []*SinaTick{}

	codesNum := len(codes)
	i := 0
	for {
		waitGroup.Add(1)
		if (i+1)*consts.REQ_CODES_NUM_MAX >= codesNum {
			go func() {
				defer waitGroup.Done()
				ticksTmp := smallTicks(codes[i*consts.REQ_CODES_NUM_MAX:])
				ticks = append(ticks, ticksTmp...)
			}()
			break
		}
		go func() {
			defer waitGroup.Done()
			ticksTmp := smallTicks(codes[i*consts.REQ_CODES_NUM_MAX : (i+1)*consts.REQ_CODES_NUM_MAX])
			ticks = append(ticks, ticksTmp...)
		}()
		i += 1
	}

	waitGroup.Wait()

	res := make(map[string]*SinaTick)
	for _, tick := range ticks {
		res[tick.Code] = tick
	}
	return res
}

func (quote *SinaQuote) Tick(code string) (*SinaTick, error) {
	res := smallTicks([]string{code})
	if len(res) == 0 {
		return nil, errors.New("invalid code")
	}
	return res[0], nil
}

func (quote *SinaQuote) Price(code string) (float64, error) {
	res, err := quote.Tick(code)
	if err != nil {
		return 0, err
	}
	return res.CurrentPrice, err
}

func parse_out_tick(msg string) (*SinaTick, error) {
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

func smallTicks(codes []string) []*SinaTick {
	res := []*SinaTick{}
	contents := fetchTickContents(codes)
	for _, content := range strings.Split(contents, "\n") {
		tick, err := parse_out_tick(content)
		if err == nil {
			res = append(res, tick)
		}
	}
	return res
}

func fetchTickContents(codes []string) string {
	utils.ToExchangeCodes(codes)
	res, err := http.Get(consts.SINA_BASE_URL + strings.Join(codes, ","))
	if err != nil {
		return ""
	}
	byteArr, err := io.ReadAll(res.Body)
	res.Body.Close()
	if err != nil {
		return ""
	}
	byteArr, err = utils.UTF8ToGBK(byteArr)
	if err != nil {
		return ""
	}
	return string(byteArr)
}
