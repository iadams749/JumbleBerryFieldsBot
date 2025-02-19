package game

const (
	// different berry types
	Pest Berry = iota
	Jumbleberry
	Sugarberry
	Pickleberry
	Moonberry
)

type Berry int

// ANSI RGB color codes used for printing berries to the console
const (
	grey   = "\033[38;2;169;169;169m" // Dark Gray (RGB)
	red    = "\033[38;2;255;0;0m"     // Bright Red (RGB)
	yellow = "\033[38;2;255;255;0m"   // True Yellow (RGB)
	green  = "\033[38;2;0;255;0m"     // Bright Green (RGB)
	purple = "\033[38;2;148;0;211m"   // Deep Purple (RGB)
	reset  = "\033[0m"                // Reset color
)

// String() method for Berry with color formatting
func (b Berry) String() string {
	switch b {
	case Pest:
		return grey + "PEST" + reset
	case Jumbleberry:
		return red + "JBRY" + reset
	case Sugarberry:
		return yellow + "SBRY" + reset
	case Pickleberry:
		return green + "PBRY" + reset
	case Moonberry:
		return purple + "MBRY" + reset
	default:
		return "Unknown Berry"
	}
}
