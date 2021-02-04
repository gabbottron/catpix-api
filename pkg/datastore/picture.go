package datastore

import (
	"log"

	"github.com/gabbottron/catpix-api/pkg/datatypes"
)

const (
	InsertPictureQuery string = `INSERT INTO picture 
		(catpixuserid, filename) 
		VALUES($1, $2)
    RETURNING pictureid, catpixuserid, filename, createdate, modifieddate`

	DeletePictureQuery string = `DELETE FROM picture 
    WHERE pictureid = $1 AND catpixuserid = $2 RETURNING filename`
)

func InsertPictureRecord(obj *datatypes.PictureJSON) error {
	// attempt to insert the new picture record
	err := db.QueryRow(InsertPictureQuery, obj.CatPixUserID, obj.FileName).Scan(
		&obj.PictureID, &obj.CatPixUserID, &obj.FileName, &obj.CreateDate, &obj.ModifiedDate)

	if err != nil {
		log.Printf("Error inserting picture record -> %s", err.Error())
		if IsConstraintViolation(err) {
			is_dup, msg := IsDuplicateKeyViolation(err)
			if is_dup {
				return &DatastoreError{msg, ErrorConstraintViolation}
			}
			// Handle the db constraint error
			return &DatastoreError{ErrorConstraintViolationString, ErrorConstraintViolation}
		}
		return err
	}

	return nil
}

func DeletePictureRecord(pictureID int, userID int) (error, string) {
	var filename string

	err := db.QueryRow(DeletePictureQuery, pictureID, userID).Scan(&filename)
	if err != nil {
		return &DatastoreError{ErrorUnknownDatabaseErrorString, ErrorUnknownDatabaseError}, ""
	}

	return nil, filename
}
