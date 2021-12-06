package polarLog

import (
	"path/filepath"
	"time"
)

type PolarLog struct {
	Path string
	Format string
	FileName string
	ExtraData ExtraData
}

type ExtraData struct {
	Date string
}
//初始化
func (polarLog *PolarLog) Default() *PolarLog  {
	polarLog.Path = polarLog.GetPath(polarLog.Path)
	polarLog.FileName = polarLog.GetFileNameByCurrentTime(polarLog.Format)
	polarLog.GetExtraData()
	return polarLog
}
//获取文件目录
func (polarLog *PolarLog) GetPath(path string) (abs string) {
	abs, _ = filepath.Abs(path)
	return
}
//根据时间生成文件名称
func (polarLog *PolarLog) GetFileNameByCurrentTime(format string) (filename string) {
	if len(format) == 0 {
		filename = time.Now().Format("2006010215")
	}else {
		filename = time.Now().Format(format)
	}
	return
}

//获取额外数据
func (polarLog *PolarLog) GetExtraData()  {
	polarLog.ExtraData.Date = time.Now().Format("2006-01-02 15:04:05")
}
