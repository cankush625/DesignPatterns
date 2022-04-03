package object_pool_design_pattern

type connection struct {
	id string
}

func (conn *connection) getID() string {
	return conn.id
}
