package manager

// Target is the collection of items that
// make a tunnel unique and part of a pool
type Target struct {
	Hash byte
	Pool string
	Size int
	User string
}

func pooler() {

}

// Hasher generates a hash for the tunnel
// based on the launch arguments in the config
func Hasher() {

}

// SetTarget assembles a list of Targets
// based on the config file
func SetTarget() {

}
