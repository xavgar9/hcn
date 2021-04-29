DELETE FROM Sessions;
DELETE FROM CCases_HCN;
DELETE FROM Courses_CCases;
ALTER TABLE Courses_CCases AUTO_INCREMENT = 1;
DELETE FROM Courses_HCN;
ALTER TABLE Courses_HCN AUTO_INCREMENT = 1;
DELETE FROM Announcements;
ALTER TABLE Announcements AUTO_INCREMENT = 1;
DELETE FROM Students_Courses;
ALTER TABLE Students_Courses AUTO_INCREMENT = 1;
DELETE FROM Students;
DELETE FROM Solved_HCN;
ALTER TABLE Solved_HCN AUTO_INCREMENT = 1;
DELETE FROM Activities;
ALTER TABLE Activities AUTO_INCREMENT = 1;
DELETE FROM HCN;
ALTER TABLE HCN AUTO_INCREMENT = 1;
DELETE FROM Clinical_Cases;
ALTER TABLE Clinical_Cases AUTO_INCREMENT = 1;
DELETE FROM Courses;
ALTER TABLE Courses AUTO_INCREMENT = 1;
DELETE FROM Teachers;

-- Students
INSERT INTO Students(ID,Name,Email) VALUES (10001,'Daniel Gómez Sermeño','goma@email.com');
INSERT INTO Students(ID,Name,Email) VALUES (10002,'Xavier Garzón López','xavgar9@email.com');
INSERT INTO Students(ID,Name,Email) VALUES (10003,'Juan F. Gil','transfer10@email.com');
INSERT INTO Students(ID,Name,Email) VALUES (10004,'Edgar Silva','ednosil@email.com');
INSERT INTO Students(ID,Name,Email) VALUES (10005,'Juanita María Parra Villamíl','juanitamariap@email.com');
INSERT INTO Students(ID,Name,Email) VALUES (10006,'Sebastián Rodríguez Osorio Silva','sebasrosorio98@email.com');
INSERT INTO Students(ID,Name,Email) VALUES (10007,'Andrés Felipe Garcés','andylukast@email.com');

-- Teachers
INSERT INTO Teachers(ID,Name,Email,Password) VALUES (50001,'Gerardo Mauricio Sarria','gerardo@email.com', '4024fb06e1423da90b80f0274e8e4476');
INSERT INTO Teachers(ID,Name,Email,Password) VALUES (50002,'Juan Carlos Martinez','juan@email.com', 'a94652aa97c7211ba8954dd15a3cf838');
INSERT INTO Teachers(ID,Name,Email,Password) VALUES (50003,'Jhoan Lozano Rojas','jhoan@email.com', '88ca9791c0f2e27a503c23b74896b377');
INSERT INTO Teachers(ID,Name,Email,Password) VALUES (50004,'Xavier Garzón López','xavier@email.com', '0f5366b3b19afc3184d23bc73d8cd311');

-- Courses
INSERT INTO Courses(ID,TeacherID,Name,CreationDate) VALUES (1,50001,'Introducción a Matlab', '2021-01-01 12:00:00');
INSERT INTO Courses(ID,TeacherID,Name,CreationDate) VALUES (2,50001,'Matlab avanzado', '2021-01-01 12:20:08');
INSERT INTO Courses(ID,TeacherID,Name,CreationDate) VALUES (3,50002,'Clases de piano', '2021-01-06 15:21:50');
INSERT INTO Courses(ID,TeacherID,Name,CreationDate) VALUES (4,50003,'Manejando en Cali', '2021-01-05 11:40:12');

-- Announcements
INSERT INTO Announcements(ID,CourseID,Title,Description,CreationDate) VALUES
    (1,1,'¡Bienvenidos al curso!','Este es un curso básico de Matlab. LOS AMO.', "2021-01-17 13:34:28");
INSERT INTO Announcements(ID,CourseID,Title,Description,CreationDate) VALUES
    (2,1,'¡Primera tarea!','Resuelvan una matriz dispersa 100x100.', "2021-01-17 13:34:28");
INSERT INTO Announcements(ID,CourseID,Title,Description,CreationDate) VALUES
    (3,1,'Hola a todos','Hola muchachos, los quiero mucho. Estudien bye!', "2021-01-17 13:34:28");
INSERT INTO Announcements(ID,CourseID,Title,Description,CreationDate) VALUES
    (4,1,'Material guía','Busquen en Youtube. "Accidentes de tránsito graves sin censura."', "2021-01-17 13:34:28");

-- HCN
INSERT INTO HCN(ID,TeacherID,MongoID) VALUES (1,50001,"607ec7dee81d0518b08d3da0");
INSERT INTO HCN(ID,TeacherID,MongoID) VALUES (2,50001,"607ec7dee81d0518b08d3db0");

-- Clinical_Cases
INSERT INTO Clinical_Cases(ID,Title,Description,Media,TeacherID) VALUES
    (1,"El joven parchado","Benjamón era un joven con IMC PARCHADO.","Li4vYWN0aXZpdGllc3Jlc291cmNlcy9pbWcxLnBuZw==",50001);
INSERT INTO Clinical_Cases(ID,Title,Description,Media,TeacherID) VALUES
    (2,"El pianista de la selva","Re La Mi Do#","Li4vYWN0aXZpdGllc3Jlc291cmNlcy9pbWcyLnBuZw==",50002);
INSERT INTO Clinical_Cases(ID,Title,Description,Media,TeacherID) VALUES
    (3,"Muerte accidental","¿Por qué se fue? ¿Y por qué murió? ¿Por qué el Señor me la quitó? Se ha ido al cielo y para poder ir yo...","Li4vYWN0aXZpdGllc3Jlc291cmNlcy9FbFVsdGltb0Jlc28ubXAz",50003);

-- Activities
INSERT INTO Activities(ID,Title,Description,Type,CreationDate,LimitDate,CourseID,ClinicalCaseID,HCNID,Difficulty) VALUES
    (1,'Primera tarea, matrices dispersas','Re easy pri, solo busquen en Google.','Calificable','2021-01-08 12:00:00','2021-01-08 20:00:00',1,1,1,3);
INSERT INTO Activities(ID,Title,Description,Type,CreationDate,LimitDate,CourseID,ClinicalCaseID,HCNID,Difficulty) VALUES
    (2,'Actividad de prueba','Por favor ignoren esta actividad, gracias.','Prueba','2021-01-09 11:43:21','2021-01-19 10:59:59',2,2,2,1);

-- Courses_HCN
INSERT INTO Courses_HCN(ID,CourseID,HCNID,Displayable) VALUES (1,1,1,1);
INSERT INTO Courses_HCN(ID,CourseID,HCNID,Displayable) VALUES (2,1,2,0);

-- CCases_HCN
INSERT INTO CCases_HCN(ID,ClinicalCaseID,HCNID) VALUES (1,1,1);

-- Courses_CCases
INSERT INTO Courses_CCases(ID,ClinicalCaseID,CourseID,Displayable) VALUES (1,1,1,1);
INSERT INTO Courses_CCases(ID,ClinicalCaseID,CourseID,Displayable) VALUES (2,2,2,1);
INSERT INTO Courses_CCases(ID,ClinicalCaseID,CourseID,Displayable) VALUES (3,3,3,0);

-- Students_Courses
INSERT INTO Students_Courses(ID,CourseID,StudentID) VALUES (1,1,10001);
INSERT INTO Students_Courses(ID,CourseID,StudentID) VALUES (2,1,10002);
INSERT INTO Students_Courses(ID,CourseID,StudentID) VALUES (3,1,10003);
INSERT INTO Students_Courses(ID,CourseID,StudentID) VALUES (4,1,10004);
INSERT INTO Students_Courses(ID,CourseID,StudentID) VALUES (5,2,10005);
INSERT INTO Students_Courses(ID,CourseID,StudentID) VALUES (6,2,10006);
INSERT INTO Students_Courses(ID,CourseID,StudentID) VALUES (7,2,10007);
INSERT INTO Students_Courses(ID,CourseID,StudentID) VALUES (8,3,10001);
INSERT INTO Students_Courses(ID,CourseID,StudentID) VALUES (9,3,10002);
INSERT INTO Students_Courses(ID,CourseID,StudentID) VALUES (10,3,10007);

-- Solved_HCN
INSERT INTO Solved_HCN(ID,ActivityID,OriginalHCN,MongoID,Solver,Reviewed) VALUES (1,1,1,"607ec7dee81d0518b08d3da1",10001,0);
INSERT INTO Solved_HCN(ID,ActivityID,OriginalHCN,MongoID,Solver,Reviewed) VALUES (2,1,1,"607ec7dee81d0518b08d3da2",10002,0);
INSERT INTO Solved_HCN(ID,ActivityID,OriginalHCN,MongoID,Solver,Reviewed) VALUES (3,1,1,"607ec7dee81d0518b08d3da3",10003,0);
INSERT INTO Solved_HCN(ID,ActivityID,OriginalHCN,MongoID,Solver,Reviewed) VALUES (4,1,1,"607ec7dee81d0518b08d3da4",10004,0);
INSERT INTO Solved_HCN(ID,ActivityID,OriginalHCN,MongoID,Solver,Reviewed) VALUES (5,1,1,"607ec7dee81d0518b08d3da5",50001,0);

INSERT INTO Solved_HCN(ID,ActivityID,OriginalHCN,MongoID,Solver,Reviewed) VALUES (6,2,2,"607ec7dee81d0518b08d3db1",10001,1);
INSERT INTO Solved_HCN(ID,ActivityID,OriginalHCN,MongoID,Solver,Reviewed) VALUES (7,2,2,"607ec7dee81d0518b08d3db2",10002,1);
INSERT INTO Solved_HCN(ID,ActivityID,OriginalHCN,MongoID,Solver,Reviewed) VALUES (8,2,2,"607ec7dee81d0518b08d3db3",10003,1);
INSERT INTO Solved_HCN(ID,ActivityID,OriginalHCN,MongoID,Solver,Reviewed) VALUES (9,2,2,"607ec7dee81d0518b08d3db4",10004,1);
INSERT INTO Solved_HCN(ID,ActivityID,OriginalHCN,MongoID,Solver,Reviewed) VALUES (10,2,2,"607ec7dee81d0518b08d3db5",50001,0);