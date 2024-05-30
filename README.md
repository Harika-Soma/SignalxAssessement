1. The project is written in golang, graphql and mongodb
2. go mod init github.com/Harika-Soma/SignalxAssessment-SupplyChain
3. cd SignalxAssessment-SupplyChain
4. First we need to initialize the go mod using, go mod init <project-name/package-name> [supplychain]
5. Next, we need to install the graphql packages using, go get -d github.com/99designs/gqlgen@VERSION
6. Next, we need to initialize the graphql server using, go run github.com/99designs/gqlgen init
7. Then we will find the project structure with graphql files like graph packages, server files, models and schemas.
8. Next, we need to define our schema according to project.
9. Next, need to generate the schema when ever schema changes using, go run github.com/99designs/gqlgen generate
10. By generating the above command some of the other packages also need to get or install using, go get <packages>.
11. After that we need to initialize the hooks/bson.go file, which is used to generate bson command for the models in graph folder to access data from the mongodb database
12. Next, we need to add database connection folders, errors and logs folders.
13. Next, we need to add store function to read the data fromt the database.
14. After generating the schema resolver files will be generated according to our queries or APIs in graph folder.
15. Next, we need to implement the logic and connection to the store function to read the data.
16. Next, implemented authentication and authorization using the createUserLogin function with jwt tokens.
17. While accessing the APIs we need to call the API createUserLogin to get the token value and need to access the other APIs.
18. Also, Implemented the errors and logs in store function and in resolver function as well.
19. Also, added .env file to access the mongodb connection.
20. Then running the server. go run server.go
21. The server is running following line will excute.
2024/05/30 14:39:18 connect to http://localhost:8080/ for GraphQL playground
    We can test the APIs in the Graphql Playground.
