# chi-swagger
chi middleware to automatically generate RESTful API documentation with Swagger 2.0.

## why use [swaggo/swag](https://github.com/swaggo/swag)?
- It provides a [declarative](https://swaggo.github.io/swaggo.io/declarative_comments_format/) way to put to put comments which can be used fo swagger docs.

- Easy to understand/implement and doesn't require to write any code.


## Swag Usage

- Add comments on your API handlers using the declarative syntax explained [here](https://swaggo.github.io/swaggo.io/declarative_comments_format/)

-  Install [swaggo/swag](https://github.com/swaggo/swag)?

```
        go get -u github.com/swaggo/swag/cmd/swag
```

- Run the Swag in your Go project root folder which contains `main.go` file.

``` 
        swag init
```

If your `main.go` file is not in root but uses the models defined in root, you can provide the path of `main.go` file.

```
        swag init -d "./" -g "$FOLDER_NAME/main.go"
```

Swag will parse comments and generate required files(docs folder and docs/doc.go).

## Using chi with swagger

[chi](https://github.com/go-chi/chi) is a golang f/w for lightweight, idiomatic and composable router for building Go HTTP services

Chi's [mount](https://github.com/go-chi/chi/blob/master/mux.go#L279) utility will come handy for handling swagger

Example `main.go` with chi implementation.

```
// @title chi-swagger example APIs
// @version 1.0
// @description chi-swagger example APIs
// @BasePath /
func main() {
	var timeout = 2 * time.Minute

	var routes = []router.Route{
		router.Route{
			Method:      "GET",
			Path:        "/",
			HandlerFunc: handlers.RootHandler,
		},
	}

	log.Println("Launching the server")
	r := router.NewRouter(routes)
	r.Mount("/swagger", httpSwagger.WrapHandler)

	server := http.Server{
		ReadTimeout:  timeout,
		WriteTimeout: timeout,
		Handler:      r,
	}

	err := server.ListenAndServe()
	if err != nil {
		log.Printf("Failed to launch api server:%+v\n", err)
	}
}

```
