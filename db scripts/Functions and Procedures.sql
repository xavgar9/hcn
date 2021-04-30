SET GLOBAL log_bin_trust_function_creators = 1;

###########################################################################################
                                    #CreateActivity
###########################################################################################
DROP FUNCTION IF EXISTS CreateActivity;
DELIMITER //
CREATE FUNCTION CreateActivity(title VARCHAR(50), description VARCHAR(3000), type VARCHAR(50), creationDate DATETIME, limitDate DATETIME, courseID INT, clinicalCaseID INT, hcnID INT, difficulty INT)
  RETURNS VARCHAR(20)
 
  BEGIN
    DECLARE s VARCHAR(20);    
    IF (EXISTS(SELECT 1 FROM HCN WHERE HCN.ID=hcnID)) THEN
            INSERT INTO Activities (Title, Description, Type, CreationDate, LimitDate, CourseID, ClinicalCaseID, HCNID, Difficulty)
        		VALUES (title, description, type, creationDate, limitDate, courseID, clinicalCaseID, hcnID, difficulty);
            
            SET s = "Update";
	ELSE

            SET s = "HCN doesnt exist";
    END IF;
    RETURN s;
  END //
DELIMITER ;

###########################################################################################
                                    #SaveToken
###########################################################################################
DROP FUNCTION IF EXISTS SaveToken;
DELIMITER //
CREATE FUNCTION SaveToken(email VARCHAR(100), token VARCHAR(500), expirationDate VARCHAR(30))
  RETURNS VARCHAR(20)

  BEGIN
    DECLARE teacherID INT;    
    DECLARE s VARCHAR(20);

    IF (EXISTS(SELECT 1 FROM Teachers WHERE Teachers.Email=email)) THEN
        SELECT Teachers.ID INTO teacherID FROM Teachers WHERE Teachers.Email=email;
        IF (EXISTS(SELECT 1 FROM Sessions WHERE Sessions.TeacherID=teacherID)) THEN
          UPDATE Sessions SET Token=token, ExpirationDate=SUBSTRING(expirationDate, 1, 19) WHERE Sessions.TeacherID=teacherID;
          SET s = "Updated";
        ELSE
            INSERT INTO Sessions (TeacherID, Token, ExpirationDate)
                VALUES (teacherID, token, SUBSTRING(expirationDate, 1, 19));
            SET s = "Insert";
        END IF;                                        
    ELSE
        SET s = "User doesnt exist";
    END IF;
    RETURN s;
  END //
DELIMITER ;

###########################################################################################
                                    #IsValidToken
###########################################################################################
DROP FUNCTION IF EXISTS IsValidToken;
DELIMITER //
CREATE FUNCTION IsValidToken(teacherID INT, token VARCHAR(500))
  RETURNS VARCHAR(20)

  BEGIN
    DECLARE s VARCHAR(20);

    IF (EXISTS(SELECT 1 FROM Sessions WHERE Sessions.TeacherID=teacherID AND Sessions.Token=token AND Sessions.ExpirationDate>NOW())) THEN
            SET s = "True";
    ELSE
            SET s = "False";
    END IF;
    RETURN s;
  END //
DELIMITER ;

###########################################################################################
                                    #AreCredentialsValid
###########################################################################################
DROP FUNCTION IF EXISTS AreCredentialsValid;
DELIMITER //
CREATE FUNCTION AreCredentialsValid(email VARCHAR(100), password VARCHAR(100))
  RETURNS VARCHAR(20)

  BEGIN
    DECLARE s VARCHAR(20);

    IF (EXISTS(SELECT 1 FROM Teachers WHERE Teachers.Email=email AND Teachers.Password=password)) THEN
            SET s = "True";
    ELSE
            SET s = "False";
    END IF;
    RETURN s;
  END //
DELIMITER ;