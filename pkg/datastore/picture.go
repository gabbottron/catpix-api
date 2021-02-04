package datastore

import (
	"database/sql"
	"log"
	"strconv"

	"github.com/gabbottron/catpix-api/pkg/datatypes"
)

const (
	SelectPictureQuery string = `SELECT pictureid, catpixuserid, filename, 
      createdate, modifieddate
    FROM picture WHERE pictureid = $1`

	SelectPictureProtectedQuery string = `SELECT pictureid, catpixuserid, filename, 
    createdate, modifieddate
    FROM picture WHERE pictureid = $1 AND catpixuserid = $2`

	SelectAllPicturesQuery string = `SELECT pictureid, catpixuserid, filename, 
    createdate, modifieddate FROM picture 
    ORDER BY modifieddate DESC`

	SelectAllPicturesByUserQuery string = `SELECT pictureid, catpixuserid, filename, 
    createdate, modifieddate FROM picture WHERE catpixuserid = $1 
    ORDER BY modifieddate DESC`

	InsertPictureQuery string = `INSERT INTO picture 
		(catpixuserid, filename) 
		VALUES($1, $2)
    RETURNING pictureid, catpixuserid, filename, createdate, modifieddate`

	DeletePictureQuery string = `DELETE FROM picture 
    WHERE pictureid = $1 AND catpixuserid = $2 RETURNING filename`

	UpdatePictureFileQuery string = `UPDATE picture 
    SET filename = $3 
    WHERE pictureid = $1 AND catpixuserid = $2 
    RETURNING pictureid, catpixuserid, filename, createdate, modifieddate`
)

func GetPictureByID(pictureID int) (datatypes.PictureJSON, error) {
	var obj datatypes.PictureJSON

	err := db.QueryRow(SelectPictureQuery, pictureID).Scan(
		&obj.PictureID, &obj.CatPixUserID, &obj.FileName, &obj.CreateDate, &obj.ModifiedDate)

	if err != nil {
		log.Printf("Error fetching picture record -> %s", err.Error())
		if err == sql.ErrNoRows {
			return obj, &DatastoreError{ErrorRecordNotFoundString, ErrorRecordNotFound}
		}
	}

	return obj, err
}

// this will only return the picture if you own it
func GetPictureByIDProtected(pictureID int, userID int) (datatypes.PictureJSON, error) {
	var obj datatypes.PictureJSON

	err := db.QueryRow(SelectPictureProtectedQuery, pictureID, userID).Scan(
		&obj.PictureID, &obj.CatPixUserID, &obj.FileName, &obj.CreateDate, &obj.ModifiedDate)

	if err != nil {
		log.Printf("Error fetching picture record -> %s", err.Error())
		if err == sql.ErrNoRows {
			return obj, &DatastoreError{ErrorRecordNotFoundString, ErrorRecordNotFound}
		}
	}

	return obj, err
}

func GetAllPictures(userID int) ([]datatypes.PictureJSON, error) {
	results := make([]datatypes.PictureJSON, 0)

	var rows *sql.Rows
	var err error

	if userID < 1 {
		rows, err = db.Query(SelectAllPicturesQuery)
	} else {
		rows, err = db.Query(SelectAllPicturesByUserQuery, userID)
	}

	if err != nil {
		log.Println(err)
	}
	defer rows.Close()

	// The count of pictures processed
	count := 0

	for rows.Next() {
		var obj datatypes.PictureJSON

		err = rows.Scan(&obj.PictureID, &obj.CatPixUserID,
			&obj.FileName, &obj.CreateDate, &obj.ModifiedDate)

		// If the scan was successful, load the row
		if err == nil {
			results = append(results, obj)
			count++
		}
	}

	// Show the count of rows
	log.Println("Rows returned: " + strconv.Itoa(count))

	if err = rows.Err(); err != nil {
		// Abnormal termination of the rows loop
		// close should be called automatically in this case
		log.Println(err)
	}

	return results, nil
}

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

func DeletePictureRecord(pictureID int, userID int) (string, error) {
	var filename string

	err := db.QueryRow(DeletePictureQuery, pictureID, userID).Scan(&filename)
	if err != nil {
		return "", &DatastoreError{ErrorUnknownDatabaseErrorString, ErrorUnknownDatabaseError}
	}

	return filename, nil
}

func UpdatePictureRecord(obj *datatypes.PictureJSON) error {
	err := db.QueryRow(UpdatePictureFileQuery, obj.PictureID, obj.CatPixUserID, obj.FileName).Scan(
		&obj.PictureID, &obj.CatPixUserID, &obj.FileName, &obj.CreateDate, &obj.ModifiedDate)
	if err != nil {
		log.Printf("Error updating picture record -> %s", err.Error())
		if err == sql.ErrNoRows {
			return &DatastoreError{ErrorRecordNotFoundString, ErrorRecordNotFound}
		}
		if IsConstraintViolation(err) {
			// Handle the db constraint error
			return &DatastoreError{ErrorConstraintViolationString, ErrorConstraintViolation}
		}
		return err
	}

	return nil
}
