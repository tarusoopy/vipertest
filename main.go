package main

import (
	"fmt"
	flag "github.com/spf13/pflag"
	"github.com/spf13/viper"
	"os"
)

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
		DbPort string
		DbName string
	}
	PcRegister struct {
		NewPc_Endpoint      string
		IdMatch_Endpoint    string
		IpRegister_Endpoint string
		Port                string
		Logfile             string
	}
}

func main() {
	var configfile string
	var logfile string
	flag.StringVarP(
		&configfile,
		"configfile",
		"c",
		"./vsocket.ini",
		"config file path")
	flag.StringVarP(
		&logfile,
		"logfile",
		"l",
		"./flag.log",
		"log file path")
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
	viper.BindPFlag("logfile", flag.Lookup("logfile"))
	fmt.Printf("configfile: %s\n", configfile)
	//viper.AddConfigPath("/etc/variosecure")
	//viper.SetConfigType("json")
	if err := viper.ReadInConfig(); err != nil {
		return err
	}

	// 値読み込み
	if err := viper.Unmarshal(&config); err != nil {
		fmt.Println("config file Unmarshal error")
		return err
	}

	fmt.Printf("cust_id: %s\n", config.VpnInfo.Cust_id)
	fmt.Printf("vpn_network: %s\n", config.VpnInfo.Vpn_network)
	fmt.Printf("vpn_address: %s\n", config.VpnInfo.Vpn_address)
	fmt.Printf("dbuser: %s\n", config.DbInfo.DbUser)
	fmt.Printf("dbpass: %s\n", config.DbInfo.DbPass)
	fmt.Printf("dbhost: %s\n", config.DbInfo.DbHost)
	fmt.Printf("dbname: %s\n", config.DbInfo.DbName)
	fmt.Printf("newpc_endpoint: %s\n", config.PcRegister.NewPc_Endpoint)
	fmt.Printf("idmatch_endpoint: %s\n", config.PcRegister.IdMatch_Endpoint)
	fmt.Printf("ipregister_endpoint: %s\n", config.PcRegister.IpRegister_Endpoint)
	fmt.Printf("port: %s\n", config.PcRegister.Port)
	fmt.Printf("logfile: %s\n", config.PcRegister.Logfile)
	//printConfig(&config)

	return nil
}
