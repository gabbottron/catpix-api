# New User

Create a new user.

**URL** : `/user`

**Method** : `POST`

**Auth required** : NO

**Permissions required** : None

## Success Response

**Code** : `200 OK`

**Content examples**

Creates a new user.

```json
{
    "username": "codeacademy2",
    "password": "d3fHireTh@tGeoffGuy"
}
```

## Notes

*  Note, password requirements are fairly strict, you will need one lowercase, one uppercase, one number and one symbol. The password must also be at least 8 characters long and contain no runs of repeated characters.
