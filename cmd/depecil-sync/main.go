package main

import (
	"net/http"
	"fmt"

	. "github.com/julvo/htmlgo"
    a "github.com/julvo/htmlgo/attributes"
)

type Register struct {
	Name string
	Number_of_registers int
}

func handle(w http.ResponseWriter, r *http.Request) {

	registers := []Register {
		{"Teste", 10},
	}

	rows := HTML("")

	for _, register := range registers {
		rows += Tr_( 
			Td_(Text(register.Name)),
			Td_(Text(register.Number_of_registers)))
	}

	thead := Thead(
		Attr(), 
		Tr(Attr(), 
			Th_(Text("Nome do banco de dadod")),
			Th_(Text("Números de registros"))))

			
	tbody := Tbody_(rows)


	page :=
        Html5_(
            Head_(
				Title_(Text("Depecil Sync")),
				Meta(Attr(a.Charset_("utf-8"))),
				Meta(Attr(a.Name_("viewport"), a.Content_("width=device-width"), a.InitialScale_("1"))),
				Link(Attr(a.Rel_("stylesheet"), a.Href_("/style.css"))),
			),
            Body_(
				Script_(JavaScript("", "setInterval(function() { window.navigation.reload() }, 1000 * 60);")),
				H1_(Text("Depecil Sync")),
				Table_(thead,tbody)))

	fmt.Fprint(w, page)
}

func main(){
	directory := "./teste";

	http.HandleFunc("/teste", handle);


	http.Handle("/", http.FileServer(http.Dir(directory)));

	address := "localhost";
	port := 8080
	fullAddress := fmt.Sprintf("%s:%d", address, port)

	fmt.Printf("Servidor iniciado em http://%s\n", fullAddress)

	err := http.ListenAndServe(fullAddress, nil)

	if err != nil {
		fmt.Println("Erro ao iniciar o servidor:", err)
	}

}