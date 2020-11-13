package main

import (
	"fmt"
	"time"
)

type Proceso struct{
	Estado bool
	Id int
	Show bool
}

func (p *Proceso) start() {
	i := 0
	for {
		if p.Estado == false{
			return
		}
		if p.Show == true {
			fmt.Println("ID: ", p.Id, ": ", i)
		}
		time.Sleep(time.Millisecond * 500)
		i++
	}
	
}

func (p *Proceso) show() {
	p.Show = !p.Show
}


func (p *Proceso) stop() {
	p.Estado = false
}

func main() {
	var op int
	var listaprocesos []*Proceso
	auxid := 1
	flagc := 1
	flag := false

	for flagc == 1{
		fmt.Println("1.-Agregar proceso\n2.-Mostrar\n3.-Eliminar\n4.-Salir ")
		fmt.Scan(&op)

		switch op {
		case 1:
			p := &Proceso{
				Estado: true,
				Id: auxid,
				Show: flag,
			}
			go p.start()
			listaprocesos = append(listaprocesos, p)
			auxid++
			fmt.Println("Agregado")
		case 2:
			flag = !flag
			for i, value := range listaprocesos {
				value.show()
				i = i
			}
		case 3:
			var id int
			fmt.Println("ID a eliminar: ")
			fmt.Scan(&id)
			for i, value := range listaprocesos {
				if id == value.Id {
					value.stop()
					fmt.Println("Eliminado")
					listaprocesos = append(listaprocesos[:i],listaprocesos[i+1:]...)
					break
				}
			}
		case 4:
			flagc = 0
		default:
			fmt.Println("Opcion no exise")
			
		}
	}
	
}