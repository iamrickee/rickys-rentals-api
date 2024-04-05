CREATE TABLE IF NOT EXISTS `locations` (
    `id` int NOT NULL AUTO_INCREMENT,
    `address` varchar(180) NOT NULL,
    `city` varchar(180) NOT NULL,
    `state` varchar(2) NOT NULL,
    `zip` varchar(5) NOT NULL,
    PRIMARY KEY (`id`)
);
