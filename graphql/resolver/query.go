package resolver

type QueryResolver struct {
}

func NewRoot() (*QueryResolver, error) {
	return &QueryResolver{}, nil
}

type ResourceQueryArgs struct {
	ResourceId string
}

func (r QueryResolver) GetResource(arg ResourceQueryArgs) *ResourceResolver {
	return NewResource()
}

func (_ QueryResolver) Goodbye() string { return "Hello, world!" }
