#### Prerequisites:
1. execute `go build .`

#### To run from command line:
1. execute `./gotasks task add {title} {description}`
2. execute `./gotasks task get 0` to list your tasks
3. execute `./gotasks task complete {taskID}`

#### To run as web server:
1. execute `./gotasks web`
2. send HTTP requests
   1. to create a user
   ```http request
    POST localhost:8080/users
    Content-Type: application/json
    
    {
        "name": "Bob Loblaw",
        "email": "bob@loblaw.au.com",
        "password": "xyz"
    }
    ```
   2. to login and get token to send along with subsequent requests in `Authorization` header
   ```http request
    POST localhost:8080/login
    Content-Type: application/json
    
    {
        "id": :userID,
        "password": "xyz"
    }
    ```
   3. to get user
   ```http request
    GET localhost:8080/users/:userID
    Content-Type: application/json
    Authorization: Bearer xyz
    ```
   4. to create a task
   ```http request
    POST localhost:8080/tasks
    Content-Type: application/json
    Authorization: Bearer xyz
   
    {
        "title": "Feed Dog",
        "description": "myomomym"
    }
    ```
   5. to complete a task
   ```http request
    POST localhost:8080/tasks/finish/:taskID
    Content-Type: application/json
    Authorization: Bearer xyz
    ```