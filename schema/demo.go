package schema

import (
	"graphql-demo/model"

	"github.com/graphql-go/graphql"
)

var demoType = graphql.NewObject(graphql.ObjectConfig{
	Name: "Hello",
	Description: "Hello Model",
	Fields: graphql.Fields{
		"id": &graphql.Field{
			Type: graphql.Int,
		},
		"name": &graphql.Field{
			Type: graphql.String,
		},
	},
})

var queryDemo = graphql.Field{
	Name: "QueryHello",
	Description: "Query Hello",
	Type: graphql.NewList(demoType),
	Args: graphql.FieldConfigArgument{
		"id": &graphql.ArgumentConfig{
			Type: graphql.Int,
		},
		"name": &graphql.ArgumentConfig{
			Type: graphql.String,
		},
	},
	Resolve: func(p graphql.ResolveParams) (result interface{}, err error) {
		id, _ := p.Args["id"].(int)
		name, _ := p.Args["name"].(string)
		
		return (&model.Hello{}).Query(id, name)
	},
}