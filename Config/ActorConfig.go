package Config

import (
	"io/ioutil"
	"github.com/zllangct/RockGO/logger"
	"encoding/json"
	"errors"
)

type ClusterServerConf struct {
	Name     string
	Host     string
	RootPort int
	Http     []interface{} //[port, staticfile_path]
	Https    []interface{} //[port, certFile, keyFile, staticfile_path]
	NetPort  int
	DebugPort int //telnet port
	Remotes  []string
	Module   string
	Log      string
	Path     string
}

type ClusterConf struct {
	Master  *ClusterServerConf
	Servers map[string]*ClusterServerConf
}



/*
获取当前节点的父节点
*/
func (this *ClusterConf) GetRemotesByName(name string) ([]string, error) {
	server, ok := this.Servers[name]
	if ok {
		return server.Remotes, nil
	} else {
		return nil, errors.New("no server found!!!")
	}
}

/*
获取当前节点的子节点
*/
func (this *ClusterConf) GetChildsByName(name string) []string {
	names := make([]string, 0)
	for sername, ser := range this.Servers {
		for _, rname := range ser.Remotes {
			if rname == name {
				names = append(names, sername)
				break
			}
		}
	}
	return names
}

func (this *ClusterConf)Reload(path string){
	//集群服务器配置信息
	data, err := ioutil.ReadFile(cfgpath)
	if err != nil {
		logger.Error(err)
	}
	err = json.Unmarshal(data, this)
	if err != nil {
		logger.Error(err)
	}
}