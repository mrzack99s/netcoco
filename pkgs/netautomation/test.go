package netautomation

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/networklore/netrasp/pkg/netrasp"
)

func LibTest() {
	d, err := netrasp.New("x.x.x.x",
		netrasp.WithDriver("ios"),
		netrasp.WithUsernamePasswordEnableSecret("user", "password", "enable"),
	)
	if err != nil {
		log.Fatalf("unable to create client: %v", err)
	}

	ctx, cancelOpen := context.WithTimeout(context.Background(), 2000*time.Millisecond)
	defer cancelOpen()
	//Connect
	err = d.Dial(ctx)
	if err != nil {
		fmt.Printf("unable to connect: %v\n", err)

		return
	}
	defer d.Close(context.Background())

	//Enable
	err = d.Enable(context.Background())
	if err != nil {
		fmt.Printf("unable to connect: %v\n", err)
		return
	}

	ctx, cancelRun := context.WithTimeout(context.Background(), 300*time.Millisecond)
	defer cancelRun()
	output, err := d.Configure(ctx, []string{
		"int gig1/0/15", "shutdown",
	})

	fmt.Println(output)
}
