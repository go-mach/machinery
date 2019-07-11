package gm

import (
	"fmt"
	"os"
	"os/signal"
	"strings"
	"syscall"
	"time"

	"github.com/go-mach/gm/config"
)

// Machinery is the main framework structure.
type Machinery struct {
	gears        map[string]Gear
	GracefulStop chan os.Signal
	Logger       Logger
}

// NewMachinery initialize and return the main Machinery engine instance.
func NewMachinery() *Machinery {
	logger := NewLogger()
	theGoMachinery := &Machinery{
		gears:        make(map[string]Gear),
		GracefulStop: make(chan os.Signal),
		Logger:       logger,
	}

	signal.Notify(theGoMachinery.GracefulStop, syscall.SIGTERM)
	signal.Notify(theGoMachinery.GracefulStop, syscall.SIGINT)

	go func() {
		sig := <-theGoMachinery.GracefulStop
		logger.Printf("caught sig: %+v", sig)

		logger.Println("Wait for 2 second to finish processing")
		time.Sleep(2 * time.Second)

		theGoMachinery.Shutdown()
		logger.Println("All gears went down. Shutting down the Machinery.")
		logger.Println("Bye!")
		os.Exit(0)
	}()

	return theGoMachinery
}

// With and configure one or more Gears with the Machinery engine.
func (m *Machinery) With(gears ...Gear) *Machinery {
	var gearName string

	for _, gear := range gears {
		gearName = gear.Name()
		if m.gears[gearName] != nil {
			m.Logger.Printf("Gear %s already registered", gearName)
		} else {
			m.Logger.Printf("registering %s Gear", gearName)
			m.gears[gearName] = gear
		}
	}

	return m
}

// Start configure app gears and starts the machinery
func (m *Machinery) Start() {
	m.Logger.Println("configuring machinery gears")
	m.configureGears()

	m.Logger.Println("starting machinery gears")
	m.startGears()

	m.Logger.Println("app Machinery started")
	select {}
}

// Shutdown all of the registered gears.
func (m *Machinery) Shutdown() {
	m.Logger.Println("Shutting down the Machinery")
	for gearName, gear := range m.gears {
		m.Logger.Printf("shutting down the %s gear", gearName)
		gear.Shutdown()
	}
}

// configure configurable gears
func (m *Machinery) configureGears() {
	for gearName, gear := range m.gears {
		// check if the gear is Configurable
		if configurableGear, ok := gear.(Configurable); ok {
			m.Logger.Printf("the %s gear is configurable", gearName)
			gearConfig := config.Get(strings.ToLower(gearName))
			if gearConfig == nil {
				panic(fmt.Sprintf("no configuration found for gear %s", gearName))
			}
			configurableGear.Configure(config.Get(gearName))
		}
	}
}

func (m *Machinery) startGears() {
	for gearName, gear := range m.gears {
		m.Logger.Printf("starting the %s gear", gearName)
		gear.Start(m)
	}
}

// GetGear returns a Gear instance pointer
// TODO: use a map to store Gears
func (m *Machinery) GetGear(name string) Gear {
	return m.gears[name]
}
