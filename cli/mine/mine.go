package mine

import (
	"fmt"
	"errors"
	"strconv"
	"strings"

	. "github.com/elastos/Elastos.ELA.Client/rpc"

	"github.com/urfave/cli"
)

func miningAction(c *cli.Context) error {
	if c.NumFlags() == 0 {
		cli.ShowSubcommandHelp(c)
		return nil
	}

	if action := c.String("toggle"); action != "" {
		action = strings.ToLower(action)
		if action != "start" && action != "stop" {
			return errors.New("toggle argument must be [start, stop]")
		}

		result, err := CallAndUnmarshal("togglemining", Param("mining", action))
		if err != nil {
			return err
		}

		fmt.Println(result)
		return nil
	}

	if num := c.String("number"); num != "" {
		number, err := strconv.ParseInt(num, 10, 16)
		if err != nil || number < 1 {
			return errors.New("[number] must be a positive integer")
		}
		result, err := CallAndUnmarshal("manualmining", Param("count", number))
		if err != nil {
			return err
		}

		fmt.Println(result)
		return nil
	}

	return nil
}

func NewCommand() *cli.Command {
	return &cli.Command{
		Name:        "mine",
		Usage:       "toggle cpu mining or manual mine",
		Description: "With ela-cli mine, you can toggle cpu mining, or manual mine blocks.",
		ArgsUsage:   "[args]",
		Flags: []cli.Flag{
			cli.StringFlag{
				Name:  "toggle, t",
				Usage: "use --toggle [start, stop] to toggle cpu mining",
			},
			cli.StringFlag{
				Name:  "number, n",
				Usage: "user --number [number] to mine the given number of blocks",
			},
		},
		Action: miningAction,
		OnUsageError: func(c *cli.Context, err error, isSubcommand bool) error {
			return cli.NewExitError(err, 1)
		},
	}
}
