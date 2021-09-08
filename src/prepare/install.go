package prepare

import (
	"os"
	"os/exec"
	"path/filepath"

	"github.com/TwinProduction/go-color"

	"github.com/deployKubernetesInCHINA/dkc-command/src"
	"github.com/deployKubernetesInCHINA/dkc-command/src/config"
	"github.com/deployKubernetesInCHINA/dkc-command/src/pkg/log"
)

func InstallDocker() {
	log.Log.Println("install docker.")
	var dockerPath string
	absPath, _ := filepath.Abs(".")
	if src.Release == "7" {
		dockerPath = filepath.Join(absPath, config.El7DirName, "docker/")
	} else {
		dockerPath = filepath.Join(absPath, config.El8DirName, "docker/")
	}
	cmd := exec.Command("sh", "-c", "yum install --disablerepo='*' -y "+dockerPath+"/*.rpm")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stdout
	cmd.Run()
}

func InstallAnsible() {
	log.Log.Println("install ansible.")
	var pyPath, ansiblePath, localPath string
	absPath, _ := filepath.Abs(".")
	if src.Release == "7" {
		pyPath = filepath.Join(absPath, config.El7DirName, "python3/")
		localPath = filepath.Join(absPath, config.El7DirName, "local/")
	} else {
		pyPath = filepath.Join(absPath, config.El8DirName, "python3/")
		localPath = filepath.Join(absPath, config.El8DirName, "local/")
	}
	ansiblePath = filepath.Join(absPath, "ansible")
	//install python3

	cmd := exec.Command("sh", "-c", "yum install --disablerepo=* -y "+pyPath+"/*.rpm")
	log.Log.Println(cmd.String())
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stdout
	if err := cmd.Run(); err != nil {
		log.Log.Println(color.Ize(color.Red, err.Error()))
	}
	//upgrade pip
	cmd = exec.Command("python3", "-m", "pip", "install", "--no-index", pyPath+"/pip-21.0.1-py3-none-any.whl")
	log.Log.Println(cmd.String())
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stdout
	if err := cmd.Run(); err != nil {
		log.Log.Println(color.Ize(color.Red, err.Error()))
	} // install ansible

	cmd = exec.Command("python3", "-m", "pip", "install", "--no-index", "--find-link=ansible", "-r", ansiblePath+"/requirements.txt")
	log.Log.Println(cmd.String())
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stdout
	if err := cmd.Run(); err != nil {
		log.Log.Println(color.Ize(color.Red, err.Error()))
	}
	// install sshpass
	if _, err := exec.LookPath("sshpass"); err != nil {
		cmd := exec.Command("sh", "-c", "yum install --disablerepo=* -y "+localPath+"/*.rpm")
		log.Log.Println(cmd.String())
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stdout
		if err := cmd.Run(); err != nil {
			log.Log.Println(color.Ize(color.Red, err.Error()))
		}
	}
}
