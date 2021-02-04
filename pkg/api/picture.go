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
	err, filename := datastore.DeletePictureRecord(param_pictureid, claim_userid)
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

func AddPictureRoutes(r *gin.Engine, authMiddleware *jwt.GinJWTMiddleware) *gin.Engine {
	// Add protected routes
	auth := r.Group("/auth")
	auth.Use(authMiddleware.MiddlewareFunc())
	{
		// picture
		auth.POST(PictureRoute, HandlePictureCreateRequest)
		auth.DELETE(PictureExistingRoute, HandlePictureDeleteRequest)
	}

	return r
}
