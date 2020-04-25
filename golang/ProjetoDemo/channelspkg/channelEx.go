package channelspkg

import (
	"fmt"
	"reflect"
)

func AddToChannel(messages chan string, valor string) {
	fmt.Println("AddToChannel", reflect.TypeOf(messages), " : "+valor)
	messages <- valor
}

func CloseChannel(messages chan string) {
	close(messages)
}

func PrintChannelInLoop(messages chan string) {
	for {
		retorno, ok := <-messages
		if ok {
			fmt.Println("Mensagem do canal: " + retorno)
		} else {
			fmt.Println("Final do Canal")
			return
		}
	}
}
