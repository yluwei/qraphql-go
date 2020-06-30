package schema

import (
	"github.com/graphql-go/graphql"
	"github.com/graphql-go/handler"
)

// 定义跟查询节点
var rootQuery = graphql.NewObject(graphql.ObjectConfig{
	Name:        "RootQuery",
	Description: "Root Query",
	Fields: graphql.Fields{
		"hello": &queryHello, // queryHello 参考schema/hello.go
	},
})

// 定义跟查询节点
var root1Query = graphql.NewObject(graphql.ObjectConfig{
	Name:        "RootQuery",
	Description: "Root Query",
	Fields: graphql.Fields{
		"hello1": &queryHello, // queryHello 参考schema/hello.go
	},
})

// 定义Schema用于http handler处理
var schema, _ = graphql.NewSchema(graphql.SchemaConfig{
	Query:    rootQuery,
	Mutation: nil, // 需要通过GraphQL更新数据，可以定义Mutation
})

// 定义Schema用于http handler处理
var schema1, _ = graphql.NewSchema(graphql.SchemaConfig{
	Query:    root1Query,
	Mutation: nil, // 需要通过GraphQL更新数据，可以定义Mutation
})

var HelloHandler = handler.New(&handler.Config{
	Schema:   &schema,
	Pretty:   true,
	GraphiQL: true,
})

var Hello1Handler = handler.New(&handler.Config{
	Schema:   &schema1,
	Pretty:   true,
	GraphiQL: true,
})
