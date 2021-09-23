package schema

import (
	"github.com/graphql-go/graphql"
	graphql2 "graphql-demo/model"
	"math/rand"
	"time"
)

var rootQuery = graphql.NewObject(graphql.ObjectConfig{
	Name:        "RootQuery",
	Description: "Root Query",
	Fields: graphql.Fields{
		"hello": &queryDemo,
	},
})

var newsType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "News",
		Fields: graphql.Fields{
			"id": &graphql.Field{
				Type: graphql.Int,
			},
			"title": &graphql.Field{
				Type: graphql.String,
			},
			"content": &graphql.Field{
				Type: graphql.String,
			},
			"author": &graphql.Field{
				Type: graphql.String,
			},
		},
	},
)

var mutation = graphql.NewObject(graphql.ObjectConfig{
	Name: "Mutation",
	Fields: graphql.Fields{
		"create": &graphql.Field{
			Type:        newsType,
			Description: "Create new news",
			Args: graphql.FieldConfigArgument{
				"title": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(graphql.String),
				},
				"content": &graphql.ArgumentConfig{
					Type: graphql.String,
				},
				"author": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(graphql.String),
				},
			},
			Resolve: func(params graphql.ResolveParams) (interface{}, error) {
				rand.Seed(time.Now().UnixNano())
				news := graphql2.News{
					Id:      rand.Intn(100000), // generate random ID
					Title:   params.Args["title"].(string),
					Content: params.Args["content"].(string),
					Author:  params.Args["author"].(string),
				}
				graphql2.NewsList = append(graphql2.NewsList, news)
				return graphql2.NewsList, nil
			},
		},
		"update": &graphql.Field{
			Type:        newsType,
			Description: "Update news by id",
			Args: graphql.FieldConfigArgument{
				"title": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(graphql.String),
				},
				"content": &graphql.ArgumentConfig{
					Type: graphql.String,
				},
				"author": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(graphql.String),
				},
			},
			Resolve: func(params graphql.ResolveParams) (interface{}, error) {
				id, _ := params.Args["id"].(int)
				title, titleOk := params.Args["title"].(string)
				content, contentOk := params.Args["content"].(string)
				author, authorOk := params.Args["author"].(string)
				for i, p := range graphql2.NewsList {
					if id == p.Id {
						if titleOk {
							graphql2.NewsList[i].Title = title
						}
						if contentOk {
							graphql2.NewsList[i].Content = content
						}
						if authorOk {
							graphql2.NewsList[i].Author = author
						}
						break
					}
				}
				return graphql2.NewsList, nil
			},
		},
		"delete": &graphql.Field{
			Type:        newsType,
			Description: "Delete product by id",
			Args: graphql.FieldConfigArgument{
				"id": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(graphql.Int),
				},
			},
			Resolve: func(params graphql.ResolveParams) (interface{}, error) {
				id, _ := params.Args["id"].(int)
				for i, p := range graphql2.NewsList {
					if id == p.Id {
						graphql2.NewsList = append(graphql2.NewsList[:i], graphql2.NewsList[i+1:]...)
					}
				}

				return graphql2.NewsList, nil
			},
		},
	},
})

var Schema, _ = graphql.NewSchema(graphql.SchemaConfig{
	Query:    rootQuery,
	Mutation: mutation, // 需要通过GraphQL更新数据
})
