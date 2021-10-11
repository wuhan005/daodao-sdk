// Copyright 2021 E99p1ant. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package main

import (
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/pkg/errors"
	"github.com/urfave/cli/v2"
	log "unknwon.dev/clog/v2"

	"github.com/wuhan005/daodao-sdk/daodao"
)

func main() {
	defer log.Stop()
	err := log.NewConsole()
	if err != nil {
		panic(err)
	}

	app := cli.NewApp()
	app.Name = "daodao"

	app.Flags = []cli.Flag{
		&cli.StringFlag{
			Name:     "phone",
			Value:    "",
			EnvVars:  []string{"DAODAO_PHONE"},
			Required: true,
		},
		&cli.StringFlag{
			Name:     "password",
			Value:    "",
			EnvVars:  []string{"DAODAO_PASSWORD"},
			Required: true,
		},
	}

	app.Commands = []*cli.Command{
		check,
		plan,
	}
	if err := app.Run(os.Args); err != nil {
		log.Fatal("Failed to start application: %v", err)
	}
}

var check = &cli.Command{
	Name: "check",
	Action: func(c *cli.Context) error {
		phone := c.String("phone")
		password := c.String("password")

		client := daodao.NewClient()
		_, err := client.Login(phone, password)
		if err != nil {
			return errors.Wrap(err, "login")
		}

		account, err := client.SyncV2(daodao.Account)
		if err != nil {
			return errors.Wrap(err, "get daodao account")
		}

		for _, account := range account.Data.Account[0].Lists {
			name := account.BankName
			if name == "" || name == "自定义" {
				name = account.Name
			}
			log.Trace("[%s] - ￥%s", name, account.Money)
		}

		return nil
	},
}

var plan = &cli.Command{
	Name: "plan",
	Flags: []cli.Flag{
		&cli.StringFlag{
			Name: "accounts",
		},
		&cli.IntFlag{
			Name: "to",
		},
	},
	Action: func(c *cli.Context) error {
		phone := c.String("phone")
		password := c.String("password")

		planAccounts := c.String("accounts")
		planTo := c.Int("to")

		planAccountsSet := make(map[string]float64)
		for _, account := range strings.Split(planAccounts, ",") {
			planAccountsSet[account] = 0
		}

		client := daodao.NewClient()
		_, err := client.Login(phone, password)
		if err != nil {
			return errors.Wrap(err, "login")
		}

		account, err := client.SyncV2(daodao.Account)
		if err != nil {
			return errors.Wrap(err, "get daodao account")
		}

		for _, account := range account.Data.Account[0].Lists {
			name := account.BankName
			if name == "" || name == "自定义" {
				name = account.Name
			}

			if _, ok := planAccountsSet[name]; !ok {
				continue
			}
			money, err := strconv.ParseFloat(account.Money, 64)
			if err != nil {
				continue
			}

			planAccountsSet[name] = money
			log.Trace("[%s] - ￥%s", name, account.Money)
		}

		var totalAmount float64
		for _, money := range planAccountsSet {
			totalAmount += money
		}

		now := time.Now()
		var restDay int
		var resetDate time.Time
		if now.Day() < planTo {
			restDay = planTo - now.Day()
			resetDate = now.AddDate(0, 0, restDay)
		} else {
			restDay = int(now.AddDate(0, 1, 0).Sub(now).Hours()/24) - (now.Day() - planTo)
			resetDate = now.AddDate(0, 0, restDay)
		}

		maxCostPerDay := totalAmount / float64(restDay)

		log.Trace("Balance: ￥%.2f", totalAmount)
		log.Trace("%d days to %d/%d, ￥%.2f per day.", restDay, resetDate.Month(), resetDate.Day(), maxCostPerDay)
		return nil
	},
}
