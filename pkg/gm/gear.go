package gm

// Configurable is the interface to mark a gear as configurable
type Configurable interface {
	Configure(config interface{})
}

// Gear is the Machinery main building block interface.
// If a component want to be loaded into the app have to implemet this interface.
type Gear interface {
	Name() string
	Start(machinery *Machinery)
	Provide() interface{}
	Shutdown()
}

// BaseGear is the Machinery most basic building block structure.
// If a component want to be loaded into the app should derive from this.
type BaseGear struct {
	UniqueName string
}

// ConfigurableGear is a BasicGear with a config map structure.
// However, a gear will not be configured if it does not implement
// the Configurable interface.
type ConfigurableGear struct {
	BaseGear
	Config map[string]interface{}
}

// Name is the default do nothing implementation for the Gear interface Name() func.
func (bg *BaseGear) Name() string {
	return bg.UniqueName
}

// Start is the default do nothing implementation for the Gear interface Start() func.
func (bg *BaseGear) Start(machinery *Machinery) {
	// do nothing
}

// Provide is the default do nothing implementation for the Gear interface Provide() func.
// NOTE that it returns a nil interface{}.
func (bg *BaseGear) Provide() interface{} {
	// do nothing
	return nil
}

// Shutdown is the default do nothing implementation for the Gear interface Shutdown() func.
func (bg *BaseGear) Shutdown() {
	// do nothing
}

// NewBaseGear returns a new instance of BaseGear.
// Commodity constructor func to be used in actual gear construction
func NewBaseGear(uname string) BaseGear {
	return BaseGear{UniqueName: uname}
}

// NewConfigurableGear returns a new instance of ConfigurableGear.
// Commodity constructor func to be used actual gear construction
func NewConfigurableGear(uname string, config map[string]interface{}) ConfigurableGear {
	return ConfigurableGear{BaseGear: NewBaseGear(uname), Config: config}
}
