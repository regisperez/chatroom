CREATE DATABASE IF NOT EXISTS chatroom_test;
CREATE DATABASE IF NOT EXISTS chatroom;

CREATE TABLE IF NOT EXISTS chatroom.users
(
    id SERIAL,
    name TEXT NOT NULL,
    login TEXT NOT NULL,
    password TEXT NOT NULL,
    CONSTRAINT users_pkey PRIMARY KEY (id)
    );

INSERT INTO users(name, login,password) VALUES("Jobsity", "admin","$2a$10$Oo4QBJ5ggrvuX0Cb0tWd8uLHMmT6NbDUa5jFTlM9qfbsD.VvyVlae");
INSERT INTO users(name, login,password) VALUES("Regis Perez", "regis","$2a$10$.sdCQlnGqve0xqMpzrq7M.JH.uvHw0vPs3J3b/rhwxLBzdtbFVENa");
INSERT INTO users(name, login,password) VALUES("Hideo Kojima", "hideo","$2a$10$eFQkvzLZ93/Sb1aJrRK2FO3tAlq1hznS.CRuT8UW05SwS5hedPQ1.");
INSERT INTO users(name, login,password) VALUES("Lara Croft", "lara","$2a$10$/jqO1CTAftrWit41bCK8ROScnLUEvPWD.b6tKygaPeGhIKpwUwkQW");
INSERT INTO users(name, login,password) VALUES("Catarina Lopes", "catarina","$2a$10$IV991CsHckG9W/XNxWn86.vW6v1SozItNmXBIvVO0GQ1c5E2Wi/O2");



