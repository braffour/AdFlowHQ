package adflowhq.worker

import (
    "log"
    "adflowhq/activities"
    "adflowhq/agent"
    "adflowhq/workflows"
    "go.temporal.io/sdk/client"
    "go.temporal.io/sdk/worker"
)

func main() {
    c, err := client.Dial(client.Options{})
    if err!=nil { log.Fatalln(err) }
    defer c.Close()

    w := worker.New(c, "ADFLOWHQ_TASK_QUEUE", worker.Options{})
    w.RegisterWorkflow(workflows.Orchestrator)
    w.RegisterActivity(agent.GeminiAgent)
    w.RegisterActivity(activities.SyncGoogleAds)
    w.RegisterActivity(activities.SyncFacebookAds)

    err = w.Run(worker.InterruptCh())
    if err!=nil { log.Fatalln(err) }
}