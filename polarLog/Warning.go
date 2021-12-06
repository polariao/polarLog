package polarLog

import (
	"github.com/polariao/polarLog/conf"
	"runtime"
)

func (polarLog *PolarLog) Warning(content string)  {
	path := polarLog.Path + "/"
	fileName := conf.WARNING + polarLog.FileName + conf.SUFFIX
	completePath := path + fileName
	goos := runtime.GOOS

	inputFile := ""
	inputFile += polarLog.ExtraData.Date+"||"
	if goos == "windows" {
		inputFile += content+"\t\n"
	}else {
		inputFile += content+"\n"
	}
	polarLog.WriteFile(completePath,inputFile)
}
