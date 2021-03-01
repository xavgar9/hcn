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