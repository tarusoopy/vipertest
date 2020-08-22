package main

import (
	"fmt"
	flag "github.com/spf13/pflag"
	"github.com/spf13/viper"
	"os"
)

/*
type Config struct {
	vpninfo    VpnInfo
	dbinfo     DbInfo
	pcregister PcRegister
}
*/

type Config struct {
	VpnInfo struct {
		CustId     string
		VpnNetwork string
		VpnAddress string
	}
}

type PcRegister struct {
	NewPcEndpoint      string
	IdMatchEdpoint     string
	IpRegisterEndpoint string
	Port               string
}

type VpnInfo struct {
	CustId     string
	VpnNetwork string
	VpnAddress string
}

type DbInfo struct {
	DbUser string
	DbPass string
	DbHost string
	DbName string
}

func main() {
	var configfile string
	flag.StringVarP(
		&configfile,
		"configfile",
		"c",
		"./vsocket.toml",
		"config file path")
	flag.Parse()

	e := readConfig(configfile)
	if e != nil {
		fmt.Printf("%s\n", e)
		os.Exit(-1)
	}
}

func readConfig(configfile string) error {
	var vpninfo = VpnInfo{}
	var dbinfo = DbInfo{}
	var pcregister = PcRegister{}
	var config = Config{
		vpninfo:    vpninfo,
		dbinfo:     dbinfo,
		pcregister: pcregister,
	}

	// コンフィグ読み込み
	viper.SetConfigFile(configfile)
	fmt.Printf("configfile: %s\n", configfile)
	//viper.AddConfigPath("/etc/variosecure")
	//viper.SetConfigType("json")
	if err := viper.ReadInConfig(); err != nil {
		fmt.Errorf("設定ファイル読み込みエラー: %s \n", err)
		return err
	}
	fmt.Printf("cust_id: %s\n", viper.GetString("cust_id"))

	// 値読み込み
	/*
		if err := viper.Unmarshal(&config); err != nil {
			return err
		}
	*/

	config.vpninfo.CustId = viper.GetString("cust_id")
	config.vpninfo.VpnNetwork = viper.GetString("vpn_network")
	config.vpninfo.VpnAddress = viper.GetString("vpn_address")

	config.dbinfo.DbUser = viper.GetString("dbuser")
	config.dbinfo.DbPass = viper.GetString("dbpass")
	config.dbinfo.DbHost = viper.GetString("dbhost")
	config.dbinfo.DbName = viper.GetString("dbname")

	config.pcregister.NewPcEndpoint = viper.GetString("newpc_endpoint")
	config.pcregister.IdMatchEdpoint = viper.GetString("idmatch_endpoint")
	config.pcregister.IpRegisterEndpoint = viper.GetString("ipregister_endpoint")

	fmt.Printf("cust_id: %s\n", config.vpninfo.CustId)
	fmt.Printf("vpn_network: %s\n", config.vpninfo.VpnNetwork)
	fmt.Printf("vpn_address: %s\n", config.vpninfo.VpnAddress)
	fmt.Printf("dbuser: %s\n", config.dbinfo.DbUser)
	fmt.Printf("dbpass: %s\n", config.dbinfo.DbPass)
	fmt.Printf("dbhost: %s\n", config.dbinfo.DbHost)
	fmt.Printf("dbname: %s\n", config.dbinfo.DbName)
	fmt.Printf("newpc_endpoint: %s\n", config.pcregister.NewPcEndpoint)
	fmt.Printf("idmatch_endpoint: %s\n", config.pcregister.IdMatchEdpoint)
	fmt.Printf("ipregister_endpoint: %s\n", config.pcregister.IpRegisterEndpoint)
	fmt.Printf("port: %s\n", config.pcregister.Port)
	//printConfig(&config)

	return nil
}

/*
func printConfig(config *Config) {
	fmt.Printf("cust_id: %s\n", config.vpninfo.CustId)
	fmt.Printf("vpn_network: %s\n", config.vpninfo.VpnNetwork)
	fmt.Printf("vpn_address: %s\n", config.vpninfo.VpnAddress)
	fmt.Printf("dbuser: %s\n", config.dbinfo.DbUser)
	fmt.Printf("dbpass: %s\n", config.dbinfo.DbPass)
	fmt.Printf("dbhost: %s\n", config.dbinfo.DbHost)
	fmt.Printf("dbname: %s\n", config.dbinfo.DbName)
	fmt.Printf("newpc_endpoint: %s\n", config.pcregister.NewPcEndpoint)
	fmt.Printf("idmatch_endpoint: %s\n", config.pcregister.IdMatchEdpoint)
	fmt.Printf("ipregister_endpoint: %s\n", config.pcregister.IpRegisterEndpoint)
	fmt.Printf("port: %s\n", config.pcregister.Port)
}
*/
