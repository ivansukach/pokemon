package services

import (
	"github.com/ivansukach/pokemon/repositories/abilities"
	"github.com/ivansukach/pokemon/repositories/abilitiesDescription"
	"github.com/ivansukach/pokemon/repositories/arena"
	"github.com/ivansukach/pokemon/repositories/pokemonDescription"
	"github.com/ivansukach/pokemon/repositories/users"
	"github.com/ivansukach/pokemon/repositories/usersProperties"
)

type GameService struct { //I shouldn`t write ..Service
	ur  users.Repository
	upr usersProperties.Repository //afgahaasjs???
	pdr pokemonDescription.Repository
	arr arena.Repository
	adr abilitiesDescription.Repository
	abr abilities.Repository
}

func (gs *GameService) CreateUser(user *users.User) error {
	return gs.ur.Create(user)
}
func (gs *GameService) UpdateUser(user *users.User) error {
	return gs.ur.Update(user)
}
func (gs *GameService) GetUser(id string) (*users.User, error) {
	return gs.ur.Get(id)
}
func (gs *GameService) DeleteUser(id string) error {
	return gs.ur.Delete(id)
}
func (gs *GameService) ListingUsers() ([]users.User, error) {
	return gs.ur.Listing()
}

func (gs *GameService) CreateAbility(ability *abilities.Ability) error {
	return gs.abr.Create(ability)
}
func (gs *GameService) GetAbility(name string) (*abilities.Ability, error) {
	return gs.abr.Get(name)
}
func (gs *GameService) DeleteAbility(name string) error {
	return gs.abr.Delete(name)
}

func (gs *GameService) CreateAbilityDescription(description *abilitiesDescription.Description) error {
	return gs.adr.Create(description)
}
func (gs *GameService) GetAbilityDescription(name string) (*abilitiesDescription.Description, error) {
	return gs.adr.Get(name)
}
func (gs *GameService) UpdateAbilityDescription(description *abilitiesDescription.Description) error {
	return gs.adr.Update(description)
}
func (gs *GameService) DeleteAbilityDescription(name string) error {
	return gs.adr.Delete(name)
}

func (gs *GameService) CreateArena(arena *arena.Arena) (int64, error) {
	return gs.arr.Create(arena)
}
func (gs *GameService) GetArena(id string) (*arena.Arena, error) {
	return gs.arr.Get(id)
}
func (gs *GameService) UpdateArena(arena *arena.Arena) error {
	return gs.arr.Update(arena)
}
func (gs *GameService) DeleteArena(id string) error {
	return gs.arr.Delete(id)
}

func (gs *GameService) CreatePokemon(pokemon *pokemonDescription.Pokemon) error { //maybe rename db table from pokemonDescription to pokemon
	return gs.pdr.Create(pokemon)
}
func (gs *GameService) GetPokemon(name string) (*pokemonDescription.Pokemon, error) {
	return gs.pdr.Get(name)
}
func (gs *GameService) UpdatePokemon(pokemon *pokemonDescription.Pokemon) error {
	return gs.pdr.Update(pokemon)
}
func (gs *GameService) DeletePokemon(name string) error {
	return gs.pdr.Delete(name)
}
func (gs *GameService) ListingPokemons() ([]pokemonDescription.Pokemon, error) {
	return gs.pdr.Listing()
}

func (gs *GameService) CreateProperty(property *usersProperties.Properties) error {
	return gs.upr.Create(property)
}
func (gs *GameService) DeleteProperty(id string) error {
	return gs.upr.Delete(id)
}
func (gs *GameService) GetAllProperties(login string) ([]usersProperties.Properties, error) {
	return gs.upr.GetAll(login)
}

func New(users users.Repository, usersProperties usersProperties.Repository, pokemonDescription pokemonDescription.Repository,
	arena arena.Repository, abilitiesDescription abilitiesDescription.Repository, abilities abilities.Repository) *GameService {
	return &GameService{ur: users, upr: usersProperties, pdr: pokemonDescription, arr: arena, adr: abilitiesDescription, abr: abilities}
}
