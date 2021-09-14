/************ TABLES **************************/
-- user who books room
CREATE TABLE IF NOT EXISTS users (
    id          SERIAL PRIMARY KEY,
    firstname   VARCHAR(20), 
    lastname    VARCHAR(20), 
    email       VARCHAR(100) NOT NULL UNIQUE 
);


-- HOTELS (hotels, guest rooms) available 
CREATE TABLE IF NOT EXISTS hotels (
    id          SERIAL PRIMARY KEY,
    country     VARCHAR(20) NOT NULL, 
    destination VARCHAR(50) NOT NULL, 
    hotel       VARCHAR(100) UNIQUE
);

-- rooms in HOTELS 
CREATE TABLE IF NOT EXISTS rooms (
    id          SERIAL PRIMARY KEY, 
    hotelID  integer REFERENCES HOTELS,
    roomtype    VARCHAR(50) NOT NULL,
    price       INTEGER CHECK  (price > 0), --price per night
    maxpersons  INTEGER CHECK (maxpersons > 0)
);

-- bookings
CREATE TABLE IF NOT EXISTS bookings (
    id          SERIAL PRIMARY KEY,
    roomID      INTEGER REFERENCES rooms,
    checkin     DATE NOT NULL,
    checkout    DATE NOT NULL
);



/************ DATA **************************/
-- users 
INSERT INTO users(firstname, lastname, email) VALUES('Denis','Shchuka','goworld@example.com');
INSERT INTO users(firstname, lastname, email) VALUES('Olaff','Swensonn','olfs@example.com');
INSERT INTO users(firstname, lastname, email) VALUES('Mark','Johnson','mjk-1923@example.com');

-- hotels 
INSERT INTO hotels(country, destination, hotel) VALUES('Russia','Sochi','Sunrise Hotel');
INSERT INTO hotels(country, destination, hotel) VALUES('United States','Hawaii','Green Villas Hotel');
INSERT INTO hotels(country, destination, hotel) VALUES('India','Jaipur','King Palace Hotel');
INSERT INTO hotels(country, destination, hotel) VALUES('Thailand','Phuket','Holiday Inn Resort');
INSERT INTO hotels(country, destination, hotel) VALUES('United Arab Emirates','Dubai','Grand Hyatt');


-- rooms
INSERT INTO rooms(hotelID,roomtype,price,maxpersons) VALUES((SELECT id from HOTELS WHERE hotel='Sunrise Hotel'), 'Standard Single', 38.5, 1);
INSERT INTO rooms(hotelID,roomtype,price,maxpersons) VALUES((SELECT id from HOTELS WHERE hotel='Sunrise Hotel'), 'Deluxe Double', 45, 2);
INSERT INTO rooms(hotelID,roomtype,price,maxpersons) VALUES((SELECT id from HOTELS WHERE hotel='Sunrise Hotel'), 'Superior', 41, 2);
INSERT INTO rooms(hotelID,roomtype,price,maxpersons) VALUES((SELECT id from HOTELS WHERE hotel='Green Villas Hotel'), 'Standard', 50, 2);
INSERT INTO rooms(hotelID,roomtype,price,maxpersons) VALUES((SELECT id from HOTELS WHERE hotel='Green Villas Hotel'), 'Deluxe', 80, 3);
INSERT INTO rooms(hotelID,roomtype,price,maxpersons) VALUES((SELECT id from HOTELS WHERE hotel='Grand Hyatt'), 'King Deluxe', 318, 2);
INSERT INTO rooms(hotelID,roomtype,price,maxpersons) VALUES((SELECT id from HOTELS WHERE hotel='Holiday Inn Resort'), 'Standard Room', 38, 2);
INSERT INTO rooms(hotelID,roomtype,price,maxpersons) VALUES((SELECT id from HOTELS WHERE hotel='Holiday Inn Resort'), 'Double Room', 44, 2);