package workflows

import (
	activities "adflowhq/activities"
	agent "adflowhq/agent"
	"time"

	"go.temporal.io/sdk/workflow"
)

func Orchestrator(ctx workflow.Context, leadID string) error {
    ao := workflow.ActivityOptions{
        StartToCloseTimeout: time.Minute * 5,
        RetryPolicy: &workflow.RetryPolicy{
            InitialInterval: time.Second,
            BackoffCoefficient: 2.0,
            MaximumInterval: time.Minute,
            MaximumAttempts: 3,
        },
    }
    ctx = workflow.WithActivityOptions(ctx, ao)

    workflow.GetLogger(ctx).Info("Starting orchestrator workflow", "leadID", leadID)

    var analysis string
    err := workflow.ExecuteActivity(ctx, agent.GeminiAgent, "Analyze lead: " + leadID).Get(ctx, &analysis)
    if err != nil {
        workflow.GetLogger(ctx).Error("Failed to analyze lead with AI agent", "error", err)
        return err
    }
    workflow.GetLogger(ctx).Info("Agent analysis completed", "result", analysis)

    // Based on analysis, sync ads platforms
    err = workflow.ExecuteActivity(ctx, activities.SyncGoogleAds, leadID).Get(ctx, nil)
    if err != nil {
        workflow.GetLogger(ctx).Error("Failed to sync Google Ads", "error", err)
        return err
    }

    err = workflow.ExecuteActivity(ctx, activities.SyncFacebookAds, leadID).Get(ctx, nil)
    if err != nil {
        workflow.GetLogger(ctx).Error("Failed to sync Facebook Ads", "error", err)
        return err
    }

    workflow.GetLogger(ctx).Info("Orchestrator workflow completed successfully", "leadID", leadID)
    return nil
}