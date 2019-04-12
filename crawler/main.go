package main

import (
	"golang-demo/crawler/engine"
	"golang-demo/crawler/zhenai/parser"
)

const url = "http://www.zhenai.com/zhenghun"

func main() {
	engine.Run(engine.Request{
		Url:        url,
		ParserFunc: parser.ParseCityList,
	})
}
