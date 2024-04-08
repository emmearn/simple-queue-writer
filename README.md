# SIMPLE QUEUE WRITER

Consigliato l'ambiente Linux all’interno del vostro sistema Windows, dovreste trovarlo preinstallato come “Ubuntu on Windows”

## Installazione del compiler

1. *sudo apt update && apt upgrade -y*
2. *curl -LO https://go.dev/dl/go1.19.3.linux-amd64.tar.gz*
3. *sudo tar -C /usr/local -xzf go1.19.3.linux-amd64.tar.gz*
4. *sudo nano ~/.bashrc* e appendere alla fine del file

    ```bash
        # GoLang variables
        GOPRIVATE=bitbucket.org/m_arnone/*
        GOROOT=/usr/local/go
        GOPATH=$HOME/go
        PATH=$GOPATH/bin:$GOROOT/bin:$PATH
    ```

5. *source ~/.bashrc*
6. *go version* dovrebbe restituire appunto la versione 19.3
7. *nano hello.go*
    e appendere il sorgente:

    ```go
        // Hello Word in Go by Vivek Gite
        package main
        
        // Import OS and fmt packages
        import ( 
            "fmt" 
            "os" 
        )
        
        // Let us start
        func main() {
            fmt.Println("Hello, world!")  // Print simple text on screen
            fmt.Println(os.Getenv("USER"), ", Let's be friends!") // Read Linux $USER environment variable 
        }
    ```

8. *go run hello.go*

Se la shell vi saluta, avete configurato correttamente il vostro ambiente :)

9. *sudo apt install make*
10. *sudo apt install gcc*

## Comandi

La maggior parte dei comandi utili sono già disponibili nel Makefile, utilizzabile con *make + comando* per esempio *make build*.

* *go* restituisce l'elenco dei comandi disponibili
* Documentazione https://pkg.go.dev/cmd

## IDE

Esiste un IDE JetBrains Goland molto valido, ma di cui non esiste la community version come per IntelliJ.
Se in possesso della licenza JetBrains conviene usare questo.
In caso contrario si può usare Visual Studio Code con relativa estensione per linguaggio GoLang.

## Creazione nuovo modulo

1. Creare la cartella all'interno del progetto
2. spostarsi all'interno della cartella
3. Inizializzare il nuovo modulo (per esempio il modulo "commission") con:

    ```bash
    go mod init bitbucket.org/m_arnone/gogeco/commission
    ```

## Installazione dipendenze
1. ```bash
    go install github.com/golang/mock/mockgen@v1.6.0
    ```
2. ```bash
    go get golang.org/x/tools/cmd/cover
    ```

## Compilazione ad avvio

Una volta sviluppato il modulo sarà suffificente lanciare:

1. *make build*
2. *./bin/main*

## Convenzioni

1. i nomi dei file in golang si scrivono tutti in minuscolo separati da “_”
2. *./cmd* per convenzione è la cartella che contiene i file di build del package *main*
3. *./internal* è una cartella speciale riconosciuta dal *go tol* che prevede che ogni package al suo interno non sia esportato
4. *./pkg* per convenzione è la cartella che contiene il codice importabile da progetti terzi
5. le variabili, le funzioni e le costanti in camel case:
    * *WriteToDB* esportata/pubblica, visibile al di fuori del package
    * *writeToDB* sarà visibile solo nel package corrente
6. rendere private la maggior parte delle logiche, lasciando esposte solo quelle che devono essere importate da altri package
7. i file di test se hanno il suffisso *_test.go* vengono eseguiti solo dal comando “*go test*"
8. le variabili di nomi brevi dove il contesto è lampante
    * *user* diventa *u*
    * *userID* diventa *uid*
9. booleani nel nome dovrebbero contenere prefissi come *Has*, *Is*, *Can* o *Allow*
10. *i*, *j*, *k* usati come indici
11. i nomi dei package devono essere brevi ed essere significative abbreviazioni
    * *strconv* (string conversion)
    * *syscall* (system call)
    * *fmt* (formatted I/O)
D'altra parte se l'abbreviazione causa ambiguità, è da evitare.
