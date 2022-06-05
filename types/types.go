package types

import (
	"go-graphql/main/resolvers"

	"github.com/graphql-go/graphql"
)

var PlayerQueryType = graphql.NewObject(graphql.ObjectConfig{
	Name: "Players",
	Fields: graphql.Fields{
		"id": &graphql.Field{
			Type: graphql.ID,
		},
		"name": &graphql.Field{
			Type: graphql.String,
		},
		"speed_attribute": &graphql.Field{
			Type: graphql.Int,
		},
		"power_attribute": &graphql.Field{
			Type: graphql.Int,
		},
		"accuracy_attribute": &graphql.Field{
			Type: graphql.Int,
		},
		"defence_attribute": &graphql.Field{
			Type: graphql.Int,
		},
		"passing_attribute": &graphql.Field{
			Type: graphql.Int,
		},
		"style": &graphql.Field{
			Type: graphql.String,
		},
		"corner_preference": &graphql.Field{
			Type: graphql.String,
		},
		"skill": &graphql.Field{
			Type: graphql.String,
		},
		"division": &graphql.Field{
			Type: graphql.String,
		},
	},
})

var playerModifierEnumType = graphql.NewEnum(graphql.EnumConfig{
	Name: "modifier",
	Values: graphql.EnumValueConfigMap{
		"BRONZE": &graphql.EnumValueConfig{
			Value: "BRONZE",
		},
		"SILVER": &graphql.EnumValueConfig{
			Value: "SILVER",
		},
		"GOLD": &graphql.EnumValueConfig{
			Value: "GOLD",
		},
		"PLATINUM": &graphql.EnumValueConfig{
			Value: "PLATINUM",
		},
		"DIAMOND": &graphql.EnumValueConfig{
			Value: "DIAMOND",
		},
	},
})

var RootQuery = graphql.NewObject(graphql.ObjectConfig{
	Name: "Query",
	Fields: graphql.Fields{
		"player": &graphql.Field{
			Type: graphql.NewList(PlayerQueryType),
			Args: graphql.FieldConfigArgument{
				"id": &graphql.ArgumentConfig{
					Type: graphql.String,
				},
				// potential future multi args
				// "name": &graphql.ArgumentConfig{
				// 	Type: graphql.String,
				// },
			},
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				return resolvers.GetPlayer(p.Args["id"].(string)), nil
			},
		},
		"players": &graphql.Field{
			Type: graphql.NewList(PlayerQueryType),
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				return resolvers.GetPlayers(), nil
			},
		},
		"generatePlayer": &graphql.Field{
			Type: PlayerQueryType,
			Args: graphql.FieldConfigArgument{
				"modifier": &graphql.ArgumentConfig{
					Description: "generate a new player, if ommited generates a basic player",
					Type:        playerModifierEnumType,
				},
			},
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				return resolvers.Generate(p.Args["modifier"]), nil
			},
		},
	},
})
