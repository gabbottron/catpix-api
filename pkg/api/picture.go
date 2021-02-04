package api

import (
	"log"
	"net/http"
	"os"
	"path/filepath"

	"github.com/gabbottron/catpix-api/pkg/datastore"
	"github.com/gabbottron/catpix-api/pkg/datatypes"
	"github.com/gabbottron/catpix-api/pkg/jwt"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

const (
	PictureRoute         string = "/picture"
	PicturesRoute        string = "/pictures"
	PictureExistingRoute string = "/picture/:pictureid"
	PicturesByUserRoute  string = "/user/:userid/pictures"
)

func HandlePictureCreateRequest(c *gin.Context) {
	// Get the userID from the claims
	claim_userid := GetIDFromClaim(c)

	// attempt to get the file from the payload
	file, err := c.FormFile("file")
	if err != nil {
		log.Printf("HandlePictureCreateRequest -> Could not load FormFile: %s", err.Error())
		ReplyBadRequest(c, "No file in the payload!")
		return
	}

	// Retrieve file information
	extension := filepath.Ext(file.Filename)

	// Generate random file name for the new uploaded file so it doesn't override an old file with same name
	newFileName := uuid.New().String() + extension

	// File received, save it to drive
	// (normally would save to something like S3, but this is a simplification for this
	//  demo project. If it was production I'd break out a library with abstracted
	//  cloud storage functionality and call methods in there from here)
	if err := c.SaveUploadedFile(file, "/app/pictures/"+newFileName); err != nil {
		log.Printf("HandlePictureCreateRequest -> Could not save file to drive: %s", err.Error())
		ReplyInternalServerError(c, "Could not save your file!")
		return
	}

	// create picture object
	var pictureData datatypes.PictureJSON

	// load values
	pictureData.CatPixUserID = claim_userid
	pictureData.FileName = new(string)
	*pictureData.FileName = newFileName

	err = datastore.InsertPictureRecord(&pictureData)
	if err != nil {
		log.Printf("HandlePictureCreateRequest -> Could not save file record to database: %s", err.Error())
		ReplyInternalServerError(c, "Could not save your file!")
		return
	}

	c.JSON(http.StatusOK, pictureData)
	return
}

func HandlePictureDeleteRequest(c *gin.Context) {
	// Get the userID from the claims
	claim_userid := GetIDFromClaim(c)

	// Validate pictureID param in URI
	param_pictureid, err := GetIntFromParam(c, "pictureid")
	if err != nil {
		log.Println("HandlePictureDeleteRequest -> Invalid pictureID supplied as URI parameter!")
		ReplyBadRequest(c, "Malformed pictureID in URI!")
		return
	}

	// attempt to delete the record from the database
	filename, err := datastore.DeletePictureRecord(param_pictureid, claim_userid)
	if err != nil {
		log.Printf("HandlePictureDeleteRequest -> Unable to delete picture with id (%d) for user (%d)", param_pictureid, claim_userid)
		ReplyNotFound(c, "A picture with that ID doesn't exist, or you do not own it!")
		return
	}

	// if we made it here, record deletion was successful,
	// now remove the file from disk
	err = os.Remove("/app/pictures/" + filename)
	if err != nil {
		log.Println("HandlePictureDeleteRequest -> Failed to remove picture file from disk!")
		// I'm not returning an error to the user at this point, because the record
		// was removed from the database, so this error is irrelevant to the end user
	}

	c.JSON(http.StatusOK, "Picture was deleted!")

	return
}

func HandlePictureUpdateRequest(c *gin.Context) {
	// Get the userID from the claims
	claim_userid := GetIDFromClaim(c)

	// Validate pictureID param in URI
	param_pictureid, err := GetIntFromParam(c, "pictureid")
	if err != nil {
		log.Println("HandlePictureUpdateRequest -> Invalid pictureID supplied as URI parameter!")
		ReplyBadRequest(c, "Malformed pictureID in URI!")
		return
	}

	// find the picture record in the database
	picture, err := datastore.GetPictureByIDProtected(param_pictureid, claim_userid)
	if err != nil {
		log.Printf("HandlePictureUpdateRequest -> Picture not found, error: %s", err.Error())
		ReplyNotFound(c, "You don't own a picture with that ID!")
		return
	}

	// attempt to get the file from the payload
	file, err := c.FormFile("file")
	if err != nil {
		log.Printf("HandlePictureUpdateRequest -> Could not load FormFile: %s", err.Error())
		ReplyBadRequest(c, "No file in the payload!")
		return
	}

	// Retrieve file information
	extension := filepath.Ext(file.Filename)

	// Generate random file name for the new uploaded file so it doesn't override an old file with same name
	newFileName := uuid.New().String() + extension

	// File received, save it to drive
	// (normally would save to something like S3, but this is a simplification for this
	//  demo project. If it was production I'd break out a library with abstracted
	//  cloud storage functionality and call methods in there from here)
	if err := c.SaveUploadedFile(file, "/app/pictures/"+newFileName); err != nil {
		log.Printf("HandlePictureUpdateRequest -> Could not save file to drive: %s", err.Error())
		ReplyInternalServerError(c, "Could not save your file!")
		return
	}

	// capture the original filename so we can delete it from disk
	// if the datastore update is successful
	old_filename := *picture.FileName

	// set the new filename in the picture object
	*picture.FileName = newFileName

	// at this point, the new file is saved to the disk, now we will update
	// the database record to reference this new file
	err = datastore.UpdatePictureRecord(&picture)
	if err != nil {
		log.Printf("HandlePictureUpdateRequest -> Picture record update failed, error: %s", err.Error())
		ReplyInternalServerError(c, "Could not save your file!")
		return
	}

	/*
			  At this point we have managed to write the new file to disk
			  and update the database record, so we can safely remove
		    the old file.
		    NOTE: I didn't simply overwrite the old filename
			  with the new file because the file type could be different,
		    thus the extension should change.
	*/
	err = os.Remove("/app/pictures/" + old_filename)
	if err != nil {
		log.Println("HandlePictureUpdateRequest -> Failed to remove original picture file from disk!")
		// I'm not returning an error to the user at this point, because the new
		// file was written to disk
	}

	c.JSON(http.StatusOK, picture)

	return
}

func HandlePictureRequest(c *gin.Context) {
	// Validate pictureID param in URI
	param_pictureid, err := GetIntFromParam(c, "pictureid")
	if err != nil {
		log.Println("HandlePictureRequest -> Invalid pictureID supplied as URI parameter!")
		ReplyBadRequest(c, "Malformed pictureID in URI!")
		return
	}

	// find the picture record in the database
	picture, err := datastore.GetPictureByID(param_pictureid)
	if err != nil {
		log.Printf("HandlePictureRequest -> Picture not found, error: %s", err.Error())
		ReplyNotFound(c, "Could not find picture with that ID!")
		return
	}

	c.JSON(http.StatusOK, picture)

	return
}

func HandlePicturesRequest(c *gin.Context) {
	// find the picture records in the database
	pictures, err := datastore.GetAllPictures(0)
	if err != nil {
		log.Printf("HandlePicturesRequest -> Pictures not found, error: %s", err.Error())
		ReplyNotFound(c, "Could not find any pictures!")
		return
	}

	c.JSON(http.StatusOK, pictures)

	return
}

func HandlePicturesByUserRequest(c *gin.Context) {
	// Validate userID param in URI
	param_userid, err := GetIntFromParam(c, "userid")
	if err != nil {
		log.Println("HandlePicturesByUserRequest -> Invalid userID supplied as URI parameter!")
		ReplyBadRequest(c, "Malformed userID in URI!")
		return
	}

	// find the picture records in the database
	pictures, err := datastore.GetAllPictures(param_userid)
	if err != nil {
		log.Printf("HandlePicturesByUserRequest -> No pictures found for user with id %d, error: %s", param_userid, err.Error())
		ReplyNotFound(c, "Could not find any pictures for that user!")
		return
	}

	c.JSON(http.StatusOK, pictures)

	return
}

func AddPictureRoutes(r *gin.Engine, authMiddleware *jwt.GinJWTMiddleware) *gin.Engine {
	// Unprotected routes
	r.GET(PictureExistingRoute, HandlePictureRequest)
	r.GET(PicturesRoute, HandlePicturesRequest)
	r.GET(PicturesByUserRoute, HandlePicturesByUserRequest)

	// Add protected routes
	auth := r.Group("/auth")
	auth.Use(authMiddleware.MiddlewareFunc())
	{
		// must be logged in for these
		auth.POST(PictureRoute, HandlePictureCreateRequest)
		auth.DELETE(PictureExistingRoute, HandlePictureDeleteRequest)
		auth.PUT(PictureExistingRoute, HandlePictureUpdateRequest)
	}

	return r
}
