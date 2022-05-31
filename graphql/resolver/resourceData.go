package resolver

type ResourceDataResolver struct {
}

func NewResourceData() *ResourceDataResolver {
	return &ResourceDataResolver{}
}

func (r *ResourceDataResolver) Name() *string {
	name := "name"
	return &name
}

func (r *ResourceDataResolver) Description() *string {
	description := "description"
	return &description
}

func (r *ResourceDataResolver) Notes() *string {
	notes := "notes"
	return &notes
}
