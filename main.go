package main

import (
	"os"

	router "gopls-workspace/router"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)





func setupRouter() *gin.Engine {
	r := gin.Default()
gin.DisableConsoleColor()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	r.Use(gin.ErrorLogger())
	// r.Use(CORSMiddleware())
	config := cors.DefaultConfig()
	config.AllowAllOrigins = true
	// OPTIONS method for ReactJS
	config.AllowMethods = []string{"POST", "GET", "OPTIONS"}
	config.AllowCredentials = true
	config.AllowHeaders = []string{"Origin", "x-access-token", "content-type", "Content-Length", "Authorization", "Cache-Control"}
	config.ExposeHeaders = []string{"Content-Length"}
	r.Use(cors.New(config))
	// r.Use(gzip.Gzip(gzip.DefaultCompression))

	/************/
	/* Users */
	/************/
	router.Init(r)

	return r
}

func setupConfigs() {

	//  os.Setenv("TPAYDBFIN", "postgres://tamscashlessauth:t@m2c@shl3ss@uth3@1@34.68.247.131:5432/tpaydatabase")
os.Setenv("CURRENTDOMAIN", "https://4168-41-169-160-57.ngrok.io")	


os.Setenv("DBTshwane", "postgres://cogjgedlgavael:cf43a86f559ebdd296331ca10991a0bfc87dfcf1fb7c83d3407698719348a669@ec2-18-204-74-74.compute-1.amazonaws.com:5432/d7jnruc4m8g23q")	
os.Setenv("WEBSERVER_PORT", "8080")

}


func main() {
	//Uncommented When Not Debugging
	// gin.SetMode(gin.ReleaseMode)
	// export GIN_MODE=release


    // gocron.Start()
    // s := gocron.NewScheduler()
    // gocron.Every(2).Seconds().Do(c.CheckNewUser)
	//  <-s.Start()

	r := setupRouter()

	setupConfigs()

	r.Run(":" + os.Getenv("WEBSERVER_PORT"))
}
