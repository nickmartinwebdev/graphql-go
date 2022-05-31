package resolver

type ResourceResolver struct {
}

func NewResource() *ResourceResolver {
	return &ResourceResolver{}
}

func (r *ResourceResolver) Uuid() string {
	id := "uuid"
	return id
}

func (r *ResourceResolver) Versions() *[]*ResourceVersionResolver {
	return NewResourceVersions()
}
