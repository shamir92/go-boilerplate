package commands

import (
	"fmt"
	usecase "simple-invitation/domain/usecase"
	"strconv"

	"github.com/urfave/cli/v2"
)

type pingCommands struct {
	pingUsecase usecase.IPingUsecase
}

type IPingCommands interface {
	// PingController Interface
	GetPingController(c *cli.Context) error
	GetPingUseCase(c *cli.Context) error
	GetPingRepository(c *cli.Context) error
	//
}

func NewPingCommand() *pingCommands {
	var pingUsecase usecase.IPingUsecase = usecase.NewPingUsecase()
	return &pingCommands{
		pingUsecase: pingUsecase,
	}
}

func (ac *pingCommands) GetPingController(c *cli.Context) error {
	fmt.Println("pong " + strconv.FormatBool(true))
	return nil
}

func (ac *pingCommands) GetPingUseCase(c *cli.Context) error {

	value, _ := ac.pingUsecase.PingUsecase()
	fmt.Println("pong " + strconv.FormatBool(value))
	return nil
}

func (ac *pingCommands) GetPingRepository(c *cli.Context) error {
	value, _ := ac.pingUsecase.PingRepository()
	// log.Println(att)
	fmt.Println("pong " + strconv.FormatBool(value))
	return nil
}
