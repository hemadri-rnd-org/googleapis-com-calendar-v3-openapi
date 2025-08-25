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

func Calendar_events_listHandler(cfg *config.APIConfig) func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
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
		queryParams := make([]string, 0)
		if val, ok := args["alwaysIncludeEmail"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("alwaysIncludeEmail=%v", val))
		}
		if val, ok := args["eventTypes"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("eventTypes=%v", val))
		}
		if val, ok := args["iCalUID"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("iCalUID=%v", val))
		}
		if val, ok := args["maxAttendees"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("maxAttendees=%v", val))
		}
		if val, ok := args["maxResults"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("maxResults=%v", val))
		}
		if val, ok := args["orderBy"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("orderBy=%v", val))
		}
		if val, ok := args["pageToken"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("pageToken=%v", val))
		}
		if val, ok := args["privateExtendedProperty"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("privateExtendedProperty=%v", val))
		}
		if val, ok := args["q"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("q=%v", val))
		}
		if val, ok := args["sharedExtendedProperty"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("sharedExtendedProperty=%v", val))
		}
		if val, ok := args["showDeleted"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("showDeleted=%v", val))
		}
		if val, ok := args["showHiddenInvitations"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("showHiddenInvitations=%v", val))
		}
		if val, ok := args["singleEvents"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("singleEvents=%v", val))
		}
		if val, ok := args["syncToken"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("syncToken=%v", val))
		}
		if val, ok := args["timeMax"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("timeMax=%v", val))
		}
		if val, ok := args["timeMin"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("timeMin=%v", val))
		}
		if val, ok := args["timeZone"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("timeZone=%v", val))
		}
		if val, ok := args["updatedMin"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("updatedMin=%v", val))
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
		url := fmt.Sprintf("%s/calendars/%s/events%s", cfg.BaseURL, calendarId, queryString)
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
		var result models.Events
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

func CreateCalendar_events_listTool(cfg *config.APIConfig) models.Tool {
	tool := mcp.NewTool("get_calendars_calendarId_events",
		mcp.WithDescription("Returns events on the specified calendar."),
		mcp.WithString("calendarId", mcp.Required(), mcp.Description("Calendar identifier. To retrieve calendar IDs call the calendarList.list method. If you want to access the primary calendar of the currently logged in user, use the \"primary\" keyword.")),
		mcp.WithBoolean("alwaysIncludeEmail", mcp.Description("Deprecated and ignored.")),
		mcp.WithArray("eventTypes", mcp.Description("Event types to return. Optional. This parameter can be repeated multiple times to return events of different types. The default is [\"default\", \"focusTime\", \"outOfOffice\"].")),
		mcp.WithString("iCalUID", mcp.Description("Specifies an event ID in the iCalendar format to be provided in the response. Optional. Use this if you want to search for an event by its iCalendar ID.")),
		mcp.WithNumber("maxAttendees", mcp.Description("The maximum number of attendees to include in the response. If there are more than the specified number of attendees, only the participant is returned. Optional.")),
		mcp.WithNumber("maxResults", mcp.Description("Maximum number of events returned on one result page. The number of events in the resulting page may be less than this value, or none at all, even if there are more events matching the query. Incomplete pages can be detected by a non-empty nextPageToken field in the response. By default the value is 250 events. The page size can never be larger than 2500 events. Optional.")),
		mcp.WithString("orderBy", mcp.Description("The order of the events returned in the result. Optional. The default is an unspecified, stable order.")),
		mcp.WithString("pageToken", mcp.Description("Token specifying which result page to return. Optional.")),
		mcp.WithArray("privateExtendedProperty", mcp.Description("Extended properties constraint specified as propertyName=value. Matches only private properties. This parameter might be repeated multiple times to return events that match all given constraints.")),
		mcp.WithString("q", mcp.Description("Free text search terms to find events that match these terms in the following fields:\n\n- summary \n- description \n- location \n- attendee's displayName \n- attendee's email \n- organizer's displayName \n- organizer's email \n- workingLocationProperties.officeLocation.buildingId \n- workingLocationProperties.officeLocation.deskId \n- workingLocationProperties.officeLocation.label \n- workingLocationProperties.customLocation.label \nThese search terms also match predefined keywords against all display title translations of working location, out-of-office, and focus-time events. For example, searching for \"Office\" or \"Bureau\" returns working location events of type officeLocation, whereas searching for \"Out of office\" or \"Abwesend\" returns out-of-office events. Optional.")),
		mcp.WithArray("sharedExtendedProperty", mcp.Description("Extended properties constraint specified as propertyName=value. Matches only shared properties. This parameter might be repeated multiple times to return events that match all given constraints.")),
		mcp.WithBoolean("showDeleted", mcp.Description("Whether to include deleted events (with status equals \"cancelled\") in the result. Cancelled instances of recurring events (but not the underlying recurring event) will still be included if showDeleted and singleEvents are both False. If showDeleted and singleEvents are both True, only single instances of deleted events (but not the underlying recurring events) are returned. Optional. The default is False.")),
		mcp.WithBoolean("showHiddenInvitations", mcp.Description("Whether to include hidden invitations in the result. Optional. The default is False.")),
		mcp.WithBoolean("singleEvents", mcp.Description("Whether to expand recurring events into instances and only return single one-off events and instances of recurring events, but not the underlying recurring events themselves. Optional. The default is False.")),
		mcp.WithString("syncToken", mcp.Description("Token obtained from the nextSyncToken field returned on the last page of results from the previous list request. It makes the result of this list request contain only entries that have changed since then. All events deleted since the previous list request will always be in the result set and it is not allowed to set showDeleted to False.\nThere are several query parameters that cannot be specified together with nextSyncToken to ensure consistency of the client state.\n\nThese are: \n- iCalUID \n- orderBy \n- privateExtendedProperty \n- q \n- sharedExtendedProperty \n- timeMin \n- timeMax \n- updatedMin All other query parameters should be the same as for the initial synchronization to avoid undefined behavior. If the syncToken expires, the server will respond with a 410 GONE response code and the client should clear its storage and perform a full synchronization without any syncToken.\nLearn more about incremental synchronization.\nOptional. The default is to return all entries.")),
		mcp.WithString("timeMax", mcp.Description("Upper bound (exclusive) for an event's start time to filter by. Optional. The default is not to filter by start time. Must be an RFC3339 timestamp with mandatory time zone offset, for example, 2011-06-03T10:00:00-07:00, 2011-06-03T10:00:00Z. Milliseconds may be provided but are ignored. If timeMin is set, timeMax must be greater than timeMin.")),
		mcp.WithString("timeMin", mcp.Description("Lower bound (exclusive) for an event's end time to filter by. Optional. The default is not to filter by end time. Must be an RFC3339 timestamp with mandatory time zone offset, for example, 2011-06-03T10:00:00-07:00, 2011-06-03T10:00:00Z. Milliseconds may be provided but are ignored. If timeMax is set, timeMin must be smaller than timeMax.")),
		mcp.WithString("timeZone", mcp.Description("Time zone used in the response. Optional. The default is the time zone of the calendar.")),
		mcp.WithString("updatedMin", mcp.Description("Lower bound for an event's last modification time (as a RFC3339 timestamp) to filter by. When specified, entries deleted since this time will always be included regardless of showDeleted. Optional. The default is not to filter by last modification time.")),
	)

	return models.Tool{
		Definition: tool,
		Handler:    Calendar_events_listHandler(cfg),
	}
}
