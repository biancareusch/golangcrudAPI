DROP TABLE IF EXISTS `person`;
CREATE TABLE `person`(
                         `id` int(6) unsigned NOT NULL AUTO_INCREMENT,
                         `first_name` varchar(25) NOT NULL,
                         `last_name` varchar(25) NOT NULL,
                         `age` int(6) unsigned NOT NULL,
                         `date_joined` TIMESTAMP NOT NULL,
                         `date_updated` TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
                         PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=latin1 ;

DROP TABLE IF EXISTS `job`;
CREATE TABLE `job`(
                      `id` int(6) UNSIGNED NOT NULL AUTO_INCREMENT,
                      `title` varchar(250) NOT NULL,
                      `description` TEXT NOT NULL,
                      `salary` int(12) unsigned NOT NULL,
                      FK_person int UNSIGNED,
                      INDEX (FK_person),
                      FOREIGN KEY (FK_person) REFERENCES person (id),
                      PRIMARY KEY (`id`)
)ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=latin1;

INSERT INTO person(first_name, last_name, age, date_joined)
VALUES('Bianca','Reusch',27,NOW()),
      ('Lisa','Smith',28,NOW()),
      ('Mark','Williams',33,NOW());

INSERT INTO job(title, description, salary, FK_person) VALUES('Software Developer', 'Software Dev experienced in Go	', 65000,1),
                                                             ('Backend Developer','Backend Dev experienced in Go	',75000,2);

INSERT INTO job(title, description, salary) VALUES('Software DeveloperNEW', 'looking for experiences software Developer', 65000),
                                                  ('Backend DeveloperNEW','looking fro experiences dev',75000);