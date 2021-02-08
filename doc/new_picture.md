# New Picture

Create a new Picture.

**URL** : `/auth/picture`

**Method** : `POST`

**Auth required** : YES

**Permissions required** : None

## Success Response

**Code** : `200 OK`

**Response examples**

```json
{
    "pictureId": 3,
    "catPixUserId": 2,
    "fileName": "44d5034a-6ebf-45d8-8cad-156ff3690cde.jpg",
    "createDate": "2021-02-04T20:25:12.602338Z",
    "modifiedDate": "2021-02-04T20:25:12.602338Z"
}
```

## Notes

*  Body of the request should be form-data with a key called 'file' set to the local image file. See Postman collection referenced in README.md for an example.
