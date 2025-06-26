package main

import (
	"adflowhq/workflows"
	"context"
	"flag"
	"log"
	"os"

	"go.temporal.io/sdk/client"
)

func main() {
	tenant := flag.String("tenant", "default", "tenant ID")
	lead := flag.String("lead", "lead123", "lead ID")
	flag.Parse()

	temporalAddress := os.Getenv("TEMPORAL_ADDRESS")
	if temporalAddress == "" {
		temporalAddress = "localhost:7233"
	}

	c, err := client.Dial(client.Options{
		HostPort: temporalAddress,
	})
	if err != nil {
		log.Fatalln("Unable to create client", err)
	}
	defer c.Close()

	opts := client.StartWorkflowOptions{
		ID:        *tenant + "_" + *lead + "_" + "wf",
		TaskQueue: "ADFLOWHQ_TASK_QUEUE",
	}

	we, err := c.ExecuteWorkflow(context.Background(), opts, workflows.Orchestrator, *lead)
	if err != nil {
		log.Fatalln("Unable to execute workflow", err)
	}

	log.Println("Started:", we.GetID(), we.GetRunID())
}
