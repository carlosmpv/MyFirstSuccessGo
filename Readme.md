# My first project GO

### this is my first functional go project

The code is relatively organized but here is the structure for anyone who wants to read:

[App](./app.go): 
* Initializes a default engine
* Initializes database module
* Call migrations
* Call routing 
* Runs the server

[Models/Database](./models): The file handling mysql connections is in the same folder as the models
* database
    * Initializes the data base and return the connector

* models
    * Create tables
    * Migrate

[Routes](./routes): Define routes for each action
* GET /user
    * Send JSON response with all registered users

* POST /login
    * Search on database for matching results on **email** and **password**
    * Returns user's property if succeeded
    * Returns error message if not

* POST /register
    * Search on database for matching results on **email**
    * If it is not registered, then, register the new user
    * Returns user's property if succeeded
    * Returns error message if email is registered or form is invalid


    
