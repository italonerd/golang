package main

import (
	"ProjetoDemo/channelspkg"
	"ProjetoDemo/estrutura"
	"ProjetoDemo/interfacepkg"
	"ProjetoDemo/pacote"
	"ProjetoDemo/pacote/subpacote"
	"ProjetoDemo/routinespkg"
	"fmt"
	"time"
)

func main() {
	pacote.PacotePrint(1)

	fmt.Println(subpacote.SubPacotePrint("main"))

	pacote.Looping()
	pacote.Maps()
	pacote.Pointers()

	fmt.Println("")
	fmt.Println("STRUTS")
	fmt.Println("")

	var objA estrutura.ObjetoA
	objA = objA.New(1, "Um")

	fmt.Println(objA.GetDesc())
	fmt.Println(objA.SetDesc("DOIS"))
	fmt.Println(objA.GetDesc())

	fmt.Println("")
	fmt.Println("INTERFACE")
	fmt.Println("")

	var r interfacepkg.Rect
	r = r.New(3, 4)

	var c interfacepkg.Circle
	c = c.New(5)

	fmt.Println(r)
	fmt.Println(c)

	interfacepkg.Measure(r)
	interfacepkg.Measure(c)

	fmt.Println("")
	fmt.Println("ROUTINES")
	fmt.Println("")

	go routinespkg.Say("world")

	routinespkg.Say("hello")

	fmt.Println("")
	fmt.Println("CHANNELS")
	fmt.Println("")

	messages := make(chan string)
	go channelspkg.PrintChannelInLoop(messages)

	go channelspkg.AddToChannel(messages, "Um")
	go channelspkg.AddToChannel(messages, "Dois")
	go channelspkg.AddToChannel(messages, "Tres")

	fmt.Println("Dormiu")
	time.Sleep(3 * time.Second)
	fmt.Println("Acordou")

	close(messages)
	//go channelspkg.AddToChannel(messages, "Quatro")

	fmt.Println("Dormiu")
	time.Sleep(3 * time.Second)
	fmt.Println("Acordou")

}
