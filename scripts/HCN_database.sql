-- Created by Vertabelo (http://vertabelo.com)
-- Last modification date: 2021-01-11 04:33:55.608

-- tables
-- Table: Activities
CREATE TABLE Activities (
    Id int NOT NULL AUTO_INCREMENT,
    Title varchar(50) NOT NULL,
    Description varchar(3000) NOT NULL,
    Type varchar(50) NOT NULL,
    CreationDate datetime NOT NULL,
    LimitDate datetime NOT NULL,
    CoursesId int NOT NULL,
    ClinicalCasesId int NOT NULL,
    HCNId int NOT NULL,
    Difficulty int NOT NULL,
    CONSTRAINT Activities_pk PRIMARY KEY (Id)
);

-- Table: Announcements
CREATE TABLE Announcements (
    Id int NOT NULL AUTO_INCREMENT,
    CoursesId int NOT NULL,
    Title varchar(100) NOT NULL,
    Description varchar(2000) NOT NULL,
    CreationDate datetime NOT NULL,
    CONSTRAINT Announcements_pk PRIMARY KEY (Id)
);

-- Table: Clinical_Cases
CREATE TABLE Clinical_Cases (
    Id int NOT NULL AUTO_INCREMENT,
    Title varchar(100) NOT NULL,
    Description varchar(3000) NOT NULL,
    Media varchar(2000) NOT NULL,
    TeachersId int NOT NULL,
    HCNId int NOT NULL,
    CONSTRAINT Clinical_Cases_pk PRIMARY KEY (Id)
);

-- Table: Courses
CREATE TABLE Courses (
    Id int NOT NULL AUTO_INCREMENT,
    Teacher int NOT NULL,
    Name varchar(50) NOT NULL,
    CreationDate datetime NOT NULL,
    CONSTRAINT Id PRIMARY KEY (Id)
);

-- Table: Courses_CCases
CREATE TABLE Courses_CCases (
    Id int NOT NULL AUTO_INCREMENT,
    ClinicalCasesId int NOT NULL,
    CoursesId int NOT NULL,
    Dispayable bool NOT NULL,
    CONSTRAINT Courses_CCases_pk PRIMARY KEY (Id)
);

-- Table: Courses_HCN
CREATE TABLE Courses_HCN (
    Id int NOT NULL AUTO_INCREMENT,
    CoursesId int NOT NULL,
    HcnId int NOT NULL,
    Displayable bool NOT NULL,
    CONSTRAINT Courses_HCN_pk PRIMARY KEY (Id)
);

-- Table: Courses_Students
CREATE TABLE Courses_Students (
    Id int NOT NULL AUTO_INCREMENT,
    CoursesId int NOT NULL,
    StudentsId int NOT NULL,
    CONSTRAINT Courses_Students_pk PRIMARY KEY (Id)
);

-- Table: Feedbacks
CREATE TABLE Feedbacks (
    Id int NOT NULL AUTO_INCREMENT,
    ActivitiesId int NOT NULL,
    StudentsId int NOT NULL,
    CONSTRAINT Feedbacks_pk PRIMARY KEY (Id)
);

-- Table: HCN
CREATE TABLE HCN (
    Id int NOT NULL AUTO_INCREMENT,
    TeachersId int NOT NULL,
    CONSTRAINT HCN_pk PRIMARY KEY (Id)
);

-- Table: Students
CREATE TABLE Students (
    Id int NOT NULL,
    Name varchar(100) NOT NULL,
    Email varchar(100) NOT NULL,
    CONSTRAINT Students_pk PRIMARY KEY (Id)
);

-- Table: Teachers
CREATE TABLE Teachers (
    Id int NOT NULL,
    Name varchar(100) NOT NULL,
    Email varchar(100) NOT NULL,
    CONSTRAINT Teachers_pk PRIMARY KEY (Id)
);

-- foreign keys
-- Reference: Activities_Clinical_Cases (table: Activities)
ALTER TABLE Activities ADD CONSTRAINT Activities_Clinical_Cases FOREIGN KEY Activities_Clinical_Cases (ClinicalCasesId)
    REFERENCES Clinical_Cases (Id);

-- Reference: Activities_Courses (table: Activities)
ALTER TABLE Activities ADD CONSTRAINT Activities_Courses FOREIGN KEY Activities_Courses (CoursesId)
    REFERENCES Courses (Id);

-- Reference: Activities_HCN (table: Activities)
ALTER TABLE Activities ADD CONSTRAINT Activities_HCN FOREIGN KEY Activities_HCN (HCNId)
    REFERENCES HCN (Id);

-- Reference: Announcements_Courses (table: Announcements)
ALTER TABLE Announcements ADD CONSTRAINT Announcements_Courses FOREIGN KEY Announcements_Courses (CoursesId)
    REFERENCES Courses (Id);

-- Reference: Clinical_Cases_HCN (table: Clinical_Cases)
ALTER TABLE Clinical_Cases ADD CONSTRAINT Clinical_Cases_HCN FOREIGN KEY Clinical_Cases_HCN (HCNId)
    REFERENCES HCN (Id);

-- Reference: Clinical_Cases_Teachers (table: Clinical_Cases)
ALTER TABLE Clinical_Cases ADD CONSTRAINT Clinical_Cases_Teachers FOREIGN KEY Clinical_Cases_Teachers (TeachersId)
    REFERENCES Teachers (Id);

-- Reference: Courses_CCases_Clinical_Cases (table: Courses_CCases)
ALTER TABLE Courses_CCases ADD CONSTRAINT Courses_CCases_Clinical_Cases FOREIGN KEY Courses_CCases_Clinical_Cases (ClinicalCasesId)
    REFERENCES Clinical_Cases (Id);

-- Reference: Courses_CCases_Courses (table: Courses_CCases)
ALTER TABLE Courses_CCases ADD CONSTRAINT Courses_CCases_Courses FOREIGN KEY Courses_CCases_Courses (CoursesId)
    REFERENCES Courses (Id);

-- Reference: Courses_HCN_Courses (table: Courses_HCN)
ALTER TABLE Courses_HCN ADD CONSTRAINT Courses_HCN_Courses FOREIGN KEY Courses_HCN_Courses (CoursesId)
    REFERENCES Courses (Id);

-- Reference: Courses_HCN_HCN (table: Courses_HCN)
ALTER TABLE Courses_HCN ADD CONSTRAINT Courses_HCN_HCN FOREIGN KEY Courses_HCN_HCN (HcnId)
    REFERENCES HCN (Id);

-- Reference: Courses_Students_Courses (table: Courses_Students)
ALTER TABLE Courses_Students ADD CONSTRAINT Courses_Students_Courses FOREIGN KEY Courses_Students_Courses (CoursesId)
    REFERENCES Courses (Id);

-- Reference: Courses_Students_Students (table: Courses_Students)
ALTER TABLE Courses_Students ADD CONSTRAINT Courses_Students_Students FOREIGN KEY Courses_Students_Students (StudentsId)
    REFERENCES Students (Id);

-- Reference: Feedbacks_Activities (table: Feedbacks)
ALTER TABLE Feedbacks ADD CONSTRAINT Feedbacks_Activities FOREIGN KEY Feedbacks_Activities (ActivitiesId)
    REFERENCES Activities (Id);

-- Reference: Feedbacks_Students (table: Feedbacks)
ALTER TABLE Feedbacks ADD CONSTRAINT Feedbacks_Students FOREIGN KEY Feedbacks_Students (StudentsId)
    REFERENCES Students (Id);

-- Reference: HCN_Teachers (table: HCN)
ALTER TABLE HCN ADD CONSTRAINT HCN_Teachers FOREIGN KEY HCN_Teachers (TeachersId)
    REFERENCES Teachers (Id);

-- Reference: Techers_Courses (table: Courses)
ALTER TABLE Courses ADD CONSTRAINT Techers_Courses FOREIGN KEY Techers_Courses (Teacher)
    REFERENCES Teachers (Id);


-- Adding some data
-- Students
INSERT INTO Students(Id,Name,Email) VALUES (1,'Daniel Gómez Sermeño','goma@email.com');
INSERT INTO Students(Id,Name,Email) VALUES (2,'Xavier Garzón López','xavgar9@email.com');
INSERT INTO Students(Id,Name,Email) VALUES (3,'Juan F. Gil','transfer10@email.com');
INSERT INTO Students(Id,Name,Email) VALUES (4,'Edgar Silva','ednosil@email.com');
INSERT INTO Students(Id,Name,Email) VALUES (5,'Juanita María Parra Villamíl','juanitamariap@email.com');
INSERT INTO Students(Id,Name,Email) VALUES (6,'Sebastián Rodríguez Osorio Silva','sebasrosorio98@email.com');
INSERT INTO Students(Id,Name,Email) VALUES (7,'Andrés Felipe Garcés','andylukast@email.com');
-- Teachers
INSERT INTO Teachers(Id,Name,Email) VALUES (1,'Benjamín Calderón Silva','matlab@email.com');
INSERT INTO Teachers(Id,Name,Email) VALUES (2,'Oscar David Hurtado Zapata','oscrdh@email.com');
INSERT INTO Teachers(Id,Name,Email) VALUES (3,'Christian Camilo Ortiz','camilorto@email.com');
-- Courses
INSERT INTO Courses(Id,Teacher,Name, CreationDate) VALUES (1,1,'Introducción a Matlab', '2021-01-01 12:00:00');
INSERT INTO Courses(Id,Teacher,Name, CreationDate) VALUES (2,1,'Matlab avanzado', '2021-01-01 12:20:08');
INSERT INTO Courses(Id,Teacher,Name, CreationDate) VALUES (3,2,'Clases de piano', '2021-01-06 15:21:50');
INSERT INTO Courses(Id,Teacher,Name, CreationDate) VALUES (4,3,'Manejando en Cali', '2021-01-05 11:40:12');

-- Announcements
INSERT INTO Announcements(Id,CoursesId,Title,Description,CreationDate) VALUES
    (1,1,'¡Bienvenidos al curso!','Este es un curso básico de Matlab. LOS AMO.', NOW());
INSERT INTO Announcements(Id,CoursesId,Title,Description,CreationDate) VALUES
    (2,1,'¡Primera tarea!','Resuelvan una matriz dispersa 100x100.', NOW());
INSERT INTO Announcements(Id,CoursesId,Title,Description,CreationDate)
    VALUES (3,1,'Hola a todos','Hola muchachos, los quiero mucho. Estudien bye!', NOW());
INSERT INTO Announcements(Id,CoursesId,Title,Description,CreationDate)
    VALUES (4,1,'Material guía','Busquen en Youtube. "Accidentes de tránsito graves sin censura."', NOW());

-- HCN
INSERT INTO HCN(Id,TeachersId) VALUES (1,1);
INSERT INTO HCN(Id,TeachersId) VALUES (2,1);
INSERT INTO HCN(Id,TeachersId) VALUES (3,1);
INSERT INTO HCN(Id,TeachersId) VALUES (4,2);
INSERT INTO HCN(Id,TeachersId) VALUES (5,3);

-- Clinical_Cases
INSERT INTO Clinical_Cases(Id,Title,Description,Media,TeachersId,HCNId) VALUES
    (1,"El joven parchado","Benjamón era un joven con IMC PARCHADO.","../activitiesresources/img1.png",1,1);
INSERT INTO Clinical_Cases(Id,Title,Description,Media,TeachersId,HCNId) VALUES
    (2,"El pianista de la selva","Re La Mi Do#","../activitiesresources/img2.png",2,2);
INSERT INTO Clinical_Cases(Id,Title,Description,Media,TeachersId,HCNId) VALUES
    (3,"Muerte accidental","¿Por qué se fue? ¿Y por qué murió? ¿Por qué el Señor me la quitó? Se ha ido al cielo y para poder ir yo...","../activitiesresources/ElUltimoBeso.mp3",3,3);

-- Activities
INSERT INTO Activities(Id,Title,Description,Type,CreationDate,LimitDate,CoursesId,ClinicalCasesId,HCNId,Difficulty) VALUES
    (1,'Primera tarea, matrices dispersas','Re easy pri, solo busquen en Google.','Calificable','2021-01-08 12:00:00','2021-01-08 20:00:00',1,1,1,3);
INSERT INTO Activities(Id,Title,Description,Type,CreationDate,LimitDate,CoursesId,ClinicalCasesId,HCNId,Difficulty) VALUES
    (2,'Actividad de prueba','Por favor ignoren esta actividad, gracias.','Prueba','2021-01-09 11:43:21','2021-01-19 10:59:59',2,2,2,1);

-- Feedbacks
INSERT INTO Feedbacks(Id,ActivitiesId,StudentsId) VALUES (1,1,1);
INSERT INTO Feedbacks(Id,ActivitiesId,StudentsId) VALUES (2,1,2);
INSERT INTO Feedbacks(Id,ActivitiesId,StudentsId) VALUES (3,1,3);
INSERT INTO Feedbacks(Id,ActivitiesId,StudentsId) VALUES (4,1,4);
INSERT INTO Feedbacks(Id,ActivitiesId,StudentsId) VALUES (5,1,5);
INSERT INTO Feedbacks(Id,ActivitiesId,StudentsId) VALUES (6,1,6);
INSERT INTO Feedbacks(Id,ActivitiesId,StudentsId) VALUES (7,1,7);

-- End of file.