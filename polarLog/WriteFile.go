package polarLog

import (
	"io"
	"os"
)

func  (polarLog *PolarLog) WriteFile(completePath ,content string)   {
	var f *os.File
	if polarLog.CheckFileIsExist(completePath) { //如果文件存在
		f,_ = os.OpenFile(completePath, os.O_APPEND, 0666) //打开文件
	} else {
		if polarLog.CheckFileIsExist(polarLog.Path) {
			f, _ = os.Create(completePath) //创建文件
		}else {
			_ = os.MkdirAll(polarLog.Path, 0666)  //创建目录
			f, _ = os.Create(completePath) //创建文件
		}
	}
	_, _ = io.WriteString(f, content) //写入文件(字符串)
}

//判断文件是否存在
func (polarLog *PolarLog) CheckFileIsExist(filename string) bool  {
	stat, err := os.Stat(filename)
	if stat != nil && err == nil {
		return true
	}else if  os.IsNotExist(err) {
		return false
	}
	return false
}

