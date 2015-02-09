package main

import (
	"fmt"
	"github.com/tinylib/msgp/msgp"
	"github.com/tinylib/synapse"
	"net"
	"os"
)

type echo struct{}

func (e echo) ServeCall(req synapse.Request, res synapse.ResponseWriter) {
	var r msgp.Raw
	err := req.Decode(&r)
	if err != nil {
		res.Error(synapse.BadRequest)
		return
	}
	res.Send(&r)
}

func main() {
	l, err := net.Listen("tcp", ":7070")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	go synapse.Serve(l, echo{})

	// START OMIT
	cl, err := synapse.Dial("tcp", ":7070", 200)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	err = cl.Call("echo", synapse.String("hello, world!"), synapse.JSPipe(os.Stdout))
	// END OMIT
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	cl.Close()
	l.Close()
}
