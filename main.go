package main

import (
	"log"
	"os"
	"runtime"

	"github.com/urfave/cli"
)

// EINVAL Error Code: #define EINVAL          22      /* Invalid argument */
const EINVAL = 23

const MaxInt32 = 2147483647
const UnLimitedTime = MaxInt32
const MissPercentageVal = 0

func main() {
	var coresCount int
	var timeSeconds int
	var percentage int

	app := cli.NewApp()

	app.Version = "0.0.1"

	app.Flags = []cli.Flag{
		cli.IntFlag{
			Name:        "coresCount, c",
			Value:       runtime.NumCPU(),
			Usage:       "how many cores",
			Destination: &coresCount,
		},
		cli.IntFlag{
			Name:        "timeSeconds, t",
			Value:       UnLimitedTime,
			Usage:       "how long",
			Destination: &timeSeconds,
		},
		cli.IntFlag{
			Name:        "percentage, p",
			Value:       MissPercentageVal,
			Usage:       "percentage of each specify cores",
			Destination: &percentage,
		},
	}

	app.Action = func(c *cli.Context) error {
		// fmt.Println("coresCount: ", coresCount)
		// fmt.Println("timeSeconds: ", timeSeconds)
		// fmt.Println("percentage: ", percentage)
		// fmt.Println("------")

		if coresCount < 1 || coresCount > runtime.NumCPU() {
			return cli.NewExitError("coresCount not correct must between 1 - `max CPU cores`", EINVAL)
		}

		if timeSeconds <= 0 {
			return cli.NewExitError("timeSeconds not correct must be positive int", EINVAL)
		}

		if percentage <= 0 || percentage > 100 {
			return cli.NewExitError("percentage must between 1 - 100", EINVAL)
		}
		RunCPULoad(coresCount, timeSeconds, percentage)
		return nil
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
