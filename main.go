package main

import (
	"fmt"
	"github.com/fsnotify/fsnotify"
	flag "github.com/spf13/pflag"
	"github.com/spf13/viper"
	"os"
)

type Config struct {
	vpninfo    VpnInfo
	dbinfo     DbInfo
	pcregister PcRegister
}

type PcRegister struct {
	NewPcEndpoint      string
	IdMatchEdpoint     string
	IpRegisterEndpoint string
	Port               string
}

type VpnInfo struct {
	CustId     string `config:"cust_id"`
	VpnNetwork string `config:"vpnnetwork"`
	VpnAddress string `config:"vpnaddress"`
}

type DbInfo struct {
	DbUser string `config:"dbuser"`
	DbPass string `config:"dbpass"`
	DbHost string `config:"dbhost"`
	DbName string `config:"dbname"`
}

var config Config

func main() {
	var configfile string
	flag.StringVarP(
		&configfile,
		"configfile",
		"c",
		"/etc/variosecure/vsocket.toml",
		"config file path")
	flag.Parse()

	e := readConfig(configfile)
	if e != nil {
		os.Exit(-1)
	}
}

func readConfig(configfile string) error {
	// コンフィグ読み込み
	viper.SetConfigFile(configfile)
	viper.AddConfigPath("/etc/variosecure")
	if err := viper.ReadInConfig(); err != nil {
		fmt.Errorf("設定ファイル読み込みエラー: %s \n", err)
		return err
	}
	viper.WatchConfig()
	viper.OnConfigChange(func(e fsnotify.Event) {
		if err := viper.ReadInConfig(); err != nil {
			fmt.Errorf("設定ファイル読み込みエラー: %s \n", err)
		}
	})

	// 値読み込み
	viper.Unmarshal(&config)
	/*
		vpninfo.CustId = viper.GetString("cust_id")
		vpninfo.VpnNetwork = viper.GetString("vpn_network")
		vpninfo.VpnAddress = viper.GetString("vpn_address")

		dbinfo.DbUser = viper.GetString("dbuser")
		dbinfo.DbPass = viper.GetString("dbpass")
		dbinfo.DbHost = viper.GetString("dbhost")
		dbinfo.DbName = viper.GetString("dbname")

		pcregister.NewPcEndpoint = viper.GetString("newpc_endpoint")
		pcregister.IdMatchEdpoint = viper.GetString("idmatch_endpoint")
		pcregister.IpRegisterEndpoint = viper.GetString("ipregister_endpoint")
	*/

	return nil
}
