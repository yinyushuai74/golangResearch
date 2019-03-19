package user

import "fmt"

func (c Status) String() string {
	switch c {
	case Offline:
		return "Offline"
	case Online:
		return "Online"
	case Disable:
		return "Disable"
	case Deleted:
		return "Deleted"
	}
	return fmt.Sprintf("Status(%d)", c)
}
