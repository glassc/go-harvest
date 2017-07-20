// +build integration

package harvest_test

import (
	"go-harvest/harvest"
	"flag"
	"os"

)

var client *harvest.Client


func init() {
	var account, username, password string

	flag.StringVar(&account, "account", "", "Harvest account to use")
	flag.StringVar(&username, "username", "", "Username for the Harvest account")
	flag.StringVar(&password, "password", "", "Password for the Harvest account")
	flag.Parse()
	
	if (account == "") || username == "" || password == "" {
		flag.Usage()
		os.Exit(2)
	}

	transport := harvest.BasicAuthTransport{
		Username: username,
		Password: password,
	}

	client = harvest.NewClient(account, transport.Client())

}
