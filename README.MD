# Dbproto usage example
This is a simple example of how to use the dbproto library in Go to create a complete backend example.
## Prerequisites
Go installed on your machine
Set up the environment variable for the AES key as described [here](https://github.com/Malpizarr/dbproto?tab=readme-ov-file#environment-variable-setup-guide-for-aes-key)


## How to run
1. Clone the repository
2. Navigate to the project directory
3. Run the following command:
```bash
go run main.go
```
The server will start running on port 8080.

## API Endpoints
The application provides the following endpoints:  

### User
- POST /users: Create a new user
- PUT /users: Update an existing user
- GET /users: Get all users
- GET /users/{username}: Get a user by username
- DELETE /users/{username}: Delete a user by username
### Post
- POST /posts: Create a new post
- GET /posts: Get all posts
- GET /posts/{username}: Get posts by username
- PUT /posts: Update an existing post
- DELETE /posts/{id}: Delete a post by id
### Friendship
- POST /friendships: Create a new friendship
- GET /friendships/{id}: Get a friendship by id
- GET /friendships: Get all friendships
- PUT /friendships/accept/{id}: Accept a friendship by id
- PUT /friendships/reject/{id}: Reject a friendship by id
- DELETE /friendships/{id}: Delete a friendship by id
## Testing the API
You can use curl or any API testing tool such as Postman to test the API. Here are some example curl commands:
```
# Replace <username>, <email>, <password>, <title>, <content>, <status>, and <id> with actual values
```
```bash
# Create a user
curl -X POST -H "Content-Type: application/json" -d '{"username":"<username>", "email":"<email>", "password":"<password>"}' http://localhost:8080/users
```
```bash
# Update a user
curl -X PUT -H "Content-Type: application/json" -d '{"username":"<username>", "email":"<email>", "password":"<password>"}' http://localhost:8080/users
```
```bash
# Get a user
curl -X GET http://localhost:8080/users/<username>
```
```bash
# Delete a user
curl -X DELETE http://localhost:8080/users/<username>
```
```bash
# Create a post
curl -X POST -H "Content-Type: application/json" -d '{"username":"<username>", "title":"<title>", "content":"<content>"}' http://localhost:8080/posts
``` 
```bash
# Update a post
curl -X PUT -H "Content-Type: application/json" -d '{"id":<id>, "username":"<username>", "title":"<title>", "content":"<content>"}' http://localhost:8080/posts
``` 
```bash
# Delete a post
curl -X DELETE http://localhost:8080/posts/<id>
``` 
```bash
# Create a friendship
curl -X POST -H "Content-Type: application/json" -d '{"user1":"<username1>", "user2":"<username2>", "status":"<status>"}' http://localhost:8080/friendships
``` 
```bash
# Accept a friendship
curl -X PUT http://localhost:8080/friendships/accept/<id>
``` 
```bash
# Reject a friendship
curl -X PUT http://localhost:8080/friendships/reject/<id>
``` 
```bash
# Delete a friendship
curl -X DELETE http://localhost:8080/friendships/<id>
```
#### **_Please replace username, email, password, title, content, status, and id with actual values._**

