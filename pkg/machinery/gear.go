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
func (bg BaseGear) Provide() interface{} {
	// do nothing
	return nil
}

// // Configurable is the interface to mark a gear as configurable
// type Configurable interface {
// 	Configure(config interface{})
// }

// // Gear is the Machinery main building block interface.
// // If a component want to be loaded into the app have to implemet this interface.
// type Gear interface {
// 	Name() string
// 	Start(machinery *Machinery)
// 	Provide() interface{}
// 	Shutdown()
// 	SetLogger(logger logger.Logger)
// }

// // BaseGear is the Machinery most basic building block structure.
// // If a component want to be loaded into the app should derive from this.
// type BaseGear struct {
// 	UniqueName string
// 	Logger     logger.Logger
// }

// // ConfigurableGear is a BasicGear with a config map structure.
// // However, a gear will not be configured if it does not implement
// // the Configurable interface.
// type ConfigurableGear struct {
// 	BaseGear
// 	Config map[string]interface{}
// }

// // Configure is a setter method to inject config structure in a ConfgiurableGear.
// func (cg ConfigurableGear) Configure(config interface{}) {
// 	cg.Config = config.(map[string]interface{})
// }

// // Name is the default do nothing implementation for the Gear interface Name() func.
// func (bg BaseGear) Name() string {
// 	return bg.UniqueName
// }

// // Start is the default do nothing implementation for the Gear interface Start() func.
// func (bg BaseGear) Start(machinery *Machinery) {
// 	// do nothing
// }

// // Provide is the default do nothing implementation for the Gear interface Provide() func.
// // NOTE that it returns a nil interface{}.
// func (bg BaseGear) Provide() interface{} {
// 	// do nothing
// 	return nil
// }

// // Shutdown is the default do nothing implementation for the Gear interface Shutdown() func.
// func (bg BaseGear) Shutdown() {
// 	// do nothing
// }

// // SetLogger for this Gear.
// func (bg BaseGear) SetLogger(logger logger.Logger) {
// 	bg.Logger = logger
// }

// // NewBaseGear returns a new instance of BaseGear.
// // Commodity constructor func to be used in actual gear construction
// func NewBaseGear(uname string) BaseGear {
// 	return &BaseGear{UniqueName: uname}
// }

// // NewConfigurableGear returns a new instance of ConfigurableGear.
// // Commodity constructor func to be used actual gear construction
// func NewConfigurableGear(uname string, config map[string]interface{}) ConfigurableGear {
// 	return ConfigurableGear{BaseGear: NewBaseGear(uname), Config: config}
// }
