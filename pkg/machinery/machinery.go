package machinery

import (
	"fmt"
	"os"
	"os/signal"
	"strings"
	"syscall"
	"time"

	"github.com/go-mach/machinery/pkg/logger"

	"github.com/go-mach/machinery/pkg/config"
)

// Machinery is the main framework structure.
type Machinery struct {
	gears        map[string]Gear
	GracefulStop chan os.Signal
	Logger       logger.Logger
}

// NewMachinery initialize and return the main Machinery engine instance.
func NewMachinery() *Machinery {
	// create the Machinery
	theLogger := logger.NewLogger(config.GetConfiguration().Log)
	theGoMachinery := &Machinery{
		gears:        make(map[string]Gear),
		GracefulStop: make(chan os.Signal),
		Logger:       theLogger,
	}

	// set up os signal notifications
	signal.Notify(theGoMachinery.GracefulStop, syscall.SIGTERM)
	signal.Notify(theGoMachinery.GracefulStop, syscall.SIGINT)

	// start a go-func to trigger os signals and gently shutdown the Machinery
	go func() {
		sig := <-theGoMachinery.GracefulStop
		theLogger.Printf("caught sig: %+v", sig)

		theLogger.Println("Wait for 2 second to finish processing")
		time.Sleep(2 * time.Second)

		theGoMachinery.Shutdown()
		theLogger.Println("All gears went down. Shutting down the Machinery.")
		theLogger.Println("Bye!")
		os.Exit(0)
	}()

	return theGoMachinery
}

// With register one or more Gears with the Machinery engine.
func (m *Machinery) With(gears ...Gear) *Machinery {
	var gearName string

	for _, gear := range gears {
		gearName = gear.Name()
		if m.gears[gearName] != nil {
			m.Logger.Printf("Gear %s already registered", gearName)
		} else {
			m.Logger.Printf("registering %s Gear", gearName)
			gear.SetLogger(m.Logger)
			m.gears[gearName] = gear
		}
	}

	return m
}

// Start starts up the Machinery.
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

// configure configurable gears.
func (m *Machinery) configureGears() {
	for gearName, gear := range m.gears {
		// check if the gear is Configurable
		//if configurableGear, ok := gear.(Configurable); ok {
		m.Logger.Printf("the %s gear is configurable", gearName)
		gearConfig := config.Get(strings.ToLower(gearName))
		if gearConfig == nil {
			panic(fmt.Sprintf("no configuration found for gear %s", gearName))
		}
		m.Logger.Printf("found configuration for %s gear: %v", gearName, gearConfig)
		gear.Configure(config.Get(gearName))
		//configurableGear.Configure(config.Get(gearName))
		//}
	}
}

// starts all the configured gears.
func (m *Machinery) startGears() {
	for gearName, gear := range m.gears {
		m.Logger.Printf("starting the %s gear", gearName)
		gear.Start(m)
	}
}

// GetGear returns a Gear instance pointer
func (m *Machinery) GetGear(name string) Gear {
	gear := m.gears[name]
	return gear
}
