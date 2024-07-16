# web-app-starter
This is a starter project for any webapp backend I want to build in GO.

This uses GO-CHI this is a plugin built on the default http package.
Although you can build this as  a fullstack app and make it monolith, but you kmow im just so used with microservices. Since firebase is free for react apps i love just building react frontend and make it static.

# External Packages used in this project, and should be enough to run all your server

    GOCHI -  [Github]("https://github.com/go-chi/chi")
    PQ(TO LOAD POSTGRES) [Github]("https://github.com/lib/pq")
    GODOENV(TO LOAD ENV) [Github]("github.com/joho/godotenv")

    for these golang packages after cloning the project run "go mod tidy"

    I use sqlc to generate for my internal package 
    sqlc configuration is in "./sqlc.yaml"
    You can change your driver if you dont want to use postgres and thatbmean you change https://github.com/lib/pq to your desired driver.

    I have 1 example each for both the sql schema and queries, you can create more base on your needs.
    I use sqlc from my terminal - install using "go install github.com/sqlc-dev/sqlc/cmd/sqlc@latest"
    Run sqlc generate to generate all your internal code, i will leave the one for the two example, you can remove them base on your need and sqlc will update the internal code.

    With this you have your db in your c for all the http handler and you can talk to db using c.DB.YourFunction