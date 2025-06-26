package main

import (
	activities "adflowhq/activities"
	agent "adflowhq/agent"
	workflows "adflowhq/workflows"
	"log"
	"os"

	"go.temporal.io/sdk/client"
	"go.temporal.io/sdk/worker"
)

func main() {
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

	w := worker.New(c, "ADFLOWHQ_TASK_QUEUE", worker.Options{})
	w.RegisterWorkflow(workflows.Orchestrator)
	w.RegisterActivity(agent.GeminiAgent)
	w.RegisterActivity(activities.SyncGoogleAds)
	w.RegisterActivity(activities.SyncFacebookAds)

	err = w.Run(worker.InterruptCh())
	if err != nil {
		log.Fatalln("Unable to start worker", err)
	}
}
