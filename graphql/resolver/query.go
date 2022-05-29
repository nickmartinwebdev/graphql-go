package resolver

import (
	"github.com/google/uuid"
	"github.com/nickmartinwebdev/graphql-go/internal/resource"
)

type QueryResolver struct{
	
}

func NewRoot() (*QueryResolver, error) {
	return &QueryResolver{}, nil
}

type ResourceQueryArgs struct{
	ResourceId string
}

func (r QueryResolver) GetResource(arg ResourceQueryArgs) *resource.Resource {
	return &resource.Resource{Uuid: uuid.UUID{} }
}

func (_ QueryResolver) Goodbye() string { return "Hello, world!" }