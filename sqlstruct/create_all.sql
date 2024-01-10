CREATE TABLE IF NOT EXISTS Language (
    language TEXT,
    PRIMARY KEY(language)
);

CREATE TABLE IF NOT EXISTS Country (
    id SERIAL,
    name TEXT NOT NULL UNIQUE,
    capital TEXT NOT NULL,
    population BIGINT NOT NULL,
    area FLOAT8 NOT NULL,
    currency CHAR(3) NOT NULL,
    language TEXT REFERENCES Language(language),
    climate TEXT,
    PRIMARY KEY(id)
);
CREATE TABLE IF NOT EXISTS Hotel (
    hotel_id SERIAL,
    name TEXT NOT NULL,
    address TEXT NOT NULL,
    rating SMALLINT CHECK (rating >= 1 AND rating <= 5), -- Assuming the rating is between 0 and 5
    country_id integer REFERENCES Country(id) ON DELETE CASCADE NOT NULL,
    PRIMARY KEY (hotel_id)
);

CREATE TABLE IF NOT EXISTS Hotel_room (
    room_id INTEGER,
    hotel_id INTEGER REFERENCES Hotel(hotel_id) ON DELETE CASCADE NOT NULL,
    places_number SMALLINT NOT NULL,
    price_per_night MONEY,
    PRIMARY KEY(room_id, hotel_id)
);




CREATE TABLE IF NOT EXISTS Tourist (
    id SERIAL,
    first_name TEXT NOT NULL,
    last_name TEXT NOT NULL,
    age SMALLINT NOT NULL check (age >= 18),
    gender BOOLEAN NOT NULL, 
    citizenship_id INTEGER REFERENCES Country(id) ON DELETE CASCADE NOT NULL,
    login TEXT UNIQUE NOT NULL,
    password_hash bytea NOT NULL,
    PRIMARY KEY (id)
);

CREATE TABLE IF NOT EXISTS Tourist_hotel (
    tourist_id INTEGER REFERENCES Tourist(id) ON DELETE CASCADE NOT NULL,
    start_date DATE,
    hotel_id INTEGER REFERENCES Hotel(hotel_id) ON DELETE CASCADE NOT NULL,
    room_id INTEGER NOT NULL,
    end_date DATE,
    CONSTRAINT hotel_room_ref FOREIGN KEY (hotel_id, room_id) REFERENCES Hotel_room(hotel_id, room_id) ON DELETE CASCADE,
    PRIMARY KEY (tourist_id, start_date)
);

CREATE TABLE IF NOT EXISTS Agency (
    id SERIAL,
    name TEXT NOT NULL,
    address TEXT NOT NULL,
    contact_number TEXT NOT NULL,
    PRIMARY KEY(id)
);

CREATE TABLE IF NOT EXISTS Tour (
    id SERIAL,
    name TEXT NOT NULL,
    travel_agency_id INTEGER REFERENCES Agency(id) ON DELETE CASCADE NOT NULL,
    start_date DATE NOT NULL,
    end_date DATE check (end_date >= start_date) not null,
    satisfaction_level float8 CHECK (satisfaction_level >= 1 AND satisfaction_level <= 5), 
    price MONEY NOT NULL,
    PRIMARY KEY(id)
);

CREATE TABLE IF NOT EXISTS Tourist_tour (
    tourist_id INTEGER REFERENCES Tourist(id) ON DELETE CASCADE NOT NULL,
    tour_id INTEGER REFERENCES Tour(id) ON DELETE CASCADE NOT NULL,
    PRIMARY KEY(tourist_id, tour_id)
);

CREATE TABLE IF NOT EXISTS Review_src_type (
    type TEXT,
    PRIMARY KEY(type)
);

CREATE TABLE IF NOT EXISTS Review_src (
    id SERIAL,
    name TEXT NOT NULL,
    type TEXT REFERENCES Review_src_type(type) ON DELETE CASCADE NOT NULL,
    address TEXT NOT NULL,
    PRIMARY KEY(id)
);

CREATE TABLE IF NOT EXISTS Tour_review (
    tourist_id INTEGER REFERENCES Tourist(id) ON DELETE CASCADE NOT NULL,
    tour_id INTEGER REFERENCES Tour(id) ON DELETE CASCADE NOT NULL,
    review_text TEXT,
    rating SMALLINT NOT NULL CHECK (rating >= 1 AND rating <= 5), -- Assuming the rating is between 0 and 5
    datetime TIMESTAMP NOT NULL,
    review_src_id INTEGER REFERENCES Review_src(id) ON DELETE CASCADE NOT NULL,
    PRIMARY KEY(tourist_id, tour_id)
);

CREATE TABLE IF NOT EXISTS Guide (
    id SERIAL,
    first_name TEXT NOT NULL,
    last_name TEXT NOT NULL,
    age INTEGER NOT NULL,
    gender BOOLEAN NOT NULL,
    agency_id INTEGER REFERENCES Agency(id) ON DELETE CASCADE NOT NULL,
    PRIMARY KEY(id)
);

CREATE TABLE IF NOT EXISTS Guide_lang (
    guide_id INTEGER REFERENCES Guide(id) ON DELETE CASCADE NOT NULL,
    language TEXT REFERENCES Language(language) ON DELETE CASCADE NOT NULL,
    PRIMARY KEY(guide_id, language)
);

CREATE TABLE IF NOT EXISTS Guide_tour (
    guide_id INTEGER REFERENCES Guide(id) ON DELETE CASCADE NOT NULL,
    tour_id INTEGER REFERENCES Tour(id) ON DELETE CASCADE NOT NULL,
    PRIMARY KEY(guide_id, tour_id)
);

CREATE TABLE IF NOT EXISTS Transport (
    type TEXT,
    PRIMARY KEY(type)
);

CREATE TABLE IF NOT EXISTS Trip (
    id SERIAL,
    src_country_id INTEGER REFERENCES Country(id) ON DELETE CASCADE NOT NULL,
    dst_country_id INTEGER REFERENCES Country(id) ON DELETE CASCADE NOT NULL,
    transport_type TEXT REFERENCES Transport(type) ON DELETE CASCADE NOT NULL,
    datetime_start TIMESTAMP NOT NULL,
    datetime_end TIMESTAMP check(datetime_end >= datetime_start) NOT NULL,
    price MONEY NOT NULL,
    tour_id INTEGER REFERENCES Tour(id) ON DELETE CASCADE NOT NULL,
    PRIMARY KEY(id)
);

CREATE TABLE IF NOT EXISTS Visa (
    id SERIAL,
    name TEXT NOT NULL,
    destination_country_id INTEGER REFERENCES Country(id) ON DELETE CASCADE NOT NULL,
    tourist_id INTEGER REFERENCES Tourist(id) ON DELETE CASCADE NOT NULL,
    issue_date DATE NOT NULL,
    expiry_date DATE NOT NULL,
    PRIMARY KEY(id)
);

CREATE TABLE IF NOT EXISTS Attraction (
    id SERIAL,
    name TEXT NOT NULL,
    address TEXT NOT NULL,
    description TEXT,
    ticket_price MONEY,
    country_id INTEGER REFERENCES Country(id) ON DELETE CASCADE NOT NULL,
    PRIMARY KEY(id)
);

CREATE TABLE IF NOT EXISTS Tourist_attraction (
    tourist_id INTEGER REFERENCES Tourist(id) ON DELETE CASCADE NOT NULL,
    attraction_id INTEGER REFERENCES Attraction(id) ON DELETE CASCADE NOT NULL,
    PRIMARY KEY(tourist_id, attraction_id)
);
