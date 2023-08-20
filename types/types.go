package types

import (
	"fmt"
	"go-graphql/main/resolvers"
	"sync"

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
		"experience_required": &graphql.Field{
			Type: graphql.Int,
		},
		"current_experience": &graphql.Field{
			Type: graphql.Int,
		},
		"image_url": &graphql.Field{
			Type: graphql.String,
		},
	},
})

var playerImageType = graphql.NewObject(graphql.ObjectConfig{
	Name: "Players",
	Fields: graphql.Fields{
		"id": &graphql.Field{
			Type: graphql.ID,
		},
		"url": &graphql.Field{
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

var playerAttributeEnumType = graphql.NewEnum(graphql.EnumConfig{
	Name: "attribute",
	Values: graphql.EnumValueConfigMap{
		"speed": &graphql.EnumValueConfig{
			Value: "Speed",
		},
		"power": &graphql.EnumValueConfig{
			Value: "Power",
		},
		"accuracy": &graphql.EnumValueConfig{
			Value: "Accuracy",
		},
		"defence": &graphql.EnumValueConfig{
			Value: "Defence",
		},
		"passing": &graphql.EnumValueConfig{
			Value: "Passing",
		},
	},
})

var RootQuery = graphql.NewObject(graphql.ObjectConfig{
	Name: "Query",
	Fields: graphql.Fields{
		"player": &graphql.Field{
			Type: PlayerQueryType,
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
		"train": &graphql.Field{
			Type: PlayerQueryType,
			Args: graphql.FieldConfigArgument{
				"id": &graphql.ArgumentConfig{
					Description: "train an existing player, if ommited returns nothing",
					Type:        graphql.String,
				},
				"attribute": &graphql.ArgumentConfig{
					Description: "attribute that needs to be trained, if ommited returns nothing",
					Type:        playerAttributeEnumType,
				},
			},
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				return resolvers.Train(p.Args["id"].(string), p.Args["attribute"].(string)), nil
			},
		},
		"image": &graphql.Field{
			Type: playerImageType,
		},

		"goroutine": &graphql.Field{
			Type: graphql.String,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				resultChan := make(chan int)
				go ConcurrentFibonacci(103, resultChan)
				// go FibonacciMemoization(1021124123, resultChan)
				result := <-resultChan
				fmt.Printf("Fibonacci(%d) = %d\n", 1021124123, result)
				// Simulate a time-consuming operation
				// For demonstration purposes, we use a sleep here.
				// In a real-world scenario, you might have a database query or other complex operation.
				// The use of goroutine here allows other requests to be processed while this one is sleeping.
				// wg := sync.WaitGroup{}
				// wg.Add(1)
				// go func() {
				// 	defer wg.Done()
				// 	// Simulate a time-consuming operation
				// 	// In a real-world application, you would replace this with your actual logic.
				// 	// For example, fetching data from a database, calling an external API, etc.
				// 	// Here, we just sleep for 2 seconds to simulate a time-consuming task.
				// 	// Replace this with your actual logic.
				// 	// The result will be sent back to the client when this goroutine completes.
				// 	// Other incoming requests can still be processed in the meantime.
				// 	fmt.Println("Start long operation...")
				// 	// Simulate a time-consuming operation by sleeping for 2 seconds.
				// 	// In a real-world scenario, replace this with your actual logic.
				// 	// For example, fetching data from a database or making an external API call.
				// 	// The use of goroutine here allows other requests to be processed concurrently.
				// 	FibonacciRecursion(1241)
				// 	fmt.Println("Long operation completed.")
				// }()

				// We don't wait for the goroutine to complete here,
				// so the endpoint will return immediately while the long operation continues in the background.
				return "We're working it out", nil
			},
		},
	},
})

func FibonacciIterative(n int) int {
	if n <= 1 {
		return n
	}

	prev := 0
	current := 1

	for i := 2; i <= n; i++ {
		// Calculate the next Fibonacci number
		next := prev + current

		// Update the previous and current values for the next iteration
		prev = current
		current = next
	}

	return current
}

func ConcurrentFibonacci(n int, resultChan chan<- int) {
	resultChan <- FibonacciIterative(n)
}

var memo = make(map[int]int)
var mutex = sync.Mutex{} // Mutex to synchronize access to the memo map

func FibonacciMemoization(n int, resultChan chan<- int) {
	// Check if the value is already in the memo
	mutex.Lock()
	if val, ok := memo[n]; ok {
		mutex.Unlock()
		resultChan <- val
		return
	}
	mutex.Unlock()

	// Calculate the value recursively and store it in the memo
	if n <= 1 {
		memo[n] = n
	} else {
		// Calculate the Fibonacci value recursively
		// and send the result through the channel
		leftChan := make(chan int)
		rightChan := make(chan int)

		go FibonacciMemoization(n-1, leftChan)
		go FibonacciMemoization(n-2, rightChan)

		left := <-leftChan
		right := <-rightChan

		memo[n] = left + right
	}

	// Send the result through the channel
	resultChan <- memo[n]
}
