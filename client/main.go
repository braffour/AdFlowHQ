package adflowhq.client

import (
    "context"
    "flag"
    "log"
    "go.temporal.io/sdk/client"
    "adflowhq/workflows"
)

func main() {
    tenant := flag.String("tenant", "default", "tenant ID")
    lead := flag.String("lead", "lead123", "lead ID")
    flag.Parse()

    c, err := client.Dial(client.Options{})
    if err!=nil { log.Fatalln(err) }
    defer c.Close()

    opts := client.StartWorkflowOptions{
        ID: *tenant + "_" + *lead + "_" + "wf",
        TaskQueue: "ADFLOWHQ_TASK_QUEUE",
    }

    we, err := c.ExecuteWorkflow(context.Background(), opts, workflows.Orchestrator, *lead)
    if err!=nil { log.Fatalln(err) }

    log.Println("Started:", we.GetID(), we.GetRunID())
}