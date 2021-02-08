# Get all Pictures

Gets all pictures in the system.

**URL** : `/pictures`

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
    },
    {
        "pictureId": 2,
        "catPixUserId": 1,
        "fileName": "af2b46fe-3fbf-484d-b37b-160bc572db0d.jpg",
        "createDate": "2021-02-04T20:24:23.406362Z",
        "modifiedDate": "2021-02-04T20:24:23.406362Z"
    },
    {
        "pictureId": 1,
        "catPixUserId": 1,
        "fileName": "0d2e6b70-535a-4718-a243-46a1c7e9a5e6.jpg",
        "createDate": "2021-02-04T20:23:06.925238Z",
        "modifiedDate": "2021-02-04T20:23:06.925238Z"
    }
]
```

## Notes

*  To access a picture, combine the filename with the API URI like so: http://localhost:8080/catpix/af2b46fe-3fbf-484d-b37b-160bc572db0d.jpg

