DELETE FROM Feedbacks;
ALTER TABLE Feedbacks AUTO_INCREMENT = 1;
DELETE FROM CCases_HCN;
DELETE FROM Courses_CCases;
DELETE FROM Courses_HCN;
DELETE FROM Announcements;
ALTER TABLE Announcements AUTO_INCREMENT = 1;
DELETE FROM Students_Courses;
DELETE FROM Students;
DELETE FROM Activities;
DELETE FROM HCN;
ALTER TABLE HCN AUTO_INCREMENT = 1;
DELETE FROM Clinical_Cases;
ALTER TABLE Clinical_Cases AUTO_INCREMENT = 1;
DELETE FROM Courses;
ALTER TABLE Courses AUTO_INCREMENT = 1;
DELETE FROM Teachers;

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
    (1,1,'¡Bienvenidos al curso!','Este es un curso básico de Matlab. LOS AMO.', "2021-01-17 13:34:28");
INSERT INTO Announcements(ID,CourseID,Title,Description,CreationDate) VALUES
    (2,1,'¡Primera tarea!','Resuelvan una matriz dispersa 100x100.', "2021-01-17 13:34:28");
INSERT INTO Announcements(ID,CourseID,Title,Description,CreationDate) VALUES
    (3,1,'Hola a todos','Hola muchachos, los quiero mucho. Estudien bye!', "2021-01-17 13:34:28");
INSERT INTO Announcements(ID,CourseID,Title,Description,CreationDate) VALUES
    (4,1,'Material guía','Busquen en Youtube. "Accidentes de tránsito graves sin censura."', "2021-01-17 13:34:28");

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
INSERT INTO CCases_HCN(ID,ClinicalCaseID,HCNID) VALUES (2,2,2);
INSERT INTO CCases_HCN(ID,ClinicalCaseID,HCNID) VALUES (3,3,3);

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