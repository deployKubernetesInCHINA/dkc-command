package cmd

import (
	"os"
	"testing"
)

func TestDownloadCache(t *testing.T){
	cmd := newDownloadCMD()
	cmd.SetArgs([]string{"cache","-m","http://packages-internal.bizconf.cn/packages/offline/"})
	cmd.SetOut(os.Stdout)
	cmd.Execute()
}

