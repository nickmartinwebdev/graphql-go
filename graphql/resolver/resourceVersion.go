package resolver

type ResourceVersionResolver struct {
}

func NewResourceVersion() *ResourceVersionResolver {
	return &ResourceVersionResolver{}
}

func NewResourceVersions() *[]*ResourceVersionResolver {
	return &[]*ResourceVersionResolver{}
}

func (r *ResourceVersionResolver) VersionId() string {
	id := "versionId"
	return id
}

func (r *ResourceVersionResolver) CreatedAt() string {
	version := "created at"
	return version
}

func (r *ResourceVersionResolver) ResourceData() *ResourceDataResolver {
	return NewResourceData()
}
