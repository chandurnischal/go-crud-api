CREATE TABLE users (
    id INT AUTO_INCREMENT PRIMARY KEY,
    name TEXT NOT NULL,
    location TEXT NOT NULL
);

INSERT INTO users (name, location) VALUES 
('Frodo Baggins', 'Shire'),
('Samwise Gamgee', 'Shire'),
('Gandalf the Grey', 'Middle-earth'),
('Aragorn', 'Gondor'),
('Legolas', 'Mirkwood'),
('Gimli', 'Erebor'),
('Boromir', 'Gondor'),
('Galadriel', 'Lothlórien'),
('Saruman', 'Isengard'),
('Gollum', 'Misty Mountains'),
('Elrond', 'Rivendell'),
('Arwen', 'Rivendell'),
('Éowyn', 'Rohan'),
('Faramir', 'Gondor'),
('Denethor', 'Gondor');