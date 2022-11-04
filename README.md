# Fantasy Hockey
> A tool for managing my fantasy hockey team

### Features
- The ability to (`GET` REST, `query` GQL) all current NHL players
- The ability to add (`POST`/`PUT` REST, `mutation` GQL) an NHL player to a roster
- The ability to remove (`PATCH`/`DELETE` REST, `mutation` GQL) a NHL player from a roster

### Design

#### Schema

```graphql
schema {
  query: Query
  mutation: Mutation
}

type Player {
  playerID: ID
  name: String
  team: String
  position: String
  rostered: Boolean
}

type Query {
  players: [Player]
} 

type Mutation {
  addPlayer(name: String, team: String, position: String, rostered: Boolean): Player
  removePlayer(playerID: ID): Boolean
}i
```

#### API

##### Query GetAllPlayers

```graphql
query GetPlayers {
  players {
    name
    team
    position
    rostered
  }
}
```

```json
{
  "data": {
    "players": [
      {
        "playerID": 42069GOAT
        "name": "Nikita Kucherov",
        "team": "Tampa Bay Lightning",
        "position": "RW",
        "rostered": true
      },
    ]
  }
}
```

##### Mutation AddPlayer

```graphql
mutation AddPlayerToRoster {
  addPlayer(name: "Nikita Kucherov", team: "Tampa Bay Lightning", position: "RW", rostered: true) {
    playerID
    name
    team
    position
  }
}
```

```json
{
  "data": {
    "addPlayer": {
      "playerID": 42069GOAT,
      "name": "Nikita Kucherov",
      "team": "Tampa Bay Lightning",
      "position": "RW",
      "rostered": true 
    }
  }
}
```
##### Mutation RemovePlayer

```graphql
mutation RemovePlayerFromRoster {
  removePlayer(playerID: 42069GOAT) {
    rostered
  }
}
```

```json
{
  "data": {
    "rostered": false
  }
}
```

### Tools and technologies

Go: 

GraphQL:
https://pkg.go.dev/github.com/graphql-go/graphql#section-readme
https://github.com/graphql-go/graphql 

Svelte:

Mongo:
https://www.mongodb.com/languages/golang