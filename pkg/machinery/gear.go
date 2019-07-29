package machinery

import (
	"github.com/go-mach/machinery/pkg/logger"
)

// Gear is the Machinery main building block interface.
type Gear interface {
	Name() string
	Start(*Machinery)
	Provide() interface{}
	Shutdown()
	SetLogger(logger logger.Logger)
	Configure(config interface{})
}

// BaseGear .
type BaseGear struct {
	Uname  string
	Config interface{}
	Logger logger.Logger
}

// Name .
func (bg *BaseGear) Name() string {
	return bg.Uname
}

// Shutdown .
func (bg *BaseGear) Shutdown() {
	bg.Logger.Printf("%s went down", bg.Uname)
}

// Start is the default do nothing implementation for the Gear interface Start() func.
func (bg *BaseGear) Start(machinery *Machinery) {
	bg.Logger.Fatalf("Please, provide a Start() method implementation for the %s gear", bg.Uname)
}

// Provide is the default do nothing implementation for the Gear interface Provide() func.
// NOTE that it returns a nil interface{}.
func (bg *BaseGear) Provide() interface{} {
	return nil
}

// SetLogger default Logger setter method. It should remain
// not overridden in concrete gears.
func (bg *BaseGear) SetLogger(logger logger.Logger) {
	bg.Logger = logger
}

// Configure default config setter method. It could be overridden
//in concrete gears to work on config type.
func (bg *BaseGear) Configure(config interface{}) {
	bg.Config = config
}
