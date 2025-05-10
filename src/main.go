package main

import (
	"fmt"
	"github.com/go-resty/resty/v2"
)

func clear() {
	fmt.Print("\033[2J\033[H")
}

type UserProfile struct {

	Name string `"json:name"`
	Login string `"json:login"`
	Bio interface{} `json:"bio"`

}


type RepoInfo struct {

	Name string `json:"name"`
	Description interface{} `json:"description"`
	Language string `json:"language"`

}

func viewProfile(Username string) {
	client := resty.New()
	user := &UserProfile{}
	reqUrl := fmt.Sprintf("https://api.github.com/users/%s", Username)
	resp, err := client.R().
	SetHeader("Accept", "application/vnd.github.v3+json").
	SetResult(&user).
	Get(reqUrl)

	if err != nil {
		fmt.Println("\033[33m[i]\033[0m unexpected error")
	}

	status := resp.StatusCode()

	if status == 200 {
		clear()
		fmt.Println("\033[33m[i]\033[0m | Name: ", user.Name)
		fmt.Println("\033[33m[i]\033[0m | Login: ", user.Login)
		fmt.Println("\033[33m[i]\033[0m | Description: ", user.Bio)
	} else if status == 404 {
		fmt.Println("[404] Not found.")
	} else {
		fmt.Println("[%d]", status)
	}
	fmt.Scanln()
}

func viewRepo(Usrname string, Reepo string) {
	client := resty.New()
	repo := &RepoInfo{}
	reqUrl := fmt.Sprintf("https://api.github.com/repos/%s/%s", Usrname, Reepo)
	resp, err := client.R().
	SetHeader("Accept", "application/vnd.github.v3+json").
	SetResult(&repo).
	Get(reqUrl)
	
	status := resp.StatusCode()

	if err != nil {
		fmt.Println("\033[33m[i]\033[0m unexpected error.")
	} 

	if status == 200 {
		clear()
		fmt.Println("\033[33m[i]\033[0m | Name: ", repo.Name)
		fmt.Println("\033[33m[i]\033[0m | Description: ", repo.Description)
		fmt.Println("\033[33m[i]\033[0m | Language: ", repo.Language)
	} else if status == 404 {
		fmt.Println("[404] Not found.")
	} else {
		fmt.Printf("{%d}", status)
	}
	fmt.Scanln()
}

func main() {
	var user string
	var repo string
	var chc int

	for {
		clear()
		fmt.Println("\033[33m[i]\033[0m Bem vindo ao gitview!")
		fmt.Println("\033[33m1.\033[0m Visualizar perfil")
		fmt.Println("\033[33m2.\033[0m Visualizar repositório")

		fmt.Print("> ")
		fmt.Scan(&chc)

		if chc == 1 {
			fmt.Print("[Username]$ ")
			fmt.Scan(&user)
			viewProfile(user)
		} else if chc == 2 {
			fmt.Print("[Repo]$ ")
			fmt.Scan(&repo)
			fmt.Print("[Username]$ ")
			fmt.Scan(&user)
			viewRepo(user, repo)
		}
	}
}
