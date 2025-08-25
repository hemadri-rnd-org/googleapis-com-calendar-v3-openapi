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

func Calendar_acl_updateHandler(cfg *config.APIConfig) func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
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
		ruleIdVal, ok := args["ruleId"]
		if !ok {
			return mcp.NewToolResultError("Missing required path parameter: ruleId"), nil
		}
		ruleId, ok := ruleIdVal.(string)
		if !ok {
			return mcp.NewToolResultError("Invalid path parameter: ruleId"), nil
		}
		queryParams := make([]string, 0)
		if val, ok := args["sendNotifications"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("sendNotifications=%v", val))
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
		var requestBody models.AclRule
		
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
		url := fmt.Sprintf("%s/calendars/%s/acl/%s%s", cfg.BaseURL, calendarId, ruleId, queryString)
		req, err := http.NewRequest("PUT", url, bytes.NewBuffer(bodyBytes))
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
		var result models.AclRule
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

func CreateCalendar_acl_updateTool(cfg *config.APIConfig) models.Tool {
	tool := mcp.NewTool("put_calendars_calendarId_acl_ruleId",
		mcp.WithDescription("Updates an access control rule."),
		mcp.WithString("calendarId", mcp.Required(), mcp.Description("Calendar identifier. To retrieve calendar IDs call the calendarList.list method. If you want to access the primary calendar of the currently logged in user, use the \"primary\" keyword.")),
		mcp.WithString("ruleId", mcp.Required(), mcp.Description("ACL rule identifier.")),
		mcp.WithBoolean("sendNotifications", mcp.Description("Whether to send notifications about the calendar sharing change. Note that there are no notifications on access removal. Optional. The default is True.")),
		mcp.WithString("id", mcp.Description("Input parameter: Identifier of the Access Control List (ACL) rule. See Sharing calendars.")),
		mcp.WithString("kind", mcp.Description("Input parameter: Type of the resource (\"calendar#aclRule\").")),
		mcp.WithString("role", mcp.Description("Input parameter: The role assigned to the scope. Possible values are:  \n- \"none\" - Provides no access. \n- \"freeBusyReader\" - Provides read access to free/busy information. \n- \"reader\" - Provides read access to the calendar. Private events will appear to users with reader access, but event details will be hidden. \n- \"writer\" - Provides read and write access to the calendar. Private events will appear to users with writer access, and event details will be visible. \n- \"owner\" - Provides ownership of the calendar. This role has all of the permissions of the writer role with the additional ability to see and manipulate ACLs.")),
		mcp.WithObject("scope", mcp.Description("Input parameter: The extent to which calendar access is granted by this ACL rule.")),
		mcp.WithString("etag", mcp.Description("Input parameter: ETag of the resource.")),
	)

	return models.Tool{
		Definition: tool,
		Handler:    Calendar_acl_updateHandler(cfg),
	}
}
