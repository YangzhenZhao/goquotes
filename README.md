### goquotes





### Cli

```
$ go run main.go -h       
Usage:
   [flags]

Flags:
  -c, --current_price string   Set a code to get current price.
  -h, --help                   help for this command
  -t, --tick string            Set a code to get tick.
  -T, --ticks stringArray      Set some code to get tick.

$ go run main.go -t 000001                     
time: 2021-03-23 15:00:03 +0800 CST
code: 000001
name: 平安银行
current_price: 21.23
pre_close: 21.55
open: 21.57
high: 21.65
low: 20.96
total_amount: 1.44804361341e+09
total_vol: 6.8290338e+07
ask_price: [21.24 21.25 21.26 21.27 21.28]
bid_price: [21.23 21.22 21.21 21.2 21.19]
ask_vol: [500 27700 61900 100800 125000]
bid_vol: [91900 133293 145700 271300 28700]

$ go run main.go -T 000001 -T 000002 -T 600519
time: 2021-03-23 15:00:00 +0800 CST
code: 600519
name: 贵州茅台
current_price: 1996
pre_close: 1989.99
open: 1998.88
high: 2008.8
low: 1970
total_amount: 5.459520703e+09
total_vol: 2.744619e+06
ask_price: [1996.01 1996.02 1996.28 1996.36 1996.6]
bid_price: [1996 1995.99 1995.9 1995.89 1995.88]
ask_vol: [4800 400 800 200 200]
bid_vol: [7594 1600 100 200 100]

time: 2021-03-23 15:00:03 +0800 CST
code: 000001
name: 平安银行
current_price: 21.23
pre_close: 21.55
open: 21.57
high: 21.65
low: 20.96
total_amount: 1.44804361341e+09
total_vol: 6.8290338e+07
ask_price: [21.24 21.25 21.26 21.27 21.28]
bid_price: [21.23 21.22 21.21 21.2 21.19]
ask_vol: [500 27700 61900 100800 125000]
bid_vol: [91900 133293 145700 271300 28700]

time: 2021-03-23 15:00:03 +0800 CST
code: 000002
name: 万 科Ａ
current_price: 31.42
pre_close: 31.63
open: 31.61
high: 31.66
low: 31.14
total_amount: 1.32435457394e+09
total_vol: 4.2266007e+07
ask_price: [31.43 31.44 31.45 31.46 31.47]
bid_price: [31.42 31.41 31.4 31.39 31.38]
ask_vol: [60903 144900 34900 20140 12300]
bid_vol: [175160 9900 58400 33200 2800]

```