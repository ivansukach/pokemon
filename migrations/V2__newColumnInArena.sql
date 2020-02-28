ALTER TABLE arena ADD COLUMN Id SERIAL PRIMARY KEY;
CREATE TABLE pokemonProperties(
    IdOfPokemon SERIAL REFERENCES usersProperties,
    Id TEXT REFERENCES attributes
);
