# go-ipip
[![Build Status](https://travis-ci.org/yinqiang/go-ipip.svg?branch=master)](https://travis-ci.org/yinqiang/go-ipip)
[IP数据库](http://ipip.net) Golang语言解析库

## Usage
```go
p := NewIpip()
if err := p.Load("17monipdb.dat"); err != nil {
    panic(err)
}
info, err := p.Find("8.8.8.8")
if err != nil {
    panic(err)
}
fmt.Printf("Country:%s, Region:%s, City:%s, Isp:%s\n",
    info.Country, info.Region, info.City, info.Isp)
```
