package admin

import "github.com/a-h/templ"

type Service[Data any] interface {
	FindOne(id string) (Data, error)
	FindMany(filter Data) ([]Data, error)
	UpdateOne(id string, data Data) (Data, error)
	CreateOne(data Data) (Data, error)
	DeleteOne(id string) (bool, error)
}

type Module[Data any] struct {
	Service       Service[Data]
	Item          templ.Component
	ItemContainer *templ.Component
	ListItem      templ.Component
	ListContainer templ.Component
	PageContainer templ.Component
}
