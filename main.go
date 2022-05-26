package main

import "github.com/coderile/shortenurl-go/app"

// func CreateShortUrl(res http.ResponseWriter, req *http.Request) {
// 	cl := db.DbConnection()
// 	res.Header().Add("content-type", "application/json")
// 	var shortUrl models.ShortURL
// 	fmt.Println(req.Body)
// 	json.NewDecoder(req.Body).Decode(&shortUrl)
// 	collection := cl.Database("shorturl").Collection("short")
// 	result, _ := collection.InsertOne(context.TODO(), shortUrl)
// 	json.NewEncoder(res).Encode(result)
// }

func main() {
	app.StartApp()
	// db connection

	// ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)

	// router := mux.NewRouter()
	// router.HandleFunc("/short_url", CreateShortUrl).Methods("POST")
	// fmt.Println("Server started at port 8999")
	// http.ListenAndServe(":8999", router)
	// if err := client.Ping(context.TODO(), readpref.Primary()); err != nil {
	// 	panic(err)
	// }
	// fmt.Println("Successfully connected and pinged.")

}
