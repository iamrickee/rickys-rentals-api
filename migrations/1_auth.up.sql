CREATE TABLE IF NOT EXISTS `users` (
    `id` int(11) NOT NULL auto_increment,
    `name` varchar(180) NOT NULL,
    `password` varchar(180) NOT NULL,
    `token` varchar(180) NULL,
    `token_exp` DATETIME NULL,
    PRIMARY KEY (`id`)
);