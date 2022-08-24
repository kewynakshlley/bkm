# bkm

bkm is a simple app to manage books, it was created to put in practice some aspects of the lang such as: package management and organization, pointers, http requests and database comms.


- /books
    * method: POST
    * data: {"title": "Neuromancer", "Author":"William Gibson", "Publication":"Ace"}
    * description: creates a new book
- /books
    * method: GET
    description: get all books
- /books/id
    * methods: GET
    description: get one book
- /books/id 
    * PUT
    * description: updates a book
- /books/id
    * DELETE
    * description: deletes a book


Soon enough I pretend to create a frontend app using React to use this API.

## Built With

* [GORM](https://gorm.io/index.html) - ORM library for golang
* [Gorilla MUX](https://github.com/gorilla/mux) - Requests multiplexer used to HTTP routing and URL matcher
* [Docker](https://www.docker.com/) - Docker-compose used to serve a mysql database


