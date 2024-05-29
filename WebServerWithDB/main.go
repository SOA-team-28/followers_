package main

func startServer() {
	/*
		database, err := db.InitDB()
		if err != nil {
			log.Fatal("Failed to initialize database:", err)
		}

		router := mux.NewRouter().StrictSlash(true)

		followerHandler := handler.NewUserHandler(database)
		followerHandler.RegisterRoutes(router)

		router.PathPrefix("/").Handler(http.FileServer(http.Dir("./static")))
		log.Println("Server is running on port", db.Port)
		log.Fatal(http.ListenAndServe(":8082", router))
	*/

}

func main() {

	startServer()
}
