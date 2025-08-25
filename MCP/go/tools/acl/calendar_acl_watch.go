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

func Calendar_acl_watchHandler(cfg *config.APIConfig) func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
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
		if val, ok := args["maxResults"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("maxResults=%v", val))
		}
		if val, ok := args["pageToken"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("pageToken=%v", val))
		}
		if val, ok := args["showDeleted"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("showDeleted=%v", val))
		}
		if val, ok := args["syncToken"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("syncToken=%v", val))
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
		var requestBody models.Channel
		
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
		url := fmt.Sprintf("%s/calendars/%s/acl/watch%s", cfg.BaseURL, calendarId, queryString)
		req, err := http.NewRequest("POST", url, bytes.NewBuffer(bodyBytes))
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
		var result models.Channel
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

func CreateCalendar_acl_watchTool(cfg *config.APIConfig) models.Tool {
	tool := mcp.NewTool("post_calendars_calendarId_acl_watch",
		mcp.WithDescription("Watch for changes to ACL resources."),
		mcp.WithString("calendarId", mcp.Required(), mcp.Description("Calendar identifier. To retrieve calendar IDs call the calendarList.list method. If you want to access the primary calendar of the currently logged in user, use the \"primary\" keyword.")),
		mcp.WithNumber("maxResults", mcp.Description("Maximum number of entries returned on one result page. By default the value is 100 entries. The page size can never be larger than 250 entries. Optional.")),
		mcp.WithString("pageToken", mcp.Description("Token specifying which result page to return. Optional.")),
		mcp.WithBoolean("showDeleted", mcp.Description("Whether to include deleted ACLs in the result. Deleted ACLs are represented by role equal to \"none\". Deleted ACLs will always be included if syncToken is provided. Optional. The default is False.")),
		mcp.WithString("syncToken", mcp.Description("Token obtained from the nextSyncToken field returned on the last page of results from the previous list request. It makes the result of this list request contain only entries that have changed since then. All entries deleted since the previous list request will always be in the result set and it is not allowed to set showDeleted to False.\nIf the syncToken expires, the server will respond with a 410 GONE response code and the client should clear its storage and perform a full synchronization without any syncToken.\nLearn more about incremental synchronization.\nOptional. The default is to return all entries.")),
		mcp.WithString("token", mcp.Description("Input parameter: An arbitrary string delivered to the target address with each notification delivered over this channel. Optional.")),
		mcp.WithString("id", mcp.Description("Input parameter: A UUID or similar unique string that identifies this channel.")),
		mcp.WithString("kind", mcp.Description("Input parameter: Identifies this as a notification channel used to watch for changes to a resource, which is \"api#channel\".")),
		mcp.WithString("type", mcp.Description("Input parameter: The type of delivery mechanism used for this channel. Valid values are \"web_hook\" (or \"webhook\"). Both values refer to a channel where Http requests are used to deliver messages.")),
		mcp.WithString("address", mcp.Description("Input parameter: The address where notifications are delivered for this channel.")),
		mcp.WithBoolean("payload", mcp.Description("Input parameter: A Boolean value to indicate whether payload is wanted. Optional.")),
		mcp.WithString("resourceId", mcp.Description("Input parameter: An opaque ID that identifies the resource being watched on this channel. Stable across different API versions.")),
		mcp.WithString("expiration", mcp.Description("Input parameter: Date and time of notification channel expiration, expressed as a Unix timestamp, in milliseconds. Optional.")),
		mcp.WithObject("params", mcp.Description("Input parameter: Additional parameters controlling delivery channel behavior. Optional.")),
		mcp.WithString("resourceUri", mcp.Description("Input parameter: A version-specific identifier for the watched resource.")),
	)

	return models.Tool{
		Definition: tool,
		Handler:    Calendar_acl_watchHandler(cfg),
	}
}
