## GO Testing Because I Can!!!!

This repo contains a quick(very quick and dirty) demo of building a GO api. Remember, I did this in just a few hours so there is definitely some refactoring I would do with more time..

## SETUP
* **PostgreSQL** - Make sure to run the included table script in your PG database and add the connection information into the dal.go file. If you don't, good luck connecting to your DB.
* **That's it...** - Use a test client like Postman to hit the launched API. When you start it your routes will show in the console but I will list them below in case...

## ROUTES
* **Base Route** - http://localhost:8081/v1/api/
* **Get User By ID** - http://localhost:8081/v1/api/webuser/{int-ID}
* **Get User(s) By LastName** - http://localhost:8081/v1/api/webuser/lastname/{string-Name}
* **Update(Patch) User Lucky Number** - http://localhost:8081/v1/api/webuser/{int-ID}/number/{int-luckyNumber}
* **Delete User** - http://localhost:8081/v1/api/webuser/{int-ID}
* **Create(Post) User** - http://localhost:8081/v1/api/webuser
* **Example Create Body**
```
{
    "firstname": "Post",
    "lastname": "Call2",
    "luckynumber": 200,
    "age": 404
}
```

