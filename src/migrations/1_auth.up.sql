CREATE TABLE IF NOT EXISTS `users` (
    `id` int NOT NULL AUTO_INCREMENT,
    `name` varchar(180) NOT NULL,
    `password` varchar(180) NOT NULL,
    `token` varchar(180) DEFAULT NULL,
    `token_exp` datetime DEFAULT NULL,
    PRIMARY KEY (`id`),
    UNIQUE (`name`)
);
