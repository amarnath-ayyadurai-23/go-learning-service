CREATE DATABASE IF NOT EXISTS `learning`;

CREATE TABLE IF NOT EXISTS people (	  	
id integer AUTO_INCREMENT NOT NULL PRIMARY KEY,
name varchar(255) NOT NULL		
);
Insert into people (name) values("Peter");
Insert into people (name) values("ToM");
Insert into people (name) values("Amar");
