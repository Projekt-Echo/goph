package netease

import (
	"github.com/go-resty/resty/v2"
	"time"
)

var RestyClient *resty.Client

func init() {
	RestyClient = resty.New().SetTimeout(10 * time.Second)
}
