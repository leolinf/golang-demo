package fetcher

func get_header() map[string]string {
	// var Header map[string]string
	Header := make(map[string]string)
	Header["Accept"] = "text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3"
	Header["Accept-Encoding"] = "gzip, deflate"
	Header["Accept-Language"] = "zh-CN,zh;q=0.9,en-US;q=0.8,en;q=0.7"
	Header["Cache-Control"] = "max-age=0"
	Header["Connection"] = "keep-alive"
	Header["Cookie"] = "ipCityCode=10127000; sid=9f7d18e8-43dd-42f1-a04e-141a2fedbf4a; Hm_lvt_2c8ad67df9e787ad29dbd54ee608f5d2=1563432955; Hm_lpvt_2c8ad67df9e787ad29dbd54ee608f5d2=1563434354"
	Header["Host"] = "album.zhenai.com"
	Header["Upgrade-Insecure-Requests"] = "1"
	Header["User-Agent"] = "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_12_6) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/75.0.3770.100 Safari/537.36['']"
	return Header
}
