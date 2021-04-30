-- Created by Vertabelo (http://vertabelo.com)
-- Last modification date: 2021-04-19 20:30:55.316

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
    Media longblob NOT NULL,
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

-- Table: HCN
CREATE TABLE HCN (
    ID int NOT NULL AUTO_INCREMENT,
    TeacherID int NOT NULL,
    MongoID varchar(50) NOT NULL,
    CONSTRAINT HCN_pk PRIMARY KEY (ID)
);

-- Table: Sessions
CREATE TABLE Sessions (
    ID int NOT NULL AUTO_INCREMENT,
    TeacherID int NOT NULL,
    Token varchar(256) NOT NULL,
    ExpirationDate datetime NOT NULL,
    CONSTRAINT Sessions_pk PRIMARY KEY (ID)
);

-- Table: Solved_HCN
CREATE TABLE Solved_HCN (
    ID int NOT NULL AUTO_INCREMENT,
    ActivityID int NOT NULL,
    OriginalHCN int NOT NULL,
    MongoID varchar(50) NOT NULL,
    Solver int NOT NULL,
    Reviewed bool NOT NULL,
    CONSTRAINT Solved_HCN_pk PRIMARY KEY (ID)
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
    Password varchar(100) NOT NULL,
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

-- Reference: HCN_Teachers (table: HCN)
ALTER TABLE HCN ADD CONSTRAINT HCN_Teachers FOREIGN KEY HCN_Teachers (TeacherID)
    REFERENCES Teachers (ID);

-- Reference: Session_Teachers (table: Sessions)
ALTER TABLE Sessions ADD CONSTRAINT Session_Teachers FOREIGN KEY Session_Teachers (TeacherID)
    REFERENCES Teachers (ID);

-- Reference: Solved_HCN_Activities (table: Solved_HCN)
ALTER TABLE Solved_HCN ADD CONSTRAINT Solved_HCN_Activities FOREIGN KEY Solved_HCN_Activities (ActivityID)
    REFERENCES Activities (ID);

-- Reference: Solved_HCN_HCN (table: Solved_HCN)
ALTER TABLE Solved_HCN ADD CONSTRAINT Solved_HCN_HCN FOREIGN KEY Solved_HCN_HCN (OriginalHCN)
    REFERENCES HCN (ID);

-- Reference: Techers_Courses (table: Courses)
ALTER TABLE Courses ADD CONSTRAINT Techers_Courses FOREIGN KEY Techers_Courses (TeacherID)
    REFERENCES Teachers (ID);

-- Constraints
ALTER TABLE Students_Courses ADD CONSTRAINT uq_Students_Courses UNIQUE(CourseID, StudentID);
ALTER TABLE CCases_HCN ADD CONSTRAINT uq_CCases_HCN UNIQUE(ClinicalCaseID, HCNID);
ALTER TABLE Courses_HCN ADD CONSTRAINT uq_Courses_HCN UNIQUE(CourseID, HCNID);
ALTER TABLE Courses_CCases ADD CONSTRAINT uq_Courses_CCases UNIQUE(CourseID, ClinicalCaseID);

-- End of file.

