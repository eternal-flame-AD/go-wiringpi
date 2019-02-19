package wiringpi

// SetupMethod enumberates setup methods
// used to call Setup function
type SetupMethod int

const (
	// WiringPiSetup setup GPIO using WiringPi scheme, requires root
	WiringPiSetup SetupMethod = iota
	// BroadcomSetup setup GPIO using Broadcom scheme, requires root
	BroadcomSetup
	// PhysSetup setup GPIO using Broadcom scheme with only P1 access, requires root
	PhysSetup
	// SysSetup setup GPIO using Broadcom scheme, uses only exported ports
	SysSetup
)
