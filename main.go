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
		Cust_id     string
		Vpn_network string
		Vpn_address string
	}
	DbInfo struct {
		DbUser string
		DbPass string
		DbHost string
		DbName string
	}
	PcRegister struct {
		NewPc_Endpoint      string
		IdMatch_Edpoint     string
		IpRegister_Endpoint string
		Port                string
	}
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
		fmt.Println(e)
		os.Exit(-1)
	}
}

func readConfig(configfile string) error {
	var config Config

	// コンフィグ読み込み
	viper.SetConfigFile(configfile)
	fmt.Printf("configfile: %s\n", configfile)
	//viper.AddConfigPath("/etc/variosecure")
	//viper.SetConfigType("json")
	if err := viper.ReadInConfig(); err != nil {
		return err
	}
	fmt.Printf("cust_id: %s\n", viper.GetString("cust_id"))

	// 値読み込み
	if err := viper.Unmarshal(&config); err != nil {
		fmt.Println("config file Unmarshal error")
		return err
	}

	fmt.Printf("cust_id: %s\n", config.VpnInfo.Cust_id)
	fmt.Printf("vpn_network: %s\n", config.VpnInfo.Vpn_network)
	fmt.Printf("vpn_address: %s\n", config.VpnInfo.Vpn_address)
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
