package resolvers

import (
	"go-graphql/main/database"
	"go-graphql/main/generator"
	"go-graphql/main/schema"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func GetPlayer(id string) []schema.Players {
	ctx, collection := database.GetMongo("players")
	name := "The One"
	cursor, err := collection.Find(ctx, bson.M{"name": name})
	if err != nil {
		panic(err)
	}
	var player []schema.Players
	defer cursor.Close(ctx)
	for cursor.Next(ctx) {
		var col schema.Players
		if err = cursor.Decode(&col); err != nil {
			panic(err)
		}
		player = append(player, col)
	}

	return player
}

func GetPlayers() []schema.Players {
	ctx, collection := database.GetMongo("players")
	findOptions := options.Find()
	findOptions.SetLimit(2)
	cursor, err := collection.Find(ctx, bson.M{}, findOptions)
	if err != nil {
		panic(err)
	}
	var player []schema.Players
	defer cursor.Close(ctx)
	for cursor.Next(ctx) {
		var col schema.Players
		if err = cursor.Decode(&col); err != nil {
			panic(err)
		}
		player = append(player, col)
	}
	return player
}

func Generate(modifier interface{}) schema.Players {
	ctx, collection := database.GetMongo("players")
	division := "bronze"
	player := schema.Players{
		Name:               generator.Name(),
		Speed_attribute:    generator.Attribute(division),
		Power_attribute:    generator.Attribute(division),
		Accuracy_attribute: generator.Attribute(division),
		Defence_attribute:  generator.Attribute(division),
		Passing_attribute:  generator.Attribute(division),
		Style:              generator.Style(),
		Corner_preference:  generator.CornerPreference(),
		Skill:              "", //bronze characters have no skills
		Division:           division,
	}

	res, err := collection.InsertOne(ctx, player)
	if err != nil {
		panic(err)
	}
	player.ID = res.InsertedID.(primitive.ObjectID)
	return player
}

// 	defer func() {
// if err := client.Disconnect(context.TODO()); err != nil {
// panic(err)
// }
// }()
