# Update Picture

Update an existing Picture.

**URL** : `/auth/picture/:picture_id`

**Method** : `PUT`

**Auth required** : YES

**Permissions required** : Must own picture being updated

## Success Response

**Code** : `200 OK`

**Response examples**

```json
{
    "pictureId": 1,
    "catPixUserId": 1,
    "fileName": "6581fd86-cfe4-4f5c-bcc5-4547679c9d45.jpg",
    "createDate": "2021-02-04T19:43:02.762944Z",
    "modifiedDate": "2021-02-04T19:47:01.50785Z"
}
```

## Notes

*  Body of the request should be form-data with a key called 'file' set to the local image file. See Postman collection referenced in README.md for an example.
