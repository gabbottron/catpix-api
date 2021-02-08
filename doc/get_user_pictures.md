# Get all Pictures for User

Gets all Pictures owned by a User.

**URL** : `/user/:user_id/pictures`

**Method** : `GET`

**Auth required** : NO

**Permissions required** : None

## Success Response

**Code** : `200 OK`

**Response examples**

```json
[
    {
        "pictureId": 3,
        "catPixUserId": 2,
        "fileName": "44d5034a-6ebf-45d8-8cad-156ff3690cde.jpg",
        "createDate": "2021-02-04T20:25:12.602338Z",
        "modifiedDate": "2021-02-04T20:25:12.602338Z"
    }
]
```

## Notes

*  To access a picture, combine the filename with the API URI like so: http://localhost:8080/catpix/44d5034a-6ebf-45d8-8cad-156ff3690cde.jpg

