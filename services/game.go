package services

import (
	"github.com/ivansukach/pokemon/repositories/abilities"
	"github.com/ivansukach/pokemon/repositories/abilitiesDescription"
	"github.com/ivansukach/pokemon/repositories/arena"
	"github.com/ivansukach/pokemon/repositories/pokemonDescription"
	"github.com/ivansukach/pokemon/repositories/users"
	"github.com/ivansukach/pokemon/repositories/usersProperties"
)

type Game struct {
	userRepository                 users.Repository
	userPropertiesRepository       usersProperties.Repository
	pokemonDescriptionRepository   pokemonDescription.Repository
	arenaRepository                arena.Repository
	abilitiesDescriptionRepository abilitiesDescription.Repository
	abilitiesRepository            abilities.Repository
}

func (gs *Game) CreateUser(user *users.User) error {
	return gs.userRepository.Create(user)
}
func (gs *Game) UpdateUser(user *users.User) error {
	return gs.userRepository.Update(user)
}
func (gs *Game) GetUser(login string) (*users.User, error) {
	return gs.userRepository.Get(login)
}
func (gs *Game) DeleteUser(id string) error {
	return gs.userRepository.Delete(id)
}
func (gs *Game) ListingUsers() ([]users.User, error) {
	return gs.userRepository.Listing()
}

func (gs *Game) CreateAbility(ability *abilities.Ability) error {
	return gs.abilitiesRepository.Create(ability)
}
func (gs *Game) GetAbility(name string) (*abilities.Ability, error) {
	return gs.abilitiesRepository.Get(name)
}
func (gs *Game) DeleteAbility(name string) error {
	return gs.abilitiesRepository.Delete(name)
}

func (gs *Game) CreateAbilityDescription(description *abilitiesDescription.Description) error {
	return gs.abilitiesDescriptionRepository.Create(description)
}
func (gs *Game) GetAbilityDescription(name string) (*abilitiesDescription.Description, error) {
	return gs.abilitiesDescriptionRepository.Get(name)
}
func (gs *Game) UpdateAbilityDescription(description *abilitiesDescription.Description) error {
	return gs.abilitiesDescriptionRepository.Update(description)
}
func (gs *Game) DeleteAbilityDescription(name string) error {
	return gs.abilitiesDescriptionRepository.Delete(name)
}

func (gs *Game) CreateArena(arena *arena.Arena) (int64, error) {
	return gs.arenaRepository.Create(arena)
}
func (gs *Game) GetArena(id int64) (*arena.Arena, error) {
	return gs.arenaRepository.Get(id)
}
func (gs *Game) UpdateArena(arena *arena.Arena) error {
	return gs.arenaRepository.Update(arena)
}
func (gs *Game) DeleteArena(id int64) error {
	return gs.arenaRepository.Delete(id)
}

func (gs *Game) CreatePokemon(pokemon *pokemonDescription.Pokemon) error { //maybe rename db table from pokemonDescription to pokemon
	return gs.pokemonDescriptionRepository.Create(pokemon)
}
func (gs *Game) GetPokemon(name string) (*pokemonDescription.Pokemon, error) {
	return gs.pokemonDescriptionRepository.Get(name)
}
func (gs *Game) UpdatePokemon(pokemon *pokemonDescription.Pokemon) error {
	return gs.pokemonDescriptionRepository.Update(pokemon)
}
func (gs *Game) DeletePokemon(name string) error {
	return gs.pokemonDescriptionRepository.Delete(name)
}
func (gs *Game) ListingPokemons() ([]pokemonDescription.Pokemon, error) {
	return gs.pokemonDescriptionRepository.Listing()
}

func (gs *Game) CreateProperty(property *usersProperties.Properties) error {
	return gs.userPropertiesRepository.Create(property)
}
func (gs *Game) DeleteProperty(id int64) error {
	return gs.userPropertiesRepository.Delete(id)
}
func (gs *Game) GetAllProperties(login string) ([]usersProperties.Properties, error) {
	return gs.userPropertiesRepository.GetAll(login)
}
func (gs *Game) GetPokemonNameById(id int64) (string, error) {
	return gs.userPropertiesRepository.GetNameById(id)
}

func New(users users.Repository, usersProperties usersProperties.Repository, pokemonDescription pokemonDescription.Repository,
	arena arena.Repository, abilitiesDescription abilitiesDescription.Repository, abilities abilities.Repository) *Game {
	return &Game{userRepository: users, userPropertiesRepository: usersProperties, pokemonDescriptionRepository: pokemonDescription, arenaRepository: arena, abilitiesDescriptionRepository: abilitiesDescription, abilitiesRepository: abilities}
}
