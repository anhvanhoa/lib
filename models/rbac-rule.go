package models

type RbacRule struct {
	Id       int
	Name     string
	Path     string
	Method   string
	Status   bool
	Roles    []int `pg:",array"`
	AuthType string
	Service  string
}
