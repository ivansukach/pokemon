CREATE TABLE pokemonDescription(
    Name TEXT PRIMARY KEY,
    Image TEXT,
    Price INTEGER,
    HealthPoints INTEGER,
    HealthRegen INTEGER,
    Mana INTEGER,
    ManaRegen INTEGER,
    Armor REAL,
    Damage INTEGER,
    MovementSpeed INTEGER,
    AttackSpeed INTEGER,
    AttackRange INTEGER
);
CREATE TABLE users(
    Login TEXT PRIMARY KEY,
    Password TEXT,
    Name TEXT,
    Surname TEXT,
    Coins INTEGER
);
CREATE TABLE attributes(
    Id TEXT PRIMARY KEY,
    Name TEXT,
    Damage REAL,
    Armor REAL,
    Health REAL
);
CREATE TABLE abilities(
    Name TEXT REFERENCES pokemonDescription,
    Electroshock bool,
    Flamethrower bool,
    Waterflow bool
);
CREATE TABLE abilitiesDescription(
    Name TEXT PRIMARY KEY,
    Animation TEXT,
    Damage INTEGER,
    Mana INTEGER,
    AttackRange INTEGER
);



CREATE TABLE usersProperties(
    IdOfPokemon SERIAL PRIMARY KEY,
    Login TEXT REFERENCES users,
    Name TEXT REFERENCES pokemonDescription
);
CREATE TABLE arena(
    Fighter1 TEXT,
    Fighter2 TEXT,
    Health1 INTEGER,
    Health2 INTEGER,
    X1 REAL,
    Y1 REAL,
    X2 REAL,
    Y2 REAL,
    Winner TEXT
);