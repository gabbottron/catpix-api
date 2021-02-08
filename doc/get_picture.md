# Get a Picture by ID

Gets a picture by ID.

**URL** : `/picture/:picture_id`

**Method** : `GET`

**Auth required** : NO

**Permissions required** : None

## Success Response

**Code** : `200 OK`

**Response examples**

```json
{
    "pictureId": 1,
    "catPixUserId": 1,
    "fileName": "0d2e6b70-535a-4718-a243-46a1c7e9a5e6.jpg",
    "createDate": "2021-02-04T20:23:06.925238Z",
    "modifiedDate": "2021-02-04T20:23:06.925238Z"
}
```

## Notes

*  To access the picture, combine the filename with the API URI like so: http://localhost:8080/catpix/0d2e6b70-535a-4718-a243-46a1c7e9a5e6.jpg

