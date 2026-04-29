package main

func main() {
	db, err := ConnectDB()
	if err != nil {
		panic(err)
	}
	defer db.Close()

	Router(db)
}
