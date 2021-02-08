# Log In

Log in and get JWT credentials.

**URL** : `/login`

**Method** : `POST`

**Auth required** : NO

**Permissions required** : None

## Success Response

**Code** : `200 OK`

**Content examples**

This is a default user loaded into the database at container creation. You may also create your own user if you like.

```json
{
    "username": "codeacademy",
    "password": "d3fHireTh@tGeoffGuy"
}
```

## Notes

* You will need to provide the JWT in the HTTP headers of any request to an endpoint requiring authorization. Authorization: Bearer {{token}}. This is done in a macro in the provided Postman collection referenced in README.md.
