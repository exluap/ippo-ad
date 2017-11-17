package main

import (
    "fmt"
    "html/template"
    "log"
    "net/http"
    "strings"
    "github.com/exluap/translit"

    ps "github.com/gorillalabs/go-powershell"
    "github.com/gorillalabs/go-powershell/backend"
)

func sayhelloName(w http.ResponseWriter, r *http.Request) {
    r.ParseForm() //Parse url parameters passed, then parse the response packet for the POST body (request body)
    // attention: If you do not call ParseForm method, the following data can not be obtained form
    fmt.Println(r.Form) // print information on server side.
    fmt.Println("path", r.URL.Path)
    fmt.Println("scheme", r.URL.Scheme)
    fmt.Println(r.Form["url_long"])
    for k, v := range r.Form {
        fmt.Println("key:", k)
        fmt.Println("val:", strings.Join(v, ""))
    }
    fmt.Fprintf(w, "Hello astaxie!") // write data to response
}

func login(w http.ResponseWriter, r *http.Request) {
    fmt.Println("method:", r.Method) //get request method
    if r.Method == "GET" {
        t, _ := template.ParseFiles("index.gtpl")
        t.Execute(w, nil)
    } else {
        r.ParseForm()
        // logic part of log in
        fmt.Println("fam:", r.Form["surname"])
        fmt.Println("name:", r.Form["firstname"])
        fmt.Println("otch:", r.Form["endname"])
        fmt.Println("group:", r.Form["univergroup"])
        fmt.Println("kaf:", r.Form["univerkafedra"])
        fmt.Println("email:", r.Form["email"])
        fmt.Println("pass:", r.Form["password"])
       // fmt.Println(translit.Ru(strings.Join(r.Form["surname"], " ")))
        fmt.Fprintf(w, "Твой логин: " + translit.Ru(strings.Join(r.Form["firstname"], " ")) + "." + translit.Ru(strings.Join(r.Form["surname"], " ")))
        fmt.Println("New-ADUser -Name '" + translit.Ru(strings.Join(r.Form["firstname"], " ")) + "." + translit.Ru(strings.Join(r.Form["surname"], " ")) + "' -AccountPassword(ConvertTo-SecureString '" + strings.Join(r.Form["password"], " ") + "' -AsPlainText -Force) -ChangePasswordAtLogon 1 -Company '" + strings.Join(r.Form["univerkafedra"], " ") + "' -Department '" + strings.Join(r.Form["univergroup"], " ") + "' -DisplayName '" + strings.Join(r.Form["surname"], " ") + " " + strings.Join(r.Form["firstname"], " ") +  " "+ strings.Join(r.Form["endname"]," ") + "' -EmailAddress '" + strings.Join(r.Form["email"], " ") + "' -Enabled 1 -GivenName '" + strings.Join(r.Form["firstname"], " ") + " " + strings.Join(r.Form["endname"]," ") + "' -SurName '" + strings.Join(r.Form["surname"]," ") + " " + "' -HomeDirectory '\\\\POSEIDON\\movedUsers\\%username%' -SamAccountName '" + translit.Ru(strings.Join(r.Form["firstname"], " ")) + "." + translit.Ru(strings.Join(r.Form["surname"], " ")) + "' -UserPrincipalName '" + translit.Ru(strings.Join(r.Form["firstname"], " ")) + "." + translit.Ru(strings.Join(r.Form["surname"], " ")) + "'@ippo.mirea.ru -Path 'ou=Пользователи,DC=mirea,DC=ippo,DC=ru'")
        //fmt.Println("ADD-ADGroupMember 'students' –members '"+translit.Ru(strings.Join(r.Form["firstname"], " ")) + "." + translit.Ru(strings.Join(r.Form["surname"], " "))+"'")
        
        // choose a backend
        back := &backend.Local{}

        // start a local powershell process
        shell, err := ps.New(back)
        if err != nil {
            panic(err)
        }
        defer shell.Exit()

        // ... and interact with it
        stdout, _, err := shell.Execute("New-ADUser -Name '" + translit.Ru(strings.Join(r.Form["firstname"], " ")) + "." + translit.Ru(strings.Join(r.Form["surname"], " ")) + "' -AccountPassword(ConvertTo-SecureString '" + strings.Join(r.Form["password"], " ") + "' -AsPlainText -Force) -ChangePasswordAtLogon 1 -Company '" + strings.Join(r.Form["univerkafedra"], " ") + "' -Department '" + strings.Join(r.Form["univergroup"], " ") + "' -DisplayName '" + strings.Join(r.Form["surname"], " ") + " " + strings.Join(r.Form["firstname"], " ") +  " "+ strings.Join(r.Form["endname"]," ") + "' -EmailAddress '" + strings.Join(r.Form["email"], " ") + "' -Enabled 1 -GivenName '" + strings.Join(r.Form["firstname"], " ") + " " + strings.Join(r.Form["endname"]," ") + "' -SurName '" + strings.Join(r.Form["surname"]," ") + " " + "' -HomeDirectory '\\\\POSEIDON\\movedUsers\\%username%' -SamAccountName '" + translit.Ru(strings.Join(r.Form["firstname"], " ")) + "." + translit.Ru(strings.Join(r.Form["surname"], " ")) + "' -UserPrincipalName '" + translit.Ru(strings.Join(r.Form["firstname"], " ")) + "." + translit.Ru(strings.Join(r.Form["surname"], " ")) + "@ippo.mirea.ru'")

        if err != nil {
            panic(err)
        }

        
        fmt.Println(stdout) 
    }
}

func main() {
    http.HandleFunc("/", sayhelloName) // setting router rule
    http.HandleFunc("/register", login)
    err := http.ListenAndServe(":9090", nil) // setting listening port
    if err != nil {
        log.Fatal("ListenAndServe: ", err)
    }
}