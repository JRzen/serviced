package isvcs

import (
	"github.com/zenoss/glog"
	"github.com/zenoss/serviced"

	"path"
	"runtime"
)

var Mgr *Manager

const (
	IMAGE_REPO = "zctrl/isvcs"
	IMAGE_TAG  = "v1"
)

func Init() {
	Mgr = NewManager("unix:///var/run/docker.sock", imagesDir())

	if err := Mgr.Register(elasticsearch); err != nil {
		glog.Fatalf("%s", err)
	}
	if err := Mgr.Register(zookeeper); err != nil {
		glog.Fatalf("%s", err)
	}
}

func localDir(p string) string {
	homeDir := serviced.ServiceDHome()
	if len(homeDir) == 0 {
		_, filename, _, _ := runtime.Caller(1)
		homeDir = path.Dir(filename)
	}
	return path.Join(homeDir, p)
}

func imagesDir() string {
	return localDir("images")
}

func resourcesDir() string {
	return localDir("resources")
}
