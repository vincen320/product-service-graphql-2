package resolver

type (
	RootResolver     struct{}
	QueryResolver    struct{}
	MutationResolver struct{}
)

func (*RootResolver) Query() *QueryResolver {
	return &QueryResolver{}
}

func (*RootResolver) Mutation() *MutationResolver {
	return &MutationResolver{}
}
