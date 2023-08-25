IF EXISTS (SELECT * FROM INFORMATION_SCHEMA.COLUMNS WHERE [TABLE_NAME] = 'product' AND [COLUMN_NAME] = 'description')
BEGIN 
	ALTER TABLE product DROP COLUMN description;
END;
