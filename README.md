## Todo-backend
- This repository includes back-end part of Todo assignment. 

## Technologies and Tools
- Golang
- Fiber
- Mockgen
- Pact

## Scripts
To run project 
- go run main.go

To run back-end unit tests 
- go test -v

To run CDC provider test
- go test -run TestProvider ./ -v -count 1


## Related repositories
- Acceptance Test Repo: https://gitlab.com/furkan.moda/todo-acceptance-test
  Tests are pass in this repo when its runned locally but not in the pipeline because of chromium error.
- Back-end Repo: https://gitlab.com/furkan.moda/todo-backend

## Author
- Furkan Küçük (https://www.linkedin.com/in/furkan-k%C3%BC%C3%A7%C3%BCk-11b84b21b/