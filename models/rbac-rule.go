package models

type RbacRule struct {
	Id       int    `json:"id"`
	Name     string `json:"name"`
	Path     string `json:"path"`
	Method   string `json:"method"`
	Status   bool   `json:"status"`
	Roles    []int  `pg:",array" json:"roles"`
	AuthType string `json:"auth_type"`
	Service  string `json:"service"`
}
