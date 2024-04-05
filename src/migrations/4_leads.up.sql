CREATE TABLE IF NOT EXISTS `leads` (
    `id` int NOT NULL AUTO_INCREMENT,
    `first_name` varchar(180) NOT NULL,
    `last_name` varchar(180) NOT NULL,
    `email` varchar(180) NOT NULL,
    `phone` varchar(15) NULL,
    `location_id` int NOT NULL,
    `rental_id` int NOT NULL,
    `message` text NULL,
    PRIMARY KEY (`id`),
    FOREIGN KEY (`location_id`) REFERENCES locations(`id`),
    FOREIGN KEY (`rental_id`) REFERENCES rentals(`id`)
);
