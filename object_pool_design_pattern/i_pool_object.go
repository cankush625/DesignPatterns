package object_pool_design_pattern

type IPoolObject interface {
	// ID is required for comparing two objects
	getID() string
}
