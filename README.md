# kenshilabsAssignmentKumarSatyam

`Setting Up`- 
-```go mod tidy```
 then run ->
-```go run main.go```

- `POST /signup`: Create a new user account.
- `POST /signin`: Authenticate a user and generate a JWT token with expiry 24hrs.
- `POST /signout`: Invalidate the user's JWT token. Generate new token with expiry as 5 seconds as JWT token are immutable we cannot invalidate it, it can be blacklisted from server side [out of scope for this project]


- `POST /tasks`: Create a new task 
- `GET /tasks`: Retrieve all tasks for the authenticated user.
- `GET /tasks/:id`: Retrieve a specific task by ID 
- `PUT /tasks/:id`: Update a task by ID 
- `DELETE /tasks/:id`: Delete a task by ID 