package main

import "github.com/graphql-go/graphql"

// Define Player type
type Player struct {
	PlayerID   int64  `json:"id"`
	Name       string `json:"name"`
	Team       string `json:"team"`
	Position   string `json:"position"`
	Rosterable bool   `json:"rosterable"`
}

// Mock data
var players = []Player{
	{
		ID:         0,
		Name:       "Nikita Kucherov",
		Team:       "Tampa Bay Lightning",
		Position:   "RW",
		Rosterable: true,
	},
	{
		ID:         1,
		Name:       "Phil Kessel",
		Team:       "Vegas Golden Knight",
		Position:   "RW",
		Rosterable: false,
	},
	{
		ID:         2,
		Name:       "Tyler Johnson",
		Team:       "Chicago Blackhawks",
		Position:   "C",
		Rosterable: true,
	},
	{
		ID:         3,
		Name:       "Jake Debrusk",
		Team:       "Boston Bruins",
		Position:   "RW",
		Rosterable: true,
	},
	{
		ID:         4,
		Name:       "Cale Makar",
		Team:       "Colorado Avalanche",
		Position:   "D",
		Rosterable: false,
	},
}

var playerType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Player",
		Fields: graphql.Fields{
			"playerID": &graphql.Field{
				Type: graphql.Int,
			},
			"name": &graphql.Field{
				Type: graphql.String,
			},
			"team": &graphql.Field{
				Type: graphql.String,
			},
			"position": &graphql.Field{
				Type: graphql.String,
			},
			"rosterable": &graphql.Field{
				Type: graphql.Boolean,
			},
		},
	},
)

var queryType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Query",
		Fields: graphql.Fields{
			// Get all current hockey players
			// /players?query={players{playerID, name, team, position, rosterable}}
			"players": &graphql.Field{
				Type:        graphql.NewList(playerType),
				Description: "Get all current players",
				Resolve: func(params graphql.ResolveParams) (interface{}, error) {
					return players, nil
				},
			},
		},
	},
)

var mutationType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Mutation",
		Fields: graphql.Fields{
			// Add player to roster
			"addPlayer": &graphql.Field{
				Type:        playerType,
				Description: "Add player to your roster",
				Args: graphql.FieldConfigArgument{
					"playerID": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.Int),
					},
				},
			},
			// Remove player from roster
		},
	},
)

func main() {
}
