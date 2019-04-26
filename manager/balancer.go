package manager

// Metadata is the collection of items that
// make a tunnel unique and part of a pool
type Metadata struct {
	Hash byte
	Pool string
	Size int
	User string
}

// Hasher takes in the launch args from the
// config to produce a unique hash representing the tunnel or pool
func hasher() {

}

// SetMetadata assembles a list of Targets
// based on the config file
func SetMetadata() {

}
