package tools

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/calendar-api/mcp-server/config"
	"github.com/calendar-api/mcp-server/models"
	"github.com/mark3labs/mcp-go/mcp"
)

func Calendar_events_moveHandler(cfg *config.APIConfig) func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	return func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args, ok := request.Params.Arguments.(map[string]any)
		if !ok {
			return mcp.NewToolResultError("Invalid arguments object"), nil
		}
		calendarIdVal, ok := args["calendarId"]
		if !ok {
			return mcp.NewToolResultError("Missing required path parameter: calendarId"), nil
		}
		calendarId, ok := calendarIdVal.(string)
		if !ok {
			return mcp.NewToolResultError("Invalid path parameter: calendarId"), nil
		}
		eventIdVal, ok := args["eventId"]
		if !ok {
			return mcp.NewToolResultError("Missing required path parameter: eventId"), nil
		}
		eventId, ok := eventIdVal.(string)
		if !ok {
			return mcp.NewToolResultError("Invalid path parameter: eventId"), nil
		}
		queryParams := make([]string, 0)
		if val, ok := args["destination"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("destination=%v", val))
		}
		if val, ok := args["sendNotifications"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("sendNotifications=%v", val))
		}
		if val, ok := args["sendUpdates"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("sendUpdates=%v", val))
		}
		// Handle multiple authentication parameters
		if cfg.APIKey != "" {
			queryParams = append(queryParams, fmt.Sprintf("key=%s", cfg.APIKey))
		}
		if cfg.BearerToken != "" {
			queryParams = append(queryParams, fmt.Sprintf("oauth_token=%s", cfg.BearerToken))
		}
		queryString := ""
		if len(queryParams) > 0 {
			queryString = "?" + strings.Join(queryParams, "&")
		}
		url := fmt.Sprintf("%s/calendars/%s/events/%s/move%s", cfg.BaseURL, calendarId, eventId, queryString)
		req, err := http.NewRequest("POST", url, nil)
		if err != nil {
			return mcp.NewToolResultErrorFromErr("Failed to create request", err), nil
		}
		// Set authentication based on auth type
		// Handle multiple authentication parameters
		// API key already added to query string
		// API key already added to query string
		req.Header.Set("Accept", "application/json")

		resp, err := http.DefaultClient.Do(req)
		if err != nil {
			return mcp.NewToolResultErrorFromErr("Request failed", err), nil
		}
		defer resp.Body.Close()

		body, err := io.ReadAll(resp.Body)
		if err != nil {
			return mcp.NewToolResultErrorFromErr("Failed to read response body", err), nil
		}

		if resp.StatusCode >= 400 {
			return mcp.NewToolResultError(fmt.Sprintf("API error: %s", body)), nil
		}
		// Use properly typed response
		var result models.Event
		if err := json.Unmarshal(body, &result); err != nil {
			// Fallback to raw text if unmarshaling fails
			return mcp.NewToolResultText(string(body)), nil
		}

		prettyJSON, err := json.MarshalIndent(result, "", "  ")
		if err != nil {
			return mcp.NewToolResultErrorFromErr("Failed to format JSON", err), nil
		}

		return mcp.NewToolResultText(string(prettyJSON)), nil
	}
}

func CreateCalendar_events_moveTool(cfg *config.APIConfig) models.Tool {
	tool := mcp.NewTool("post_calendars_calendarId_events_eventId_move",
		mcp.WithDescription("Moves an event to another calendar, i.e. changes an event's organizer. Note that only default events can be moved; outOfOffice, focusTime and workingLocation events cannot be moved."),
		mcp.WithString("calendarId", mcp.Required(), mcp.Description("Calendar identifier of the source calendar where the event currently is on.")),
		mcp.WithString("eventId", mcp.Required(), mcp.Description("Event identifier.")),
		mcp.WithString("destination", mcp.Required(), mcp.Description("Calendar identifier of the target calendar where the event is to be moved to.")),
		mcp.WithBoolean("sendNotifications", mcp.Description("Deprecated. Please use sendUpdates instead.\n\nWhether to send notifications about the change of the event's organizer. Note that some emails might still be sent even if you set the value to false. The default is false.")),
		mcp.WithString("sendUpdates", mcp.Description("Guests who should receive notifications about the change of the event's organizer.")),
	)

	return models.Tool{
		Definition: tool,
		Handler:    Calendar_events_moveHandler(cfg),
	}
}
