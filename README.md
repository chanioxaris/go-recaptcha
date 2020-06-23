# go-recaptcha

[Google reCAPTCHA](https://www.google.com/recaptcha/intro/v3.html) (versions [2](https://developers.google.com/recaptcha/docs/display) and [3](https://developers.google.com/recaptcha/docs/v3)) verification in Golang.

## Install

To get the package:

`go get github.com/chanioxaris/go-recaptcha`

## Examples

reCAPTCHA v2:

- Simple usage with default values (httpClient: http.DefaultClient)

        package main
        
        import (
            "fmt"
        
            "github.com/chanioxaris/go-recaptcha/recaptcha"
        )
        
        func main() {
            rec, err := recaptcha.New("secret", recaptcha.WithVersion(2))
            if err != nil {
                panic(err)
            }
        
            if err = rec.Verify("response"); err != nil {
                panic(err)
            }
        
            fmt.Println("Success")
        }
        
- Simple usage with custom http client

        package main
        
        import (
            "fmt"
        
            "github.com/chanioxaris/go-recaptcha/recaptcha"
        )
        
        func main() {
            customClient := &http.Client{Timeout: time.Second * 10}
            
            rec, err := recaptcha.New("secret", recaptcha.WithVersion(2), recaptcha.WithHTTPClient(customClient))
            if err != nil {
                panic(err)
            }
        
            if err = rec.Verify("response"); err != nil {
                panic(err)
            }
        
            fmt.Println("Success")
        }
        
- Get reCAPTCHA token from request body (`g-recaptcha-response` field)

        import (
        	"fmt"
        	"net/http"
        
        	"github.com/chanioxaris/go-recaptcha/recaptcha"
        )
        
        func Handler(w http.ResponseWriter, r *http.Request) {
        	rec, err := recaptcha.New("secret", recaptcha.WithVersion(2))
        	if err != nil {
        		panic(err)
        	}
        
        	response, err := rec.GetRequestToken(r)
        	if err != nil {
        		panic(err)
        	}
        
        	if err = rec.Verify(response); err != nil {
        		panic(err)
        	}
        
        	fmt.Println("Success")
        }

reCAPTCHA v3:

- Simple usage with default values (version: 3, action: "", score: 0.5, httpClient: http.DefaultClient)

        package main
        
        import (
            "fmt"
    
            "github.com/chanioxaris/go-recaptcha/recaptcha"
        )
    
        func main() {
            rec, err := recaptcha.New("secret")
            if err != nil {
                panic(err)
            }
        
            if err = rec.Verify("response"); err != nil {
                panic(err)
            }
        
            fmt.Println("Success")
        }
        
- Simple usage with custom values

        package main
            
        import (
            "fmt"
    
            "github.com/chanioxaris/go-recaptcha/recaptcha"
        )
        
        func main() {
             customClient := &http.Client{Timeout: time.Second * 10}
             customAction := "custom-action"
             customScore := 0.7
        
            rec, err := recaptcha.New("secret", 
            		recaptcha.WithHTTPClient(customClient), 
            		recaptcha.WithAction(customAction), 
            		recaptcha.WithMinScore(customScore),
            	)
            if err != nil {
                panic(err)
            }
        
            if err = rec.Verify("response"); err != nil {
                panic(err)
            }
        
            fmt.Println("Success")
        }

- Get reCAPTCHA token from request body (`g-recaptcha-response` field)

        package main
                
        import (
            "fmt"
            "net/http"
        
            "github.com/chanioxaris/go-recaptcha/recaptcha"
        )
        
        func Handler(w http.ResponseWriter, r *http.Request) {
            rec, err := recaptcha.New("secret")
            if err != nil {
                panic(err)
            }
        
            response, err := rec.GetRequestToken(r)
            if err != nil {
                panic(err)
            }
        
            if err = rec.Verify(response); err != nil {
                panic(err)
            }
        
            fmt.Println("Success")
        }

Middleware:

- Use middleware in REST API

        package main
    
        import (
            "log"
            "net/http"
        
            "github.com/gorilla/mux"
        
            md "github.com/chanioxaris/go-recaptcha/middleware"
            "github.com/chanioxaris/go-recaptcha/recaptcha"
        )
        
        func main() {
            // Create recaptcha instance.
            rec, err := recaptcha.New("secret")
            if err != nil {
                panic(err)
            }
        
            // Setup router.
            router := mux.NewRouter().StrictSlash(true)
            // Use the recaptcha middleware.
            router.Use(md.Middleware(rec))
        
            // Setup endpoint handler.
            router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
                w.Write([]byte("A reCAPTCHA protected endpoint"))
            })
        
            // Start server.
            log.Fatal(http.ListenAndServe(":8080", router))
        }

## License

go-recaptcha is [MIT licensed](LICENSE).