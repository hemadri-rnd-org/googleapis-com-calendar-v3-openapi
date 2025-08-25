package tools

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
	"bytes"

	"github.com/calendar-api/mcp-server/config"
	"github.com/calendar-api/mcp-server/models"
	"github.com/mark3labs/mcp-go/mcp"
)

func Calendar_calendarlist_patchHandler(cfg *config.APIConfig) func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
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
		if val, ok := args["colorRgbFormat"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("colorRgbFormat=%v", val))
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
		// Create properly typed request body using the generated schema
		var requestBody models.CalendarListEntry
		
		// Optimized: Single marshal/unmarshal with JSON tags handling field mapping
		if argsJSON, err := json.Marshal(args); err == nil {
			if err := json.Unmarshal(argsJSON, &requestBody); err != nil {
				return mcp.NewToolResultError(fmt.Sprintf("Failed to convert arguments to request type: %v", err)), nil
			}
		} else {
			return mcp.NewToolResultError(fmt.Sprintf("Failed to marshal arguments: %v", err)), nil
		}
		
		bodyBytes, err := json.Marshal(requestBody)
		if err != nil {
			return mcp.NewToolResultErrorFromErr("Failed to encode request body", err), nil
		}
		url := fmt.Sprintf("%s/users/me/calendarList/%s%s", cfg.BaseURL, calendarId, queryString)
		req, err := http.NewRequest("PATCH", url, bytes.NewBuffer(bodyBytes))
		req.Header.Set("Content-Type", "application/json")
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
		var result models.CalendarListEntry
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

func CreateCalendar_calendarlist_patchTool(cfg *config.APIConfig) models.Tool {
	tool := mcp.NewTool("patch_users_me_calendarList_calendarId",
		mcp.WithDescription("Updates an existing calendar on the user's calendar list. This method supports patch semantics."),
		mcp.WithString("calendarId", mcp.Required(), mcp.Description("Calendar identifier. To retrieve calendar IDs call the calendarList.list method. If you want to access the primary calendar of the currently logged in user, use the \"primary\" keyword.")),
		mcp.WithBoolean("colorRgbFormat", mcp.Description("Whether to use the foregroundColor and backgroundColor fields to write the calendar colors (RGB). If this feature is used, the index-based colorId field will be set to the best matching option automatically. Optional. The default is False.")),
		mcp.WithString("summaryOverride", mcp.Description("Input parameter: The summary that the authenticated user has set for this calendar. Optional.")),
		mcp.WithString("accessRole", mcp.Description("Input parameter: The effective access role that the authenticated user has on the calendar. Read-only. Possible values are:  \n- \"freeBusyReader\" - Provides read access to free/busy information. \n- \"reader\" - Provides read access to the calendar. Private events will appear to users with reader access, but event details will be hidden. \n- \"writer\" - Provides read and write access to the calendar. Private events will appear to users with writer access, and event details will be visible. \n- \"owner\" - Provides ownership of the calendar. This role has all of the permissions of the writer role with the additional ability to see and manipulate ACLs.")),
		mcp.WithString("timeZone", mcp.Description("Input parameter: The time zone of the calendar. Optional. Read-only.")),
		mcp.WithString("foregroundColor", mcp.Description("Input parameter: The foreground color of the calendar in the hexadecimal format \"#ffffff\". This property supersedes the index-based colorId property. To set or change this property, you need to specify colorRgbFormat=true in the parameters of the insert, update and patch methods. Optional.")),
		mcp.WithString("location", mcp.Description("Input parameter: Geographic location of the calendar as free-form text. Optional. Read-only.")),
		mcp.WithBoolean("deleted", mcp.Description("Input parameter: Whether this calendar list entry has been deleted from the calendar list. Read-only. Optional. The default is False.")),
		mcp.WithBoolean("primary", mcp.Description("Input parameter: Whether the calendar is the primary calendar of the authenticated user. Read-only. Optional. The default is False.")),
		mcp.WithString("colorId", mcp.Description("Input parameter: The color of the calendar. This is an ID referring to an entry in the calendar section of the colors definition (see the colors endpoint). This property is superseded by the backgroundColor and foregroundColor properties and can be ignored when using these properties. Optional.")),
		mcp.WithString("kind", mcp.Description("Input parameter: Type of the resource (\"calendar#calendarListEntry\").")),
		mcp.WithString("etag", mcp.Description("Input parameter: ETag of the resource.")),
		mcp.WithArray("defaultReminders", mcp.Description("Input parameter: The default reminders that the authenticated user has for this calendar.")),
		mcp.WithObject("notificationSettings", mcp.Description("Input parameter: The notifications that the authenticated user is receiving for this calendar.")),
		mcp.WithString("backgroundColor", mcp.Description("Input parameter: The main color of the calendar in the hexadecimal format \"#0088aa\". This property supersedes the index-based colorId property. To set or change this property, you need to specify colorRgbFormat=true in the parameters of the insert, update and patch methods. Optional.")),
		mcp.WithString("description", mcp.Description("Input parameter: Description of the calendar. Optional. Read-only.")),
		mcp.WithBoolean("selected", mcp.Description("Input parameter: Whether the calendar content shows up in the calendar UI. Optional. The default is False.")),
		mcp.WithString("summary", mcp.Description("Input parameter: Title of the calendar. Read-only.")),
		mcp.WithString("id", mcp.Description("Input parameter: Identifier of the calendar.")),
		mcp.WithObject("conferenceProperties", mcp.Description("")),
		mcp.WithBoolean("hidden", mcp.Description("Input parameter: Whether the calendar has been hidden from the list. Optional. The attribute is only returned when the calendar is hidden, in which case the value is true.")),
	)

	return models.Tool{
		Definition: tool,
		Handler:    Calendar_calendarlist_patchHandler(cfg),
	}
}
