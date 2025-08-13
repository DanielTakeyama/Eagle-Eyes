package core

import (
	"fmt"
	"time"
)

type Wordlist struct {
	Wordpress []string
	Default   []string
}

func (w *Wordlist) ScanWordpress(url string) error{
	if len(w.Wordpress) > 0 {
		for _, word := range w.Wordpress {
			//Colocar a lógica aqui, esse print e sleep foi só para validar uma coisa
			fmt.Printf("Scaneando %v/%v\n", url, word)
			time.Sleep(1 * time.Second)
		}
		fmt.Println("Varredura WORDPRESS comleta.")
		return nil
	}
	return fmt.Errorf("erro: lista wordpress vazia")
}
