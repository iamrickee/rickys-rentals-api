CREATE TABLE IF NOT EXISTS `rentals` (
    `id` int NOT NULL AUTO_INCREMENT,
    `name` varchar(180) NOT NULL,
    `description` text NOT NULL,
    `image` varchar(180) NOT NULL,
    PRIMARY KEY (`id`),
    UNIQUE (`name`)
);
