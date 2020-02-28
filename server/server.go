package server

import (
	"context"
	"github.com/ivansukach/pokemon/protocol"
	"github.com/ivansukach/pokemon/repositories/arena"
	"github.com/ivansukach/pokemon/repositories/pokemonDescription"
	"github.com/ivansukach/pokemon/repositories/usersProperties"
	"github.com/ivansukach/pokemon/services"
	log "github.com/sirupsen/logrus"
)

type Server struct {
	gs *services.Game
}

func New(gs *services.Game) *Server {
	return &Server{gs: gs}
}
func (s *Server) Create(ctx context.Context, req *protocol.Pokemon) (*protocol.EmptyResponse, error) {
	pokemon := pokemonDescription.Pokemon{
		Name:          req.Name,
		Image:         req.Image,
		Price:         req.Price,
		HealthPoints:  req.HealthPoints,
		HealthRegen:   req.HealthRegen,
		Mana:          req.Mana,
		ManaRegen:     req.ManaRegen,
		Armor:         req.Armor,
		Damage:        req.Damage,
		MovementSpeed: req.MovementSpeed,
		AttackSpeed:   req.AttackSpeed,
		AttackRange:   req.AttackRange,
	}
	err := s.gs.CreatePokemon(&pokemon)
	if err != nil {
		log.Error(err)
		return &protocol.EmptyResponse{}, err
	}
	return &protocol.EmptyResponse{}, nil
}
func (s *Server) GetByName(ctx context.Context, req *protocol.GetByNameRequest) (*protocol.Pokemon, error) {
	pokemon, err := s.gs.GetPokemon(req.Name)
	if err != nil {
		log.Error(err)
		return nil, err
	}
	response := &protocol.Pokemon{
		Name:          pokemon.Name,
		Image:         pokemon.Image,
		Price:         pokemon.Price,
		HealthPoints:  pokemon.HealthPoints,
		HealthRegen:   pokemon.HealthRegen,
		Mana:          pokemon.Mana,
		ManaRegen:     pokemon.ManaRegen,
		Armor:         pokemon.Armor,
		Damage:        pokemon.Damage,
		MovementSpeed: pokemon.MovementSpeed,
		AttackSpeed:   pokemon.AttackSpeed,
		AttackRange:   pokemon.AttackRange,
	}
	return response, nil

}
func (s *Server) GetById(ctx context.Context, req *protocol.GetByIdRequest) (*protocol.Pokemon, error) {
	name, err := s.gs.GetPokemonNameById(req.Id)
	if err != nil {
		log.Error(err)
		return nil, err
	}
	pokemon, err := s.GetByName(ctx, &protocol.GetByNameRequest{Name: name}) //isn't variable dead when exit statement??
	return pokemon, nil

}
func (s *Server) Delete(ctx context.Context, req *protocol.DeleteRequest) (*protocol.EmptyResponse, error) {
	err := s.gs.DeletePokemon(req.Name)
	if err != nil {
		log.Error(err)
		return &protocol.EmptyResponse{}, err
	}
	return &protocol.EmptyResponse{}, nil
}
func (s *Server) Update(ctx context.Context, req *protocol.Pokemon) (*protocol.EmptyResponse, error) {
	pokemon := pokemonDescription.Pokemon{
		Name:          req.Name,
		Image:         req.Image,
		Price:         req.Price,
		HealthPoints:  req.HealthPoints,
		HealthRegen:   req.HealthRegen,
		Mana:          req.Mana,
		ManaRegen:     req.ManaRegen,
		Armor:         req.Armor,
		Damage:        req.Damage,
		MovementSpeed: req.MovementSpeed,
		AttackSpeed:   req.AttackSpeed,
		AttackRange:   req.AttackRange,
	}
	err := s.gs.UpdatePokemon(&pokemon)
	if err != nil {
		log.Error(err)
		return &protocol.EmptyResponse{}, err
	}
	return &protocol.EmptyResponse{}, nil
}
func (s *Server) Buy(ctx context.Context, req *protocol.BuyRequest) (*protocol.EmptyResponse, error) {
	user, err := s.gs.GetUser(req.Login)
	if err != nil {
		log.Error(err)
		return nil, err
	}
	coins := user.Coins
	pokemon, err := s.gs.GetPokemon(req.Name)
	if err != nil {
		log.Error(err)
		return nil, err
	}
	if coins < pokemon.Price {
		log.Warning("You haven't got enough money")
		return nil, nil
	}
	user.Coins = user.Coins - pokemon.Price
	err = s.gs.UpdateUser(user)
	if err != nil {
		log.Error(err)
		return nil, err
	}
	property := usersProperties.Properties{IdOfPokemon: 0, Login: user.Login, Name: pokemon.Name}
	err = s.gs.CreateProperty(&property)
	if err != nil {
		log.Error(err)
		return nil, err
	}
	return &protocol.EmptyResponse{}, nil
}
func (s *Server) Sell(ctx context.Context, req *protocol.SellRequest) (*protocol.EmptyResponse, error) {
	user, err := s.gs.GetUser(req.Login)
	if err != nil {
		log.Error(err)
		return nil, err
	}
	pokemonName, err := s.gs.GetPokemonNameById(req.Id)
	if err != nil {
		log.Error(err)
		return nil, err
	}
	pokemon, err := s.gs.GetPokemon(pokemonName)
	if err != nil {
		log.Error(err)
		return &protocol.EmptyResponse{}, err
	}
	user.Coins = user.Coins + pokemon.Price
	err = s.gs.UpdateUser(user)
	if err != nil {
		log.Error(err)
		return nil, err
	}
	err = s.gs.DeleteProperty(req.Id)
	if err != nil {
		log.Error(err)
		return nil, err
	}
	return &protocol.EmptyResponse{}, nil
}
func (s *Server) MakeDamage(ctx context.Context, req *protocol.DamageRequest) (*protocol.EmptyResponse, error) {
	arena, err := s.gs.GetArena(req.IdArena)
	if err != nil {
		log.Error(err)
		return nil, err
	}
	if req.IdOfDamageReceiver == arena.Fighter1 {
		arena.Health1 -= req.DamageAmount
	} else {
		arena.Health2 -= req.DamageAmount
	}
	err = s.gs.UpdateArena(arena)
	if err != nil {
		log.Error(err)
		return nil, err
	}
	return &protocol.EmptyResponse{}, nil

}
func (s *Server) CreateArena(ctx context.Context, req *protocol.CreateArenaRequest) (*protocol.ArenaId, error) {
	pokemonName1, err := s.gs.GetPokemonNameById(req.Id1)
	if err != nil {
		log.Error(err)
		return nil, err
	}
	pokemonName2, err := s.gs.GetPokemonNameById(req.Id2)
	if err != nil {
		log.Error(err)
		return nil, err
	}
	pokemon1, err := s.gs.GetPokemon(pokemonName1)
	if err != nil {
		log.Error(err)
		return nil, err
	}
	pokemon2, err := s.gs.GetPokemon(pokemonName2)
	if err != nil {
		log.Error(err)
		return nil, err
	}
	arena := arena.Arena{
		Id:       0,
		Fighter1: req.Id1,
		Fighter2: req.Id2,
		Health1:  pokemon1.HealthPoints,
		Health2:  pokemon2.HealthPoints,
		X1:       0,
		Y1:       0,
		X2:       0,
		Y2:       0,
		Winner:   "",
	}
	idArena, err := s.gs.CreateArena(&arena)
	if err != nil {
		log.Error(err)
		return nil, err
	}
	return &protocol.ArenaId{IdArena: idArena}, nil
}
