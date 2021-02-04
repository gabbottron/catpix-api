package api

import (
	"log"
	"net/http"
	"path/filepath"

	"github.com/gabbottron/catpix-api/pkg/datastore"
	"github.com/gabbottron/catpix-api/pkg/datatypes"
	"github.com/gabbottron/catpix-api/pkg/jwt"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

const (
	PictureRoute  string = "/picture"
	PicturesRoute string = "/pictures"
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

func AddPictureRoutes(r *gin.Engine, authMiddleware *jwt.GinJWTMiddleware) *gin.Engine {
	// Add protected routes
	auth := r.Group("/auth")
	auth.Use(authMiddleware.MiddlewareFunc())
	{
		// picture
		auth.POST(PictureRoute, HandlePictureCreateRequest)
	}

	return r
}
