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

func Calendar_events_getHandler(cfg *config.APIConfig) func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
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
		if val, ok := args["alwaysIncludeEmail"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("alwaysIncludeEmail=%v", val))
		}
		if val, ok := args["maxAttendees"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("maxAttendees=%v", val))
		}
		if val, ok := args["timeZone"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("timeZone=%v", val))
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
		url := fmt.Sprintf("%s/calendars/%s/events/%s%s", cfg.BaseURL, calendarId, eventId, queryString)
		req, err := http.NewRequest("GET", url, nil)
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

func CreateCalendar_events_getTool(cfg *config.APIConfig) models.Tool {
	tool := mcp.NewTool("get_calendars_calendarId_events_eventId",
		mcp.WithDescription("Returns an event based on its Google Calendar ID. To retrieve an event using its iCalendar ID, call the events.list method using the iCalUID parameter."),
		mcp.WithString("calendarId", mcp.Required(), mcp.Description("Calendar identifier. To retrieve calendar IDs call the calendarList.list method. If you want to access the primary calendar of the currently logged in user, use the \"primary\" keyword.")),
		mcp.WithString("eventId", mcp.Required(), mcp.Description("Event identifier.")),
		mcp.WithBoolean("alwaysIncludeEmail", mcp.Description("Deprecated and ignored. A value will always be returned in the email field for the organizer, creator and attendees, even if no real email address is available (i.e. a generated, non-working value will be provided).")),
		mcp.WithNumber("maxAttendees", mcp.Description("The maximum number of attendees to include in the response. If there are more than the specified number of attendees, only the participant is returned. Optional.")),
		mcp.WithString("timeZone", mcp.Description("Time zone used in the response. Optional. The default is the time zone of the calendar.")),
	)

	return models.Tool{
		Definition: tool,
		Handler:    Calendar_events_getHandler(cfg),
	}
}
