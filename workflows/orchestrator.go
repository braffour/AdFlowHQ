package adflowhq.workflows

import (
    "go.temporal.io/sdk/workflow"
    "time"
    "adflowhq/activities"
    "adflowhq/agent"
)

func Orchestrator(ctx workflow.Context, leadID string) error {
    ao := workflow.ActivityOptions{
        StartToCloseTimeout: time.Minute,
        RetryPolicy: &workflow.RetryPolicy{InitialInterval: time.Second, BackoffCoefficient:2},
    }
    ctx = workflow.WithActivityOptions(ctx, ao)

    var analysis string
    err := workflow.ExecuteActivity(ctx, agent.GeminiAgent, "Analyze lead: " + leadID).Get(ctx, &analysis)
    if err!=nil { return err }
    workflow.GetLogger(ctx).Info("Agent analysis", "result", analysis)

    // Based on analysis, sync ads
    err = workflow.ExecuteActivity(ctx, activities.SyncGoogleAds, leadID).Get(ctx, nil)
    if err!=nil { return err }
    err = workflow.ExecuteActivity(ctx, activities.SyncFacebookAds, leadID).Get(ctx, nil)
    if err!=nil { return err }

    return nil
}