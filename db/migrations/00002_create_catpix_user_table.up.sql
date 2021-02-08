CREATE TABLE IF NOT EXISTS catpix_user (
  CatPixUserId            SERIAL       NOT NULL,
  UserName                VARCHAR(100) NOT NULL CHECK (char_length(UserName)  >= 3)   CONSTRAINT "DF_CatPixUser_UserName"   DEFAULT(null),
  Password                VARCHAR(100) NOT NULL CHECK (char_length(Password)  >= 8)   CONSTRAINT "DF_CatPixUser_Password"   DEFAULT(null),
  UserActiveStatus        BOOLEAN      NOT NULL CONSTRAINT "DF_CatPixUser_UserStatus"   DEFAULT (TRUE),
  CreateDate              TIMESTAMP    NOT NULL CONSTRAINT "DF_CatPixUser_CreateDate"   DEFAULT (NOW()),
  ModifiedDate            TIMESTAMP    NOT NULL CONSTRAINT "DF_CatPixUser_ModifiedDate" DEFAULT (NOW()),
  CONSTRAINT "PK_CatPixUser_CatPixUserId" PRIMARY KEY (CatPixUserId),
  UNIQUE(UserName)
);

CREATE TRIGGER trigger_record_changed
  BEFORE UPDATE ON catpix_user
  FOR EACH ROW
  EXECUTE PROCEDURE record_changed();
