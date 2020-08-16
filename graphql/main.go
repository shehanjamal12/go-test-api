package main

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/graphql-go/graphql"
)

type mainstruc struct {
	df dbfunc
}

func  main() {
	fmt.Println("hello")
	var struc mainstruc
	itemlist := struc.df.getitem()
	var itemType = graphql.NewObject(
		graphql.ObjectConfig{
			Name: "Items",
			Fields: graphql.Fields{
				"ItemName": &graphql.Field{
					Type: graphql.String,
				},
				"ItemQuantity": &graphql.Field{

					Type: graphql.Int,
				},
			},
		},
	)

	fields := graphql.Fields{
		"tutorial": &graphql.Field{
			Type: itemType,

			Description: "Get Item By Name",

			Args: graphql.FieldConfigArgument{
				"name": &graphql.ArgumentConfig{
					Type: graphql.String,
				},
			},
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {

				name, ok := p.Args["name"].(string)
				if ok {

					for _, item := range itemlist {
						if string(item.ItemName) == name {
							// return our tutorial
							return item, nil
						}
					}
				}
				return nil, nil
			},
		},

		"list": &graphql.Field{
			Type:        itemType,
			Description: "Get all item",
			Resolve: func(params graphql.ResolveParams) (interface{}, error) {
				return itemlist, nil
			},
		},
	}
	rootQuery := graphql.ObjectConfig{Name: "RootQuery", Fields: fields}
	schemaConfig := graphql.SchemaConfig{Query: graphql.NewObject(rootQuery)}
	schema, err := graphql.NewSchema(schemaConfig)
	if err != nil {
		log.Fatalf("failed to create:%v", err)
	}
	query := `
		{
			list{
				
			}
		}
	`
	params := graphql.Params{Schema: schema, RequestString: query}
	r := graphql.Do(params)
	if len(r.Errors) > 0 {
		log.Fatalf("failed :%v", r.Errors)
	}
	rJSON, _ := json.Marshal(r)
	fmt.Printf("%s \n", rJSON)

}
