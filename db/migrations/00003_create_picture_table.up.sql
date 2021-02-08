CREATE TABLE IF NOT EXISTS picture (
  PictureId               SERIAL       NOT NULL,  
  CatPixUserId            INT          NOT NULL,
  FileName                VARCHAR(200) NULL,
  CreateDate              TIMESTAMP    NOT NULL CONSTRAINT "DF_Picture_CreateDate"   DEFAULT (NOW()),
  ModifiedDate            TIMESTAMP    NOT NULL CONSTRAINT "DF_Picture_ModifiedDate" DEFAULT (NOW()),
  CONSTRAINT "PK_Picture_PictureId"    PRIMARY KEY (PictureId),
  CONSTRAINT "FK_Picture_CatPixUserId" FOREIGN KEY (CatPixUserId) REFERENCES catpix_user (CatPixUserId),
  UNIQUE(FileName)
);

CREATE TRIGGER trigger_record_changed
  BEFORE UPDATE ON picture
  FOR EACH ROW
  EXECUTE PROCEDURE record_changed();
