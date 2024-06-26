package main

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		err := setupAWS(ctx)
		if err != nil {
			return err
		}

		err = setupOCI(ctx)
		if err != nil {
			return err
		}

		return nil
	})
}
