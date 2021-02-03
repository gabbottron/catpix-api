#!/bin/bash
set -e

psql -v ON_ERROR_STOP=1 --username "$POSTGRES_USER" --dbname "$POSTGRES_DB" <<-EOSQL
    CREATE USER catpix WITH PASSWORD 'mypass';
    CREATE DATABASE catpix;
    GRANT ALL PRIVILEGES ON DATABASE catpix TO catpix;

    \c catpix

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

    GRANT ALL PRIVILEGES ON ALL TABLES IN SCHEMA public TO catpix;
    GRANT USAGE, SELECT ON ALL SEQUENCES IN SCHEMA public TO catpix;
EOSQL
