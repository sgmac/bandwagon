package bandwagon

// Server interface does basic actions for a given vps.
type VirtualServer interface {
	// Lists available operating systems.
	ListImages() (*Images, error)
}
