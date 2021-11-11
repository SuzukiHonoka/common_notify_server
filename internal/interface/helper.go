package iface

type Helper interface {
	Connect()
	Create()
	Close()
	Refresh()
	GetUsers() interface{}
	AddUser(usr interface{})
	DelUser(usr interface{})
}
