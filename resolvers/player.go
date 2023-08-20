package resolvers

import (
	"fmt"
	"go-graphql/main/database"
	"go-graphql/main/generator"
	"go-graphql/main/schema"
	"math/rand"
	"reflect"
	"strings"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func GetPlayer(id string) schema.Players {
	ctx, collection := database.GetMongo("players")
	fmt.Println(id)
	objectId, err1 := primitive.ObjectIDFromHex(id)
	if err1 != nil {
		fmt.Println(err1)
	}
	var player schema.Players
	err := collection.FindOne(ctx, bson.M{"_id": objectId}).Decode(&player)
	if err != nil {
		panic(err)
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
	division := modifier.(string)
	player := schema.Players{
		Name:                generator.Name(),
		Speed_attribute:     generator.Attribute(division),
		Power_attribute:     generator.Attribute(division),
		Accuracy_attribute:  generator.Attribute(division),
		Defence_attribute:   generator.Attribute(division),
		Passing_attribute:   generator.Attribute(division),
		Style:               generator.Style(),
		Corner_preference:   generator.CornerPreference(),
		Skill:               generator.Skill(modifier), //bronze characters have no skills
		Division:            division,
		Current_experience:  0,
		Experience_required: 1000,
	}

	res, err := collection.InsertOne(ctx, player)
	if err != nil {
		panic(err)
	}
	player.ID = res.InsertedID.(primitive.ObjectID)
	return player
}

func Train(id string, attribute string) interface{} {
	ctx, collection := database.GetMongo("players")
	player := reflect.ValueOf(GetPlayer(id))
	formattedAttribute := fmt.Sprintf("%s_attribute", attribute)
	mongoId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return player
	}
	value, ok := player.FieldByName(formattedAttribute).Interface().(int32)
	if !ok {
		return player
	}
	var convertedCurrentExperience int32
	currentExperience := player.FieldByName("current_experience")
	fmt.Println(currentExperience)
	if currentExperience.Kind() == 0 {
		convertedCurrentExperience = 0
	} else {
		convertedCurrentExperience = currentExperience.Interface().(int32)
	}
	var updatedPlayer schema.Players
	responeErr := collection.FindOneAndUpdate(ctx,
		bson.M{"_id": mongoId},
		bson.M{
			"$set": bson.M{strings.ToLower(formattedAttribute): value + int32(rand.Intn((2-1)+2)), "current_experience": convertedCurrentExperience + 50}},
		options.FindOneAndUpdate().SetReturnDocument(options.After)).Decode(&updatedPlayer)
	if responeErr != nil {
		return player
	}
	return updatedPlayer
}

func ToLower(formattedAttribute string) {
	panic("unimplemented")
}

// 	defer func() {
// if err := client.Disconnect(context.TODO()); err != nil {
// panic(err)
// }
// }()

func AddNewAttributeToDocuments() {
	ctx, collection := database.GetMongo("players")
	collection.UpdateMany(ctx, bson.M{}, bson.M{"$set": bson.M{"experience_required": 1000,
		"current_experience": 0}})
}
