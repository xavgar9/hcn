-- Created by Vertabelo (http://vertabelo.com)
-- Last modification date: 2021-01-12 01:03:12.778

-- tables
-- Table: Activities
CREATE TABLE Activities (
    ID int NOT NULL AUTO_INCREMENT,
    Title varchar(50) NOT NULL,
    Description varchar(3000) NOT NULL,
    Type varchar(50) NOT NULL,
    CreationDate datetime NOT NULL,
    LimitDate datetime NOT NULL,
    CourseID int NOT NULL,
    ClinicalCaseID int NOT NULL,
    HCNID int NOT NULL,
    Difficulty int NOT NULL,
    CONSTRAINT Activities_pk PRIMARY KEY (ID)
);

-- Table: Announcements
CREATE TABLE Announcements (
    ID int NOT NULL AUTO_INCREMENT,
    CourseID int NOT NULL,
    Title varchar(100) NOT NULL,
    Description varchar(2000) NOT NULL,
    CreationDate datetime NOT NULL,
    CONSTRAINT Announcements_pk PRIMARY KEY (ID)
);

-- Table: CCases_HCN
CREATE TABLE CCases_HCN (
    ID int NOT NULL AUTO_INCREMENT,
    ClinicalCaseID int NOT NULL,
    HCNID int NOT NULL,
    CONSTRAINT CCases_HCN_pk PRIMARY KEY (ID)
);

-- Table: Clinical_Cases
CREATE TABLE Clinical_Cases (
    ID int NOT NULL AUTO_INCREMENT,
    Title varchar(100) NOT NULL,
    Description varchar(3000) NOT NULL,
    Media varchar(2000) NOT NULL,
    TeacherID int NOT NULL,
    CONSTRAINT Clinical_Cases_pk PRIMARY KEY (ID)
);

-- Table: Courses
CREATE TABLE Courses (
    ID int NOT NULL AUTO_INCREMENT,
    TeacherID int NOT NULL,
    Name varchar(50) NOT NULL,
    CreationDate datetime NOT NULL,
    CONSTRAINT Id PRIMARY KEY (ID)
);

-- Table: Courses_CCases
CREATE TABLE Courses_CCases (
    Id int NOT NULL AUTO_INCREMENT,
    ClinicalCaseID int NOT NULL,
    CourseID int NOT NULL,
    Displayable bool NOT NULL,
    CONSTRAINT Courses_CCases_pk PRIMARY KEY (Id)
);

-- Table: Courses_HCN
CREATE TABLE Courses_HCN (
    ID int NOT NULL AUTO_INCREMENT,
    CourseID int NOT NULL,
    HCNID int NOT NULL,
    Displayable bool NOT NULL,
    CONSTRAINT Courses_HCN_pk PRIMARY KEY (ID)
);

-- Table: Feedbacks
CREATE TABLE Feedbacks (
    ID int NOT NULL AUTO_INCREMENT,
    ActivityID int NOT NULL,
    StudentID int NOT NULL,
    CONSTRAINT Feedbacks_pk PRIMARY KEY (ID)
);

-- Table: HCN
CREATE TABLE HCN (
    ID int NOT NULL AUTO_INCREMENT,
    TeacherID int NOT NULL,
    CONSTRAINT HCN_pk PRIMARY KEY (ID)
);

-- Table: Students
CREATE TABLE Students (
    ID int NOT NULL,
    Name varchar(100) NOT NULL,
    Email varchar(100) NOT NULL,
    CONSTRAINT Students_pk PRIMARY KEY (ID)
);

-- Table: Students_Courses
CREATE TABLE Students_Courses (
    ID int NOT NULL AUTO_INCREMENT,
    CourseID int NOT NULL,
    StudentID int NOT NULL,
    CONSTRAINT Students_Courses_pk PRIMARY KEY (ID)
);

-- Table: Teachers
CREATE TABLE Teachers (
    ID int NOT NULL,
    Name varchar(100) NOT NULL,
    Email varchar(100) NOT NULL,
    CONSTRAINT Teachers_pk PRIMARY KEY (ID)
);

-- foreign keys
-- Reference: Activities_Clinical_Cases (table: Activities)
ALTER TABLE Activities ADD CONSTRAINT Activities_Clinical_Cases FOREIGN KEY Activities_Clinical_Cases (ClinicalCaseID)
    REFERENCES Clinical_Cases (ID);

-- Reference: Activities_Courses (table: Activities)
ALTER TABLE Activities ADD CONSTRAINT Activities_Courses FOREIGN KEY Activities_Courses (CourseID)
    REFERENCES Courses (ID);

-- Reference: Activities_HCN (table: Activities)
ALTER TABLE Activities ADD CONSTRAINT Activities_HCN FOREIGN KEY Activities_HCN (HCNID)
    REFERENCES HCN (ID);

-- Reference: Announcements_Courses (table: Announcements)
ALTER TABLE Announcements ADD CONSTRAINT Announcements_Courses FOREIGN KEY Announcements_Courses (CourseID)
    REFERENCES Courses (ID);

-- Reference: CCases_HCN_Clinical_Cases (table: CCases_HCN)
ALTER TABLE CCases_HCN ADD CONSTRAINT CCases_HCN_Clinical_Cases FOREIGN KEY CCases_HCN_Clinical_Cases (ClinicalCaseID)
    REFERENCES Clinical_Cases (ID);

-- Reference: CCases_HCN_HCN (table: CCases_HCN)
ALTER TABLE CCases_HCN ADD CONSTRAINT CCases_HCN_HCN FOREIGN KEY CCases_HCN_HCN (HCNID)
    REFERENCES HCN (ID);

-- Reference: Clinical_Cases_Teachers (table: Clinical_Cases)
ALTER TABLE Clinical_Cases ADD CONSTRAINT Clinical_Cases_Teachers FOREIGN KEY Clinical_Cases_Teachers (TeacherID)
    REFERENCES Teachers (ID);

-- Reference: Courses_CCases_Clinical_Cases (table: Courses_CCases)
ALTER TABLE Courses_CCases ADD CONSTRAINT Courses_CCases_Clinical_Cases FOREIGN KEY Courses_CCases_Clinical_Cases (ClinicalCaseID)
    REFERENCES Clinical_Cases (ID);

-- Reference: Courses_CCases_Courses (table: Courses_CCases)
ALTER TABLE Courses_CCases ADD CONSTRAINT Courses_CCases_Courses FOREIGN KEY Courses_CCases_Courses (CourseID)
    REFERENCES Courses (ID);

-- Reference: Courses_HCN_Courses (table: Courses_HCN)
ALTER TABLE Courses_HCN ADD CONSTRAINT Courses_HCN_Courses FOREIGN KEY Courses_HCN_Courses (CourseID)
    REFERENCES Courses (ID);

-- Reference: Courses_HCN_HCN (table: Courses_HCN)
ALTER TABLE Courses_HCN ADD CONSTRAINT Courses_HCN_HCN FOREIGN KEY Courses_HCN_HCN (HCNID)
    REFERENCES HCN (ID);

-- Reference: Courses_Students_Courses (table: Students_Courses)
ALTER TABLE Students_Courses ADD CONSTRAINT Courses_Students_Courses FOREIGN KEY Courses_Students_Courses (CourseID)
    REFERENCES Courses (ID);

-- Reference: Courses_Students_Students (table: Students_Courses)
ALTER TABLE Students_Courses ADD CONSTRAINT Courses_Students_Students FOREIGN KEY Courses_Students_Students (StudentID)
    REFERENCES Students (ID);

-- Reference: Feedbacks_Activities (table: Feedbacks)
ALTER TABLE Feedbacks ADD CONSTRAINT Feedbacks_Activities FOREIGN KEY Feedbacks_Activities (ActivityID)
    REFERENCES Activities (ID);

-- Reference: Feedbacks_Students (table: Feedbacks)
ALTER TABLE Feedbacks ADD CONSTRAINT Feedbacks_Students FOREIGN KEY Feedbacks_Students (StudentID)
    REFERENCES Students (ID);

-- Reference: HCN_Teachers (table: HCN)
ALTER TABLE HCN ADD CONSTRAINT HCN_Teachers FOREIGN KEY HCN_Teachers (TeacherID)
    REFERENCES Teachers (ID);

-- Reference: Techers_Courses (table: Courses)
ALTER TABLE Courses ADD CONSTRAINT Techers_Courses FOREIGN KEY Techers_Courses (TeacherID)
    REFERENCES Teachers (ID);

-- Constraints
ALTER TABLE Students_Courses ADD CONSTRAINT uq_Students_Courses UNIQUE(CourseID, StudentID);
ALTER TABLE CCases_HCN ADD CONSTRAINT uq_CCases_HCN UNIQUE(ClinicalCaseID, HCNID);
ALTER TABLE Courses_HCN ADD CONSTRAINT uq_Courses_HCN UNIQUE(CourseID, HCNID);
ALTER TABLE Courses_CCases ADD CONSTRAINT uq_Courses_CCases UNIQUE(CourseID, ClinicalCaseID);
ALTER TABLE Feedbacks ADD CONSTRAINT uq_Feedbacks UNIQUE(ActivityID, StudentID);


-- Adding some data
-- Students
INSERT INTO Students(ID,Name,Email) VALUES (1,'Daniel Gómez Sermeño','goma@email.com');
INSERT INTO Students(ID,Name,Email) VALUES (2,'Xavier Garzón López','xavgar9@email.com');
INSERT INTO Students(ID,Name,Email) VALUES (3,'Juan F. Gil','transfer10@email.com');
INSERT INTO Students(ID,Name,Email) VALUES (4,'Edgar Silva','ednosil@email.com');
INSERT INTO Students(ID,Name,Email) VALUES (5,'Juanita María Parra Villamíl','juanitamariap@email.com');
INSERT INTO Students(ID,Name,Email) VALUES (6,'Sebastián Rodríguez Osorio Silva','sebasrosorio98@email.com');
INSERT INTO Students(ID,Name,Email) VALUES (7,'Andrés Felipe Garcés','andylukast@email.com');

-- Teachers
INSERT INTO Teachers(ID,Name,Email) VALUES (1,'Benjamín Calderón Silva','matlab@email.com');
INSERT INTO Teachers(ID,Name,Email) VALUES (2,'Oscar David Hurtado Zapata','oscrdh@email.com');
INSERT INTO Teachers(ID,Name,Email) VALUES (3,'Christian Camilo Ortiz','camilorto@email.com');

-- Courses
INSERT INTO Courses(TeacherID,Name, CreationDate) VALUES (1,'Introducción a Matlab', '2021-01-01 12:00:00');
INSERT INTO Courses(TeacherID,Name, CreationDate) VALUES (1,'Matlab avanzado', '2021-01-01 12:20:08');
INSERT INTO Courses(TeacherID,Name, CreationDate) VALUES (2,'Clases de piano', '2021-01-06 15:21:50');
INSERT INTO Courses(TeacherID,Name, CreationDate) VALUES (3,'Manejando en Cali', '2021-01-05 11:40:12');

-- Announcements
INSERT INTO Announcements(ID,CourseID,Title,Description,CreationDate) VALUES
    (1,1,'¡Bienvenidos al curso!','Este es un curso básico de Matlab. LOS AMO.', NOW());
INSERT INTO Announcements(ID,CourseID,Title,Description,CreationDate) VALUES
    (2,1,'¡Primera tarea!','Resuelvan una matriz dispersa 100x100.', NOW());
INSERT INTO Announcements(ID,CourseID,Title,Description,CreationDate)
    VALUES (3,1,'Hola a todos','Hola muchachos, los quiero mucho. Estudien bye!', NOW());
INSERT INTO Announcements(ID,CourseID,Title,Description,CreationDate)
    VALUES (4,1,'Material guía','Busquen en Youtube. "Accidentes de tránsito graves sin censura."', NOW());

-- HCN
INSERT INTO HCN(ID,TeacherID) VALUES (1,1);
INSERT INTO HCN(ID,TeacherID) VALUES (2,1);
INSERT INTO HCN(ID,TeacherID) VALUES (3,1);
INSERT INTO HCN(ID,TeacherID) VALUES (4,2);
INSERT INTO HCN(ID,TeacherID) VALUES (5,3);

-- Clinical_Cases
INSERT INTO Clinical_Cases(ID,Title,Description,Media,TeacherID) VALUES
    (1,"El joven parchado","Benjamón era un joven con IMC PARCHADO.","../activitiesresources/img1.png",1);
INSERT INTO Clinical_Cases(ID,Title,Description,Media,TeacherID) VALUES
    (2,"El pianista de la selva","Re La Mi Do#","../activitiesresources/img2.png",2);
INSERT INTO Clinical_Cases(ID,Title,Description,Media,TeacherID) VALUES
    (3,"Muerte accidental","¿Por qué se fue? ¿Y por qué murió? ¿Por qué el Señor me la quitó? Se ha ido al cielo y para poder ir yo...","../activitiesresources/ElUltimoBeso.mp3",3);

-- Activities
INSERT INTO Activities(ID,Title,Description,Type,CreationDate,LimitDate,CourseID,ClinicalCaseID,HCNID,Difficulty) VALUES
    (1,'Primera tarea, matrices dispersas','Re easy pri, solo busquen en Google.','Calificable','2021-01-08 12:00:00','2021-01-08 20:00:00',1,1,1,3);
INSERT INTO Activities(ID,Title,Description,Type,CreationDate,LimitDate,CourseID,ClinicalCaseID,HCNID,Difficulty) VALUES
    (2,'Actividad de prueba','Por favor ignoren esta actividad, gracias.','Prueba','2021-01-09 11:43:21','2021-01-19 10:59:59',2,2,2,1);

-- Feedbacks
INSERT INTO Feedbacks(ID,ActivityID,StudentID) VALUES (1,1,1);
INSERT INTO Feedbacks(ID,ActivityID,StudentID) VALUES (2,1,2);
INSERT INTO Feedbacks(ID,ActivityID,StudentID) VALUES (3,1,3);
INSERT INTO Feedbacks(ID,ActivityID,StudentID) VALUES (4,1,4);
INSERT INTO Feedbacks(ID,ActivityID,StudentID) VALUES (5,1,5);
INSERT INTO Feedbacks(ID,ActivityID,StudentID) VALUES (6,1,6);
INSERT INTO Feedbacks(ID,ActivityID,StudentID) VALUES (7,1,7);

-- Courses_HCN
INSERT INTO Courses_HCN(ID,CourseID,HCNID,Displayable) VALUES (1,1,1,1);
INSERT INTO Courses_HCN(ID,CourseID,HCNID,Displayable) VALUES (2,1,2,0);
INSERT INTO Courses_HCN(ID,CourseID,HCNID,Displayable) VALUES (3,2,3,1);

-- CCases_HCN
INSERT INTO CCases_HCN(ID,ClinicalCaseID,HCNID) VALUES (1,1,1);

-- Courses_CCases
INSERT INTO Courses_CCases(ID,ClinicalCaseID,CourseID,Displayable) VALUES (1,1,1,1);
INSERT INTO Courses_CCases(ID,ClinicalCaseID,CourseID,Displayable) VALUES (2,2,2,1);
INSERT INTO Courses_CCases(ID,ClinicalCaseID,CourseID,Displayable) VALUES (3,3,3,0);

-- Students_Courses
INSERT INTO Students_Courses(ID,CourseID,StudentID) VALUES (1,1,1);
INSERT INTO Students_Courses(ID,CourseID,StudentID) VALUES (2,1,2);
INSERT INTO Students_Courses(ID,CourseID,StudentID) VALUES (3,1,3);
INSERT INTO Students_Courses(ID,CourseID,StudentID) VALUES (4,1,4);
INSERT INTO Students_Courses(ID,CourseID,StudentID) VALUES (5,2,5);
INSERT INTO Students_Courses(ID,CourseID,StudentID) VALUES (6,2,6);
INSERT INTO Students_Courses(ID,CourseID,StudentID) VALUES (7,2,7);
INSERT INTO Students_Courses(ID,CourseID,StudentID) VALUES (8,3,1);
INSERT INTO Students_Courses(ID,CourseID,StudentID) VALUES (9,3,2);
INSERT INTO Students_Courses(ID,CourseID,StudentID) VALUES (10,3,7);