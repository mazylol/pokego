# PokeGO

[![Go Reference](https://pkg.go.dev/badge/github.com/mazylol/pokego.svg)](https://pkg.go.dev/github.com/mazylol/pokego) [![Go Report Card](https://goreportcard.com/badge/github.com/mazylol/pokego)](https://goreportcard.com/report/github.com/mazylol/pokego) [![CI](https://github.com/mazylol/pokego/actions/workflows/ci.yml/badge.svg)](https://github.com/mazylol/pokego/actions/workflows/ci.yml)

PokeGO is a [Go](https://golang.org) package that provides bindings to the [Pokeapi](https://pokeapi.co). It caches to a local [SQLite](https://www.sqlite.org/index.html) database.

## Getting Started

### Installing

```sh
go get github.com/mazylol/pokego
```

### Usage

Import the package into your project.
```go
import "github.com/mazylol/pokego"
```

Use the GetPokemon function to get a Pokemon of the provided name.
```go
pokemon, err := pokego.GetPokemon("slugma")

// Print out name
fmt.Println(pokemon.Name)
```

You can also get a list of every Pokemon name.
```go
pokemonList, err := pokego.GetResourceList("pokemon", 35) // 35 is how many pokemon it will get
	
// Print out the first Pokemons name
fmt.Println(pokemonList.Results[0].Name)
```

## Features
- [x] Resource Lists/Pagination
---
- [x] Berries
- [x] Contests
- [x] Encounters
- [ ] Evolution
- [ ] Games
- [ ] Items
- [ ] Locations
- [ ] Machines
- [x] Moves
- [x] Pokemon
---
- [ ] Utility

## Contributions <3
Contributors are very welcome. I still have a lot of ground to cover and any help is greatly appreciated. Refer to the features to see what needs done.