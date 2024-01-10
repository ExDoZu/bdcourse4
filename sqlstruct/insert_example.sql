-- Insert data into the Language table
INSERT INTO Language (language) VALUES
    ('English'),
    ('Spanish'),
    ('French'),
    ('German');

-- Insert data into the Country table
INSERT INTO Country (name, capital, population, area, currency, language, climate) VALUES
    ('USA', 'Washington, D.C.', 331002651, 9833517.85, 'USD', 'English', 'Various'),
    ('France', 'Paris', 65273511, 551695.33, 'EUR', 'French', 'Temperate'),
    ('Spain', 'Madrid', 46754778, 505992.00, 'EUR', 'Spanish', 'Mediterranean'),
    ('Germany', 'Berlin', 83783942, 357022.00, 'EUR', 'German', 'Temperate');

-- Insert data into the Hotel table
INSERT INTO Hotel (name, address, rating, country_id) VALUES
    ('Grand Hotel', '100 Grand St, City', 5, 1),
    ('Beach Resort', '200 Beach Rd, Serenity', 4, 2),
    ('City Hotel', '300 Urban Ave, Metropolis', 3, 3),
    ('Alpine Lodge', '400 Mountain Rd, Tranquility', 2, 4);


-- Insert data into the Hotel_room table
INSERT INTO Hotel_room (room_id, hotel_id, places_number, price_per_night) VALUES
    (101, 1, 2, 200.00),
    (102, 1, 1, 150.00),
    (201, 2, 3, 100.00),
    (202, 2, 2, 80.00),
    (301, 3, 1, 250.00),
    (302, 3, 2, 180.00),
    (401, 4, 2, 120.00),
    (402, 4, 1, 90.00);

-- Insert data into the Tourist table
INSERT INTO Tourist (first_name, last_name, age, gender, citizenship_id, login, password_hash) VALUES
    ('John', 'Smith', 25, true, 1, 'johnsmith', '123456'),
    ('Jane', 'Doe', 30, false, 2, 'janedoe', '123456'),
    ('Bob', 'Jones', 35, true, 3, 'bobjones', '123456'),
    ('Alice', 'Miller', 40, false, 4, 'alicemiller', '123456');
    

-- Insert data into the Tourist_hotel table
INSERT INTO Tourist_hotel (tourist_id, start_date, hotel_id, room_id, end_date) VALUES
    (1, '2024-01-15', 1, 101, '2024-01-20'),
    (2, '2024-02-10', 2, 201, '2024-02-15'),
    (3, '2024-03-20', 3, 301, '2024-03-25'),
    (4, '2024-04-05', 4, 401, '2024-04-10');

-- Insert data into the Agency table
INSERT INTO Agency (name, address, contact_number) VALUES
    ('Adventure Tours', '100 Explore St, City', '555-1234'),
    ('Relaxing Getaways', '200 Tranquil Rd, Serenity', '555-5678'),
    ('City Tours Inc.', '300 Urban Ave, Metropolis', '555-9876');

-- Insert data into the Tour table
INSERT INTO Tour (name, travel_agency_id, start_date, end_date, satisfaction_level, price) VALUES
    ('Mountain Retreat', 1, '2023-01-10', '2023-01-18', 4, 1500.00),
    ('Beach Vacation', 2, '2023-02-05', '2023-02-12', 5, 1200.00),
    ('City Explorer', 3, '2024-03-15', '2024-03-22', 3, 800.00),
    ('Castle Tour', 1, '2024-04-10', '2024-04-18', 4, 1000.00);

-- Insert data into the Tourist_tour table
INSERT INTO Tourist_tour (tourist_id, tour_id) VALUES
    (1, 1),
    (2, 2),
    (3, 3),
    (4, 1);

-- Insert data into the Review_src_type table
INSERT INTO Review_src_type (type) VALUES
    ('Website'),
    ('Travel Magazine'),
    ('Social Media');

-- Insert data into the Review_src table
INSERT INTO Review_src (name, type, address) VALUES
    ('TravelReviews.com', 'Website', 'www.travelreviews.com'),
    ('Explore Magazine', 'Travel Magazine', 'www.exploremagazine.com'),
    ('Adventuregram', 'Social Media', 'www.adventuregram.com');

-- Insert data into the Tour_review table
INSERT INTO Tour_review (tourist_id, tour_id, review_text, rating, datetime, review_src_id) VALUES
    (1, 1, 'Had an amazing time in the mountains!', 5, '2023-01-20', 1),
    (2, 2, 'Beautiful beaches and great service!', 5, '2023-02-12', 2),
    (3, 3, 'City tour was okay, could be better.', 3, '2023-03-22', 3),
    (4, 1, 'Enjoyed the mountain retreat, but room could be improved.', 4, '2023-01-18', 1);

-- Insert data into the Guide table
INSERT INTO Guide (first_name, last_name, age, gender, agency_id) VALUES
    ('Michael', 'Jones', 35, true, 1),
    ('Sophie', 'Miller', 28, false, 2),
    ('David', 'Clark', 40, true, 3);

-- Insert data into the Guide_lang table
INSERT INTO Guide_lang (guide_id, language) VALUES
    (1, 'English'),
    (2, 'French'),
    (3, 'German'),
    (3, 'English');

-- Insert data into the Guide_tour table
INSERT INTO Guide_tour (guide_id, tour_id) VALUES
    (1, 1),
    (2, 2),
    (3, 3);

-- Insert data into the Transport table
INSERT INTO Transport (type) VALUES
    ('Plane'),
    ('Train'),
    ('Bus');


-- Insert data into the Trip table
INSERT INTO Trip (src_country_id, dst_country_id, transport_type, datetime_start, datetime_end, price, tour_id) VALUES
    (1, 3, 'Plane', '2024-01-10', '2024-01-18', 500.00, 1),
    (2, 4, 'Train', '2024-02-05', '2024-02-12', 300.00, 2),
    (3, 1, 'Bus', '2024-03-15', '2024-03-22', 100.00, 3),
    (1, 2, 'Plane', '2024-03-30 18:45:00', '2024-03-30 22:17:00', 500.00, 3);

-- Insert data into the Visa table
INSERT INTO Visa (name, destination_country_id, tourist_id, issue_date, expiry_date) VALUES
    ('USA Visa', 1, 1, '2024-01-01', '2024-12-31'),
    ('France Visa', 2, 2, '2024-01-15', '2024-12-31'),
    ('Spain Visa', 3, 3, '2024-02-01', '2024-12-31'),
    ('Germany Visa', 4, 4, '2024-02-15', '2024-12-31');

-- Insert data into the Attraction table
INSERT INTO Attraction (name, address, description, ticket_price, country_id) VALUES
    ('Grand Canyon', '100 Grand Canyon Rd, Arizona', 'A steep-sided canyon carved by the Colorado River in Arizona, United States.', 50.00, 1),
    ('Eiffel Tower', '200 Champ de Mars, Paris', 'A wrought-iron lattice tower on the Champ de Mars in Paris, France.', 20.00, 2),
    ('Sagrada Familia', '300 Carrer de Mallorca, Barcelona', 'A large unfinished Roman Catholic minor basilica in Barcelona, Spain.', 30.00, 3),
    ('Neuschwanstein Castle', '400 Neuschwansteinstraße, Hohenschwangau', 'A 19th-century Romanesque Revival palace on a rugged hill above the village of Hohenschwangau near Füssen in southwest Bavaria, Germany.', 40.00, 4);

-- Insert data into the Tourist_attraction table
INSERT INTO Tourist_attraction (tourist_id, attraction_id) VALUES
    ( 1, 1),
    ( 2, 2),
    ( 3, 3),
    ( 4, 4);

