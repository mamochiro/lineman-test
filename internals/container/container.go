package container

import (
	"fmt"

	"lm-test/internals/controller"
	"lm-test/internals/infrastructure/database"
	"lm-test/internals/repository"
	"lm-test/internals/service"
	"lm-test/internals/utils"

	"lm-test/internals/config"

	client "lm-test/internals/infrastructure/http_client"
	server "lm-test/internals/infrastructure/http_server"

	"go.uber.org/dig"
)

// Container ...
type Container struct {
	container *dig.Container
}

// Configure ...
func (c *Container) Configure() error {
	// config
	if err := c.container.Provide(config.NewConfiguration); err != nil {
		return err
	}

	// server
	if err := c.container.Provide(server.NewHttpServer); err != nil {
		return err
	}
	// client
	if err := c.container.Provide(client.NewHttpClient); err != nil {
		return err
	}

	// controller
	if err := c.container.Provide(controller.NewHealthZController); err != nil {
		return err
	}
	if err := c.container.Provide(controller.NewPingPongController); err != nil {
		return err
	}
	if err := c.container.Provide(controller.NewCovidController); err != nil {
		return err
	}

	// database
	if err := c.container.Provide(database.NewServerBase); err != nil {
		return err
	}

	// repository
	if err := c.container.Provide(repository.NewRepository); err != nil {
		return err
	}
	if err := c.container.Provide(repository.NewCovidRepository); err != nil {
		return err
	}

	// service
	if err := c.container.Provide(service.NewService); err != nil {
		return err
	}
	if err := c.container.Provide(service.NewCovidService); err != nil {
		return err
	}

	// utils
	if err := c.container.Provide(utils.NewUtils); err != nil {
		return err
	}
	if err := c.container.Provide(utils.NewCustomValidator); err != nil {
		return err
	}

	// If have new dependency should be set here
	return nil
}

// Start ...
func (c *Container) Start() error {
	fmt.Println("Start Container")
	if err := c.container.Invoke(func(s *server.HttpServer) {
		s.Start()
	}); err != nil {
		fmt.Printf("%s", err)
		return err
	}
	return nil
}

// NewContainer ...
func NewContainer() (*Container, error) {
	fmt.Println("this file should be set dependency injection by using uber-dig")
	d := dig.New()
	container := &Container{
		container: d,
	}
	if err := container.Configure(); err != nil {
		return nil, err
	}
	return container, nil
}
