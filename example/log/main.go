package log

import "github.com/num5/loger"

var log *loger.Log

func init() {
	// 初始化
	log = loger.NewLog(1000)

	// 设置log级别
	log.SetLevel("Debug")

	// 设置输出引擎
	log.SetEngine("file", `{"level":4, "spilt":"size", "filename":".logs/log.log", "maxsize":10}`)

	//log.DelEngine("console")

	// 设置是否输出行号
	//log.SetFuncCall(true)
}
