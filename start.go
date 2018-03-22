package main

import (
	"fmt"
	"github.com/exluap/translit"
	"html/template"
	"log"
	"net/http"
	"strings"

	ps "github.com/gorillalabs/go-powershell"
	"github.com/gorillalabs/go-powershell/backend"

	"encoding/json"
	"github.com/getsentry/raven-go"
	"os"
)

type Config struct {
	LogServer struct {
		DSN string `json:"DSN"`
	} `json:"log_server"`
}

func LoadConfiguration(file string) Config {
	var config Config
	configFile, err := os.Open(file)
	defer configFile.Close()
	if err != nil {
		raven.CaptureErrorAndWait(err, nil)
		log.Println(err.Error())
	}
	jsonParser := json.NewDecoder(configFile)
	jsonParser.Decode(&config)
	return config
}

func Init() {
	settings := LoadConfiguration("config.json")
	raven.SetDSN(settings.LogServer.DSN)
}

func login(w http.ResponseWriter, r *http.Request) {
	log.Println("method:", r.Method) //get request method

	if r.Method == "GET" {
		t, err := template.ParseFiles("index.gtpl")
		err2 := t.Execute(w, nil)

		if err != nil {
			raven.CaptureErrorAndWait(err, nil)
			log.Panic(err)
		}

		if err2 != nil {
			raven.CaptureErrorAndWait(err, nil)
			log.Panic(err)
		}

	} else {
		r.ParseForm()
		log.Println("Получил данные, начинаю работать с Active Directory")
		log.Println("method: ", r.Method)

		// logic part of log in
		/*
		        log.Println("fam:", r.Form["surname"])
		        log.Println("name:", r.Form["firstname"])
		        log.Println("otch:", r.Form["endname"])
		        log.Println("group:", r.Form["univergroup"])
				log.Println("kaf:", r.Form["univerkafedra"])
				log.Println("email:", r.Form["email"])
				log.Println("pass:", r.Form["password"])**/

		fmt.Fprintf(w, "Твой логин: "+translit.Ru(strings.Join(r.Form["firstname"], " "))+"."+translit.Ru(strings.Join(r.Form["surname"], " "))+"\n")
		fmt.Fprintf(w, "Пароль: Qwerty123"+" (этот пароль временный и работает только на первый вход) \n")
		fmt.Fprintf(w, "Теперь ты можешь входить с этими данными в Windows и полноценно работать :) \n Хорошего тебе дня!")
		log.Println("New-ADUser -Name '" + translit.Ru(strings.Join(r.Form["firstname"], " ")) + "." + translit.Ru(strings.Join(r.Form["surname"], " ")) + "' -AccountPassword(ConvertTo-SecureString 'Qwerty123' -AsPlainText -Force) -ChangePasswordAtLogon 1 -Company '" + translit.Ru(strings.Join(r.Form["univerkafedra"], " ")) + "' -Department '" + translit.Ru(strings.Join(r.Form["univergroup"], " ")) + "' -DisplayName '" + translit.Ru(strings.Join(r.Form["surname"], " ")) + " " + translit.Ru(strings.Join(r.Form["firstname"], " ")) + " " + translit.Ru(strings.Join(r.Form["endname"], " ")) + "' -EmailAddress '" + strings.Join(r.Form["email"], " ") + "' -Enabled 1 -GivenName '" + translit.Ru(strings.Join(r.Form["firstname"], " ")) + " " + translit.Ru(strings.Join(r.Form["endname"], " ")) + "' -SurName '" + translit.Ru(strings.Join(r.Form["surname"], " ")) + " " + "' -ProfilePath '\\\\POSEIDON\\movedUsers$\\%username%' -SamAccountName '" + translit.Ru(strings.Join(r.Form["firstname"], " ")) + "." + translit.Ru(strings.Join(r.Form["surname"], " ")) + "' -UserPrincipalName '" + translit.Ru(strings.Join(r.Form["firstname"], " ")) + "." + translit.Ru(strings.Join(r.Form["surname"], " ")) + "@ippo.mirea.ru' -MobilePhone" + strings.Join(r.Form["tel"], " ") + " -Path 'OU=PersonaledUsers,DC=ippo,DC=mirea,DC=ru")

		// choose a backend
		back := &backend.Local{}

		// start a local powershell process
		shell, err := ps.New(back)
		if err != nil {
			fmt.Fprintf(w, "Извини, но что-то пошло не так. Позови администратора и он решит твой вопрос")
			raven.CaptureErrorAndWait(err, nil)
			log.Panic(err)
		}
		defer shell.Exit()

		// ... and interact with it
		log.Println("Пытаюсь через PowerShell добавить пользователя")
		stdout, _, err := shell.Execute("New-ADUser -Name '" + translit.Ru(strings.Join(r.Form["firstname"], " ")) + "." + translit.Ru(strings.Join(r.Form["surname"], " ")) + "' -AccountPassword(ConvertTo-SecureString 'Qwerty123' -AsPlainText -Force) -ChangePasswordAtLogon 1 -Company '" + translit.Ru(strings.Join(r.Form["univerkafedra"], " ")) + "' -Department '" + translit.Ru(strings.Join(r.Form["univergroup"], " ")) + "' -DisplayName '" + translit.Ru(strings.Join(r.Form["surname"], " ")) + " " + translit.Ru(strings.Join(r.Form["firstname"], " ")) + " " + translit.Ru(strings.Join(r.Form["endname"], " ")) + "' -EmailAddress '" + strings.Join(r.Form["email"], " ") + "' -Enabled 1 -GivenName '" + translit.Ru(strings.Join(r.Form["firstname"], " ")) + " " + translit.Ru(strings.Join(r.Form["endname"], " ")) + "' -SurName '" + translit.Ru(strings.Join(r.Form["surname"], " ")) + " " + "' -ProfilePath '\\\\POSEIDON\\movedUsers$\\%username%' -SamAccountName '" + translit.Ru(strings.Join(r.Form["firstname"], " ")) + "." + translit.Ru(strings.Join(r.Form["surname"], " ")) + "' -UserPrincipalName '" + translit.Ru(strings.Join(r.Form["firstname"], " ")) + "." + translit.Ru(strings.Join(r.Form["surname"], " ")) + "@ippo.mirea.ru' -MobilePhone" + strings.Join(r.Form["tel"], " ") + " -Path 'OU=PersonaledUsers,DC=ippo,DC=mirea,DC=ru")

		if err != nil {
			fmt.Fprintf(w, "Извини, но что-то пошло не так. Позови администратора и он решит твой вопрос")
			raven.CaptureErrorAndWait(err, map[string]string{"Step": "execute Powershell"})
			log.Panic(err)
		}

		log.Println(stdout)
	}
}

func main() {
	Init()
	http.HandleFunc("/", login) // setting router rule
	http.HandleFunc("/register", login)
	err := http.ListenAndServe(":9090", nil) // setting listening port
	if err != nil {
		raven.CaptureErrorAndWait(err, nil)
		log.Fatal("ListenAndServe: ", err)
	}
}
