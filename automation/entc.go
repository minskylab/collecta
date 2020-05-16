package main

func main() {
	entSchema := "./ent/schema"
	gqlgenRoot := "./api"

	generateGQLGenFromEntSchema(entSchema, gqlgenRoot)
	inflateGQLGen(gqlgenRoot)
}