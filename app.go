package main

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"os/user"
	"path/filepath"

	uuid "github.com/satori/go.uuid"
)

var ConfigName = "easy-host.conf"

type H interface{}

type M map[string]interface{}

type Config struct {
	Host []*Host `json:"host"`
}
type Host struct {
	Identity string `json:"identity"`
	Name     string `json:"name"`
	Info     string `json:"info"`
	Use      string `json:"use"`
}

func HostList() ([]*Host, error) {
	nowPath := GetConfPath()
	data, err := ioutil.ReadFile(nowPath + string(os.PathSeparator) + ConfigName)
	if errors.Is(err, os.ErrNotExist) {
		return nil, errors.New("暂无数据")
	}
	conf := new(Config)
	err = json.Unmarshal(data, conf)
	if err != nil {
		return nil, err
	}
	return conf.Host, nil
}

func HostCreate(host *Host) error {
	conf := new(Config)
	nowPath := GetConfPath()
	if host.Identity != "1" {
		host.Identity = uuid.NewV4().String()
	}
	data, err := ioutil.ReadFile(nowPath + string(os.PathSeparator) + ConfigName)
	if errors.Is(err, os.ErrNotExist) {
		conf.Host = []*Host{host}
		data, _ = json.Marshal(conf)
		os.MkdirAll(nowPath, 0666)
		ioutil.WriteFile(nowPath+string(os.PathSeparator)+ConfigName, data, 0666)
		return nil
	}
	json.Unmarshal(data, conf)
	conf.Host = append(conf.Host, host)
	data, _ = json.Marshal(conf)
	ioutil.WriteFile(nowPath+string(os.PathSeparator)+ConfigName, data, 0666)
	return nil
}

func HostEdit(host *Host) error {
	conf := new(Config)
	nowPath := GetConfPath()
	data, err := ioutil.ReadFile(nowPath + string(os.PathSeparator) + ConfigName)
	if err != nil {
		return err
	}
	json.Unmarshal(data, conf)
	for i, v := range conf.Host {
		if v.Identity == host.Identity {
			conf.Host[i] = host
		} else {
			v.Use = "0"
		}
	}
	data, _ = json.Marshal(conf)
	ioutil.WriteFile(nowPath+string(os.PathSeparator)+ConfigName, data, 0666)
	return nil
}

func HostDelete(identity string) error {
	if identity == "" {
		return errors.New("标识不能为空")
	}
	conf := new(Config)
	nowPath := GetConfPath()
	data, err := ioutil.ReadFile(nowPath + string(os.PathSeparator) + ConfigName)
	if err != nil {
		return err
	}
	err = json.Unmarshal(data, conf)
	if err != nil {
		return err
	}
	for i, v := range conf.Host {
		if v.Identity == identity {
			conf.Host = append(conf.Host[:i], conf.Host[i+1:]...)
			break
		}
	}
	data, _ = json.Marshal(conf)
	ioutil.WriteFile(nowPath+string(os.PathSeparator)+ConfigName, data, 0666)
	return nil
}

func GetConfPath() string {
	current, _ := user.Current()
	return current.HomeDir
}

func HostsInit() error {
	userHome := getUserHome()
	easyHostConfig := filepath.Join(userHome, string(os.PathSeparator), ConfigName)
	// 判断文件夹是否存在，不存在创建
	if exist, _ := fileExists(userHome); !exist {
		os.MkdirAll(userHome, os.ModePerm)
	}
	hostsPath := getHostPath()
	data, err := ioutil.ReadFile(hostsPath)
	if err != nil {
		data, _ = json.Marshal("")
		ioutil.WriteFile(easyHostConfig, data, os.ModePerm)
	}
	haveSystemHost := false
	hostlist, _ := HostList()
	for i, v := range hostlist {
		if v.Identity == "1" {
			println(i)
			haveSystemHost = true
		}
	}
	if !haveSystemHost {
		host := new(Host)
		host.Identity = "1"

		host.Name = "系统默认"
		host.Info = string(data)
		host.Use = "1"
		HostCreate(host)
	}
	print(hostsPath)
	hosts = hostsPath
	return nil
}

func WriteHost(host *Host) error {
	print(hosts)
	setSystemHosts(host.Info)
	return nil
}

func fileExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

// App struct
type App struct {
	ctx context.Context
}

// NewApp creates a new App application struct
func NewApp() *App {
	return &App{}
}

// startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (a *App) startup(ctx context.Context) {
	a.ctx = ctx
}

// Greet returns a greeting for the given name
func (a *App) Greet(name string) string {
	return fmt.Sprintf("Hello %s, It's show time!", name)
}

func (a *App) HostList() H {
	conn, err := HostList()
	if err != nil {
		return M{
			"code": -1,
			"msg":  "ERROR : " + err.Error(),
		}
	}
	return M{
		"code": 200,
		"data": conn,
	}
}

func (a *App) HostCreate(host *Host) H {
	err := HostCreate(host)
	if err != nil {
		return M{
			"code": -1,
			"msg":  "ERROR : " + err.Error(),
		}
	}
	return M{
		"code": 200,
		"msg":  "新建成功",
	}
}

func (a *App) HostEdit(host *Host) H {
	println(host.Info)
	println(host.Name)
	err := HostEdit(host)
	if err != nil {
		return M{
			"code": -1,
			"msg":  "ERROR : " + err.Error(),
		}
	}
	return M{
		"code": 200,
		"msg":  "编辑成功",
	}
}
func (a *App) WriteHost(host *Host) H {

	err := WriteHost(host)
	if err != nil {
		return M{
			"code": -1,
			"msg":  "ERROR : " + err.Error(),
		}
	}
	return M{
		"code": 200,
		"msg":  "切换成功",
	}
}

func (a *App) HostDelete(identity string) H {
	err := HostDelete(identity)
	if err != nil {
		return M{
			"code": -1,
			"msg":  "ERROR : " + err.Error(),
		}
	}
	return M{
		"code": 200,
		"msg":  "删除成功",
	}
}
