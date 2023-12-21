# Pokedex Go cli app
A project curated by [boot.dev](https://boot.dev)

Go cli app - a repl for catching pokemon info from [Pokeapi](https://pokeapi.co)

## Commands
```
help
exit
map
mapb
explore
catch
inspect
pokedex
```

## Diff
- A cache holds unmarshalled structs. Generics were used to allow hold any struct in cache map.
- Api layer doesn't know about cache. It just get an url, fetches data and unmarshall it to a struct. Check and save to cache is in command(controller) layer.
- A bit different file/folder structure
- made 1/3 chance without a treshhold of a catch failure just for faster tests

## IDEAS FOR EXTENDING THE PROJECT
- Update the CLI to support the "up" arrow to cycle through previous commands
- Simulate battles between pokemon
- Add more unit tests
- Refactor your code to organize it better and make it more testable
- Keep pokemon in a "party" and allow them to level up
- Allow for pokemon that are caught to evolve after a set amount of time
- Persist a user's Pokedex to disk so they can save progress between sessions
- Use the PokeAPI to make exploration more interesting. For example, rather than typing the names of areas, maybe you are given choices of areas and just type "left" or "right"
- Random encounters with wild pokemon
- Adding support for different types of balls (Pokeballs, Great Balls, Ultra Balls, etc), which have different chances of catching pokemon