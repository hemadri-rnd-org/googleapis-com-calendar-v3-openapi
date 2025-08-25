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

func Calendar_events_importHandler(cfg *config.APIConfig) func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
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
		if val, ok := args["conferenceDataVersion"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("conferenceDataVersion=%v", val))
		}
		if val, ok := args["supportsAttachments"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("supportsAttachments=%v", val))
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
		var requestBody models.Event
		
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
		url := fmt.Sprintf("%s/calendars/%s/events/import%s", cfg.BaseURL, calendarId, queryString)
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

func CreateCalendar_events_importTool(cfg *config.APIConfig) models.Tool {
	tool := mcp.NewTool("post_calendars_calendarId_events_import",
		mcp.WithDescription("Imports an event. This operation is used to add a private copy of an existing event to a calendar."),
		mcp.WithString("calendarId", mcp.Required(), mcp.Description("Calendar identifier. To retrieve calendar IDs call the calendarList.list method. If you want to access the primary calendar of the currently logged in user, use the \"primary\" keyword.")),
		mcp.WithNumber("conferenceDataVersion", mcp.Description("Version number of conference data supported by the API client. Version 0 assumes no conference data support and ignores conference data in the event's body. Version 1 enables support for copying of ConferenceData as well as for creating new conferences using the createRequest field of conferenceData. The default is 0.")),
		mcp.WithBoolean("supportsAttachments", mcp.Description("Whether API client performing operation supports event attachments. Optional. The default is False.")),
		mcp.WithBoolean("endTimeUnspecified", mcp.Description("Input parameter: Whether the end time is actually unspecified. An end time is still provided for compatibility reasons, even if this attribute is set to True. The default is False.")),
		mcp.WithArray("attendees", mcp.Description("Input parameter: The attendees of the event. See the Events with attendees guide for more information on scheduling events with other calendar users. Service accounts need to use domain-wide delegation of authority to populate the attendee list.")),
		mcp.WithObject("gadget", mcp.Description("Input parameter: A gadget that extends this event. Gadgets are deprecated; this structure is instead only used for returning birthday calendar metadata.")),
		mcp.WithString("id", mcp.Description("Input parameter: Opaque identifier of the event. When creating new single or recurring events, you can specify their IDs. Provided IDs must follow these rules:  \n- characters allowed in the ID are those used in base32hex encoding, i.e. lowercase letters a-v and digits 0-9, see section 3.1.2 in RFC2938 \n- the length of the ID must be between 5 and 1024 characters \n- the ID must be unique per calendar  Due to the globally distributed nature of the system, we cannot guarantee that ID collisions will be detected at event creation time. To minimize the risk of collisions we recommend using an established UUID algorithm such as one described in RFC4122.\nIf you do not specify an ID, it will be automatically generated by the server.\nNote that the icalUID and the id are not identical and only one of them should be supplied at event creation time. One difference in their semantics is that in recurring events, all occurrences of one event have different ids while they all share the same icalUIDs.")),
		mcp.WithString("htmlLink", mcp.Description("Input parameter: An absolute link to this event in the Google Calendar Web UI. Read-only.")),
		mcp.WithObject("workingLocationProperties", mcp.Description("")),
		mcp.WithBoolean("anyoneCanAddSelf", mcp.Description("Input parameter: Whether anyone can invite themselves to the event (deprecated). Optional. The default is False.")),
		mcp.WithObject("outOfOfficeProperties", mcp.Description("")),
		mcp.WithBoolean("attendeesOmitted", mcp.Description("Input parameter: Whether attendees may have been omitted from the event's representation. When retrieving an event, this may be due to a restriction specified by the maxAttendee query parameter. When updating an event, this can be used to only update the participant's response. Optional. The default is False.")),
		mcp.WithObject("end", mcp.Description("")),
		mcp.WithObject("start", mcp.Description("")),
		mcp.WithString("created", mcp.Description("Input parameter: Creation time of the event (as a RFC3339 timestamp). Read-only.")),
		mcp.WithBoolean("guestsCanModify", mcp.Description("Input parameter: Whether attendees other than the organizer can modify the event. Optional. The default is False.")),
		mcp.WithString("eventType", mcp.Description("Input parameter: Specific type of the event. This cannot be modified after the event is created. Possible values are:  \n- \"default\" - A regular event or not further specified. \n- \"outOfOffice\" - An out-of-office event. \n- \"focusTime\" - A focus-time event. \n- \"workingLocation\" - A working location event.  Currently, only \"default \" and \"workingLocation\" events can be created using the API. Extended support for other event types will be made available in later releases.")),
		mcp.WithObject("reminders", mcp.Description("Input parameter: Information about the event's reminders for the authenticated user.")),
		mcp.WithString("hangoutLink", mcp.Description("Input parameter: An absolute link to the Google Hangout associated with this event. Read-only.")),
		mcp.WithObject("organizer", mcp.Description("Input parameter: The organizer of the event. If the organizer is also an attendee, this is indicated with a separate entry in attendees with the organizer field set to True. To change the organizer, use the move operation. Read-only, except when importing an event.")),
		mcp.WithArray("attachments", mcp.Description("Input parameter: File attachments for the event.\nIn order to modify attachments the supportsAttachments request parameter should be set to true.\nThere can be at most 25 attachments per event,")),
		mcp.WithString("location", mcp.Description("Input parameter: Geographic location of the event as free-form text. Optional.")),
		mcp.WithNumber("sequence", mcp.Description("Input parameter: Sequence number as per iCalendar.")),
		mcp.WithString("etag", mcp.Description("Input parameter: ETag of the resource.")),
		mcp.WithObject("source", mcp.Description("Input parameter: Source from which the event was created. For example, a web page, an email message or any document identifiable by an URL with HTTP or HTTPS scheme. Can only be seen or modified by the creator of the event.")),
		mcp.WithBoolean("locked", mcp.Description("Input parameter: Whether this is a locked event copy where no changes can be made to the main event fields \"summary\", \"description\", \"location\", \"start\", \"end\" or \"recurrence\". The default is False. Read-Only.")),
		mcp.WithString("recurringEventId", mcp.Description("Input parameter: For an instance of a recurring event, this is the id of the recurring event to which this instance belongs. Immutable.")),
		mcp.WithBoolean("guestsCanInviteOthers", mcp.Description("Input parameter: Whether attendees other than the organizer can invite others to the event. Optional. The default is True.")),
		mcp.WithString("colorId", mcp.Description("Input parameter: The color of the event. This is an ID referring to an entry in the event section of the colors definition (see the  colors endpoint). Optional.")),
		mcp.WithString("iCalUID", mcp.Description("Input parameter: Event unique identifier as defined in RFC5545. It is used to uniquely identify events accross calendaring systems and must be supplied when importing events via the import method.\nNote that the iCalUID and the id are not identical and only one of them should be supplied at event creation time. One difference in their semantics is that in recurring events, all occurrences of one event have different ids while they all share the same iCalUIDs. To retrieve an event using its iCalUID, call the events.list method using the iCalUID parameter. To retrieve an event using its id, call the events.get method.")),
		mcp.WithBoolean("privateCopy", mcp.Description("Input parameter: If set to True, Event propagation is disabled. Note that it is not the same thing as Private event properties. Optional. Immutable. The default is False.")),
		mcp.WithString("transparency", mcp.Description("Input parameter: Whether the event blocks time on the calendar. Optional. Possible values are:  \n- \"opaque\" - Default value. The event does block time on the calendar. This is equivalent to setting Show me as to Busy in the Calendar UI. \n- \"transparent\" - The event does not block time on the calendar. This is equivalent to setting Show me as to Available in the Calendar UI.")),
		mcp.WithString("description", mcp.Description("Input parameter: Description of the event. Can contain HTML. Optional.")),
		mcp.WithString("visibility", mcp.Description("Input parameter: Visibility of the event. Optional. Possible values are:  \n- \"default\" - Uses the default visibility for events on the calendar. This is the default value. \n- \"public\" - The event is public and event details are visible to all readers of the calendar. \n- \"private\" - The event is private and only event attendees may view event details. \n- \"confidential\" - The event is private. This value is provided for compatibility reasons.")),
		mcp.WithObject("originalStartTime", mcp.Description("")),
		mcp.WithObject("extendedProperties", mcp.Description("Input parameter: Extended properties of the event.")),
		mcp.WithObject("focusTimeProperties", mcp.Description("")),
		mcp.WithBoolean("guestsCanSeeOtherGuests", mcp.Description("Input parameter: Whether attendees other than the organizer can see who the event's attendees are. Optional. The default is True.")),
		mcp.WithString("summary", mcp.Description("Input parameter: Title of the event.")),
		mcp.WithString("updated", mcp.Description("Input parameter: Last modification time of the event (as a RFC3339 timestamp). Read-only.")),
		mcp.WithString("kind", mcp.Description("Input parameter: Type of the resource (\"calendar#event\").")),
		mcp.WithObject("creator", mcp.Description("Input parameter: The creator of the event. Read-only.")),
		mcp.WithArray("recurrence", mcp.Description("Input parameter: List of RRULE, EXRULE, RDATE and EXDATE lines for a recurring event, as specified in RFC5545. Note that DTSTART and DTEND lines are not allowed in this field; event start and end times are specified in the start and end fields. This field is omitted for single events or instances of recurring events.")),
		mcp.WithObject("conferenceData", mcp.Description("")),
	)

	return models.Tool{
		Definition: tool,
		Handler:    Calendar_events_importHandler(cfg),
	}
}
