package models

import (
	"context"
	"github.com/mark3labs/mcp-go/mcp"
)

type Tool struct {
	Definition mcp.Tool
	Handler    func(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error)
}

// Events represents the Events schema from the OpenAPI specification
type Events struct {
	Timezone string `json:"timeZone,omitempty"` // The time zone of the calendar. Read-only.
	Updated string `json:"updated,omitempty"` // Last modification time of the calendar (as a RFC3339 timestamp). Read-only.
	Accessrole string `json:"accessRole,omitempty"` // The user's access role for this calendar. Read-only. Possible values are: - "none" - The user has no access. - "freeBusyReader" - The user has read access to free/busy information. - "reader" - The user has read access to the calendar. Private events will appear to users with reader access, but event details will be hidden. - "writer" - The user has read and write access to the calendar. Private events will appear to users with writer access, and event details will be visible. - "owner" - The user has ownership of the calendar. This role has all of the permissions of the writer role with the additional ability to see and manipulate ACLs.
	Defaultreminders []EventReminder `json:"defaultReminders,omitempty"` // The default reminders on the calendar for the authenticated user. These reminders apply to all events on this calendar that do not explicitly override them (i.e. do not have reminders.useDefault set to True).
	Description string `json:"description,omitempty"` // Description of the calendar. Read-only.
	Etag string `json:"etag,omitempty"` // ETag of the collection.
	Items []Event `json:"items,omitempty"` // List of events on the calendar.
	Kind string `json:"kind,omitempty"` // Type of the collection ("calendar#events").
	Nextsynctoken string `json:"nextSyncToken,omitempty"` // Token used at a later point in time to retrieve only the entries that have changed since this result was returned. Omitted if further results are available, in which case nextPageToken is provided.
	Nextpagetoken string `json:"nextPageToken,omitempty"` // Token used to access the next page of this result. Omitted if no further results are available, in which case nextSyncToken is provided.
	Summary string `json:"summary,omitempty"` // Title of the calendar. Read-only.
}

// EventReminder represents the EventReminder schema from the OpenAPI specification
type EventReminder struct {
	Method string `json:"method,omitempty"` // The method used by this reminder. Possible values are: - "email" - Reminders are sent via email. - "popup" - Reminders are sent via a UI popup. Required when adding a reminder.
	Minutes int `json:"minutes,omitempty"` // Number of minutes before the start of the event when the reminder should trigger. Valid values are between 0 and 40320 (4 weeks in minutes). Required when adding a reminder.
}

// ConferenceRequestStatus represents the ConferenceRequestStatus schema from the OpenAPI specification
type ConferenceRequestStatus struct {
	Statuscode string `json:"statusCode,omitempty"` // The current status of the conference create request. Read-only. The possible values are: - "pending": the conference create request is still being processed. - "success": the conference create request succeeded, the entry points are populated. - "failure": the conference create request failed, there are no entry points.
}

// EventWorkingLocationProperties represents the EventWorkingLocationProperties schema from the OpenAPI specification
type EventWorkingLocationProperties struct {
	TypeField string `json:"type,omitempty"` // Type of the working location. Possible values are: - "homeOffice" - The user is working at home. - "officeLocation" - The user is working from an office. - "customLocation" - The user is working from a custom location. Any details are specified in a sub-field of the specified name, but this field may be missing if empty. Any other fields are ignored. Required when adding working location properties.
	Customlocation map[string]interface{} `json:"customLocation,omitempty"` // If present, specifies that the user is working from a custom location.
	Homeoffice interface{} `json:"homeOffice,omitempty"` // If present, specifies that the user is working at home.
	Officelocation map[string]interface{} `json:"officeLocation,omitempty"` // If present, specifies that the user is working from an office.
}

// EventDateTime represents the EventDateTime schema from the OpenAPI specification
type EventDateTime struct {
	Date string `json:"date,omitempty"` // The date, in the format "yyyy-mm-dd", if this is an all-day event.
	Datetime string `json:"dateTime,omitempty"` // The time, as a combined date-time value (formatted according to RFC3339). A time zone offset is required unless a time zone is explicitly specified in timeZone.
	Timezone string `json:"timeZone,omitempty"` // The time zone in which the time is specified. (Formatted as an IANA Time Zone Database name, e.g. "Europe/Zurich".) For recurring events this field is required and specifies the time zone in which the recurrence is expanded. For single events this field is optional and indicates a custom time zone for the event start/end.
}

// Settings represents the Settings schema from the OpenAPI specification
type Settings struct {
	Kind string `json:"kind,omitempty"` // Type of the collection ("calendar#settings").
	Nextpagetoken string `json:"nextPageToken,omitempty"` // Token used to access the next page of this result. Omitted if no further results are available, in which case nextSyncToken is provided.
	Nextsynctoken string `json:"nextSyncToken,omitempty"` // Token used at a later point in time to retrieve only the entries that have changed since this result was returned. Omitted if further results are available, in which case nextPageToken is provided.
	Etag string `json:"etag,omitempty"` // Etag of the collection.
	Items []Setting `json:"items,omitempty"` // List of user settings.
}

// FreeBusyRequest represents the FreeBusyRequest schema from the OpenAPI specification
type FreeBusyRequest struct {
	Timezone string `json:"timeZone,omitempty"` // Time zone used in the response. Optional. The default is UTC.
	Calendarexpansionmax int `json:"calendarExpansionMax,omitempty"` // Maximal number of calendars for which FreeBusy information is to be provided. Optional. Maximum value is 50.
	Groupexpansionmax int `json:"groupExpansionMax,omitempty"` // Maximal number of calendar identifiers to be provided for a single group. Optional. An error is returned for a group with more members than this value. Maximum value is 100.
	Items []FreeBusyRequestItem `json:"items,omitempty"` // List of calendars and/or groups to query.
	Timemax string `json:"timeMax,omitempty"` // The end of the interval for the query formatted as per RFC3339.
	Timemin string `json:"timeMin,omitempty"` // The start of the interval for the query formatted as per RFC3339.
}

// CalendarList represents the CalendarList schema from the OpenAPI specification
type CalendarList struct {
	Nextsynctoken string `json:"nextSyncToken,omitempty"` // Token used at a later point in time to retrieve only the entries that have changed since this result was returned. Omitted if further results are available, in which case nextPageToken is provided.
	Etag string `json:"etag,omitempty"` // ETag of the collection.
	Items []CalendarListEntry `json:"items,omitempty"` // Calendars that are present on the user's calendar list.
	Kind string `json:"kind,omitempty"` // Type of the collection ("calendar#calendarList").
	Nextpagetoken string `json:"nextPageToken,omitempty"` // Token used to access the next page of this result. Omitted if no further results are available, in which case nextSyncToken is provided.
}

// EventOutOfOfficeProperties represents the EventOutOfOfficeProperties schema from the OpenAPI specification
type EventOutOfOfficeProperties struct {
	Autodeclinemode string `json:"autoDeclineMode,omitempty"` // Whether to decline meeting invitations which overlap Out of office events. Valid values are declineNone, meaning that no meeting invitations are declined; declineAllConflictingInvitations, meaning that all conflicting meeting invitations that conflict with the event are declined; and declineOnlyNewConflictingInvitations, meaning that only new conflicting meeting invitations which arrive while the Out of office event is present are to be declined.
	Declinemessage string `json:"declineMessage,omitempty"` // Response message to set if an existing event or new invitation is automatically declined by Calendar.
}

// FreeBusyGroup represents the FreeBusyGroup schema from the OpenAPI specification
type FreeBusyGroup struct {
	Calendars []string `json:"calendars,omitempty"` // List of calendars' identifiers within a group.
	Errors []Error `json:"errors,omitempty"` // Optional error(s) (if computation for the group failed).
}

// Setting represents the Setting schema from the OpenAPI specification
type Setting struct {
	Id string `json:"id,omitempty"` // The id of the user setting.
	Kind string `json:"kind,omitempty"` // Type of the resource ("calendar#setting").
	Value string `json:"value,omitempty"` // Value of the user setting. The format of the value depends on the ID of the setting. It must always be a UTF-8 string of length up to 1024 characters.
	Etag string `json:"etag,omitempty"` // ETag of the resource.
}

// Event represents the Event schema from the OpenAPI specification
type Event struct {
	Location string `json:"location,omitempty"` // Geographic location of the event as free-form text. Optional.
	Sequence int `json:"sequence,omitempty"` // Sequence number as per iCalendar.
	Etag string `json:"etag,omitempty"` // ETag of the resource.
	Source map[string]interface{} `json:"source,omitempty"` // Source from which the event was created. For example, a web page, an email message or any document identifiable by an URL with HTTP or HTTPS scheme. Can only be seen or modified by the creator of the event.
	Locked bool `json:"locked,omitempty"` // Whether this is a locked event copy where no changes can be made to the main event fields "summary", "description", "location", "start", "end" or "recurrence". The default is False. Read-Only.
	Recurringeventid string `json:"recurringEventId,omitempty"` // For an instance of a recurring event, this is the id of the recurring event to which this instance belongs. Immutable.
	Guestscaninviteothers bool `json:"guestsCanInviteOthers,omitempty"` // Whether attendees other than the organizer can invite others to the event. Optional. The default is True.
	Colorid string `json:"colorId,omitempty"` // The color of the event. This is an ID referring to an entry in the event section of the colors definition (see the colors endpoint). Optional.
	Icaluid string `json:"iCalUID,omitempty"` // Event unique identifier as defined in RFC5545. It is used to uniquely identify events accross calendaring systems and must be supplied when importing events via the import method. Note that the iCalUID and the id are not identical and only one of them should be supplied at event creation time. One difference in their semantics is that in recurring events, all occurrences of one event have different ids while they all share the same iCalUIDs. To retrieve an event using its iCalUID, call the events.list method using the iCalUID parameter. To retrieve an event using its id, call the events.get method.
	Privatecopy bool `json:"privateCopy,omitempty"` // If set to True, Event propagation is disabled. Note that it is not the same thing as Private event properties. Optional. Immutable. The default is False.
	Transparency string `json:"transparency,omitempty"` // Whether the event blocks time on the calendar. Optional. Possible values are: - "opaque" - Default value. The event does block time on the calendar. This is equivalent to setting Show me as to Busy in the Calendar UI. - "transparent" - The event does not block time on the calendar. This is equivalent to setting Show me as to Available in the Calendar UI.
	Description string `json:"description,omitempty"` // Description of the event. Can contain HTML. Optional.
	Visibility string `json:"visibility,omitempty"` // Visibility of the event. Optional. Possible values are: - "default" - Uses the default visibility for events on the calendar. This is the default value. - "public" - The event is public and event details are visible to all readers of the calendar. - "private" - The event is private and only event attendees may view event details. - "confidential" - The event is private. This value is provided for compatibility reasons.
	Originalstarttime EventDateTime `json:"originalStartTime,omitempty"`
	Extendedproperties map[string]interface{} `json:"extendedProperties,omitempty"` // Extended properties of the event.
	Focustimeproperties EventFocusTimeProperties `json:"focusTimeProperties,omitempty"`
	Guestscanseeotherguests bool `json:"guestsCanSeeOtherGuests,omitempty"` // Whether attendees other than the organizer can see who the event's attendees are. Optional. The default is True.
	Summary string `json:"summary,omitempty"` // Title of the event.
	Updated string `json:"updated,omitempty"` // Last modification time of the event (as a RFC3339 timestamp). Read-only.
	Kind string `json:"kind,omitempty"` // Type of the resource ("calendar#event").
	Creator map[string]interface{} `json:"creator,omitempty"` // The creator of the event. Read-only.
	Recurrence []string `json:"recurrence,omitempty"` // List of RRULE, EXRULE, RDATE and EXDATE lines for a recurring event, as specified in RFC5545. Note that DTSTART and DTEND lines are not allowed in this field; event start and end times are specified in the start and end fields. This field is omitted for single events or instances of recurring events.
	Conferencedata ConferenceData `json:"conferenceData,omitempty"`
	Status string `json:"status,omitempty"` // Status of the event. Optional. Possible values are: - "confirmed" - The event is confirmed. This is the default status. - "tentative" - The event is tentatively confirmed. - "cancelled" - The event is cancelled (deleted). The list method returns cancelled events only on incremental sync (when syncToken or updatedMin are specified) or if the showDeleted flag is set to true. The get method always returns them. A cancelled status represents two different states depending on the event type: - Cancelled exceptions of an uncancelled recurring event indicate that this instance should no longer be presented to the user. Clients should store these events for the lifetime of the parent recurring event. Cancelled exceptions are only guaranteed to have values for the id, recurringEventId and originalStartTime fields populated. The other fields might be empty. - All other cancelled events represent deleted events. Clients should remove their locally synced copies. Such cancelled events will eventually disappear, so do not rely on them being available indefinitely. Deleted events are only guaranteed to have the id field populated. On the organizer's calendar, cancelled events continue to expose event details (summary, location, etc.) so that they can be restored (undeleted). Similarly, the events to which the user was invited and that they manually removed continue to provide details. However, incremental sync requests with showDeleted set to false will not return these details. If an event changes its organizer (for example via the move operation) and the original organizer is not on the attendee list, it will leave behind a cancelled event where only the id field is guaranteed to be populated.
	Endtimeunspecified bool `json:"endTimeUnspecified,omitempty"` // Whether the end time is actually unspecified. An end time is still provided for compatibility reasons, even if this attribute is set to True. The default is False.
	Attendees []EventAttendee `json:"attendees,omitempty"` // The attendees of the event. See the Events with attendees guide for more information on scheduling events with other calendar users. Service accounts need to use domain-wide delegation of authority to populate the attendee list.
	Gadget map[string]interface{} `json:"gadget,omitempty"` // A gadget that extends this event. Gadgets are deprecated; this structure is instead only used for returning birthday calendar metadata.
	Id string `json:"id,omitempty"` // Opaque identifier of the event. When creating new single or recurring events, you can specify their IDs. Provided IDs must follow these rules: - characters allowed in the ID are those used in base32hex encoding, i.e. lowercase letters a-v and digits 0-9, see section 3.1.2 in RFC2938 - the length of the ID must be between 5 and 1024 characters - the ID must be unique per calendar Due to the globally distributed nature of the system, we cannot guarantee that ID collisions will be detected at event creation time. To minimize the risk of collisions we recommend using an established UUID algorithm such as one described in RFC4122. If you do not specify an ID, it will be automatically generated by the server. Note that the icalUID and the id are not identical and only one of them should be supplied at event creation time. One difference in their semantics is that in recurring events, all occurrences of one event have different ids while they all share the same icalUIDs.
	Htmllink string `json:"htmlLink,omitempty"` // An absolute link to this event in the Google Calendar Web UI. Read-only.
	Workinglocationproperties EventWorkingLocationProperties `json:"workingLocationProperties,omitempty"`
	Anyonecanaddself bool `json:"anyoneCanAddSelf,omitempty"` // Whether anyone can invite themselves to the event (deprecated). Optional. The default is False.
	Outofofficeproperties EventOutOfOfficeProperties `json:"outOfOfficeProperties,omitempty"`
	Attendeesomitted bool `json:"attendeesOmitted,omitempty"` // Whether attendees may have been omitted from the event's representation. When retrieving an event, this may be due to a restriction specified by the maxAttendee query parameter. When updating an event, this can be used to only update the participant's response. Optional. The default is False.
	End EventDateTime `json:"end,omitempty"`
	Start EventDateTime `json:"start,omitempty"`
	Created string `json:"created,omitempty"` // Creation time of the event (as a RFC3339 timestamp). Read-only.
	Guestscanmodify bool `json:"guestsCanModify,omitempty"` // Whether attendees other than the organizer can modify the event. Optional. The default is False.
	Eventtype string `json:"eventType,omitempty"` // Specific type of the event. This cannot be modified after the event is created. Possible values are: - "default" - A regular event or not further specified. - "outOfOffice" - An out-of-office event. - "focusTime" - A focus-time event. - "workingLocation" - A working location event. Currently, only "default " and "workingLocation" events can be created using the API. Extended support for other event types will be made available in later releases.
	Reminders map[string]interface{} `json:"reminders,omitempty"` // Information about the event's reminders for the authenticated user.
	Hangoutlink string `json:"hangoutLink,omitempty"` // An absolute link to the Google Hangout associated with this event. Read-only.
	Organizer map[string]interface{} `json:"organizer,omitempty"` // The organizer of the event. If the organizer is also an attendee, this is indicated with a separate entry in attendees with the organizer field set to True. To change the organizer, use the move operation. Read-only, except when importing an event.
	Attachments []EventAttachment `json:"attachments,omitempty"` // File attachments for the event. In order to modify attachments the supportsAttachments request parameter should be set to true. There can be at most 25 attachments per event,
}

// ConferenceParameters represents the ConferenceParameters schema from the OpenAPI specification
type ConferenceParameters struct {
	Addonparameters ConferenceParametersAddOnParameters `json:"addOnParameters,omitempty"`
}

// CalendarListEntry represents the CalendarListEntry schema from the OpenAPI specification
type CalendarListEntry struct {
	Defaultreminders []EventReminder `json:"defaultReminders,omitempty"` // The default reminders that the authenticated user has for this calendar.
	Notificationsettings map[string]interface{} `json:"notificationSettings,omitempty"` // The notifications that the authenticated user is receiving for this calendar.
	Backgroundcolor string `json:"backgroundColor,omitempty"` // The main color of the calendar in the hexadecimal format "#0088aa". This property supersedes the index-based colorId property. To set or change this property, you need to specify colorRgbFormat=true in the parameters of the insert, update and patch methods. Optional.
	Description string `json:"description,omitempty"` // Description of the calendar. Optional. Read-only.
	Selected bool `json:"selected,omitempty"` // Whether the calendar content shows up in the calendar UI. Optional. The default is False.
	Summary string `json:"summary,omitempty"` // Title of the calendar. Read-only.
	Id string `json:"id,omitempty"` // Identifier of the calendar.
	Conferenceproperties ConferenceProperties `json:"conferenceProperties,omitempty"`
	Hidden bool `json:"hidden,omitempty"` // Whether the calendar has been hidden from the list. Optional. The attribute is only returned when the calendar is hidden, in which case the value is true.
	Summaryoverride string `json:"summaryOverride,omitempty"` // The summary that the authenticated user has set for this calendar. Optional.
	Accessrole string `json:"accessRole,omitempty"` // The effective access role that the authenticated user has on the calendar. Read-only. Possible values are: - "freeBusyReader" - Provides read access to free/busy information. - "reader" - Provides read access to the calendar. Private events will appear to users with reader access, but event details will be hidden. - "writer" - Provides read and write access to the calendar. Private events will appear to users with writer access, and event details will be visible. - "owner" - Provides ownership of the calendar. This role has all of the permissions of the writer role with the additional ability to see and manipulate ACLs.
	Timezone string `json:"timeZone,omitempty"` // The time zone of the calendar. Optional. Read-only.
	Foregroundcolor string `json:"foregroundColor,omitempty"` // The foreground color of the calendar in the hexadecimal format "#ffffff". This property supersedes the index-based colorId property. To set or change this property, you need to specify colorRgbFormat=true in the parameters of the insert, update and patch methods. Optional.
	Location string `json:"location,omitempty"` // Geographic location of the calendar as free-form text. Optional. Read-only.
	Deleted bool `json:"deleted,omitempty"` // Whether this calendar list entry has been deleted from the calendar list. Read-only. Optional. The default is False.
	Primary bool `json:"primary,omitempty"` // Whether the calendar is the primary calendar of the authenticated user. Read-only. Optional. The default is False.
	Colorid string `json:"colorId,omitempty"` // The color of the calendar. This is an ID referring to an entry in the calendar section of the colors definition (see the colors endpoint). This property is superseded by the backgroundColor and foregroundColor properties and can be ignored when using these properties. Optional.
	Kind string `json:"kind,omitempty"` // Type of the resource ("calendar#calendarListEntry").
	Etag string `json:"etag,omitempty"` // ETag of the resource.
}

// Colors represents the Colors schema from the OpenAPI specification
type Colors struct {
	Event map[string]interface{} `json:"event,omitempty"` // A global palette of event colors, mapping from the color ID to its definition. An event resource may refer to one of these color IDs in its colorId field. Read-only.
	Kind string `json:"kind,omitempty"` // Type of the resource ("calendar#colors").
	Updated string `json:"updated,omitempty"` // Last modification time of the color palette (as a RFC3339 timestamp). Read-only.
	Calendar map[string]interface{} `json:"calendar,omitempty"` // A global palette of calendar colors, mapping from the color ID to its definition. A calendarListEntry resource refers to one of these color IDs in its colorId field. Read-only.
}

// FreeBusyCalendar represents the FreeBusyCalendar schema from the OpenAPI specification
type FreeBusyCalendar struct {
	Busy []TimePeriod `json:"busy,omitempty"` // List of time ranges during which this calendar should be regarded as busy.
	Errors []Error `json:"errors,omitempty"` // Optional error(s) (if computation for the calendar failed).
}

// ConferenceParametersAddOnParameters represents the ConferenceParametersAddOnParameters schema from the OpenAPI specification
type ConferenceParametersAddOnParameters struct {
	Parameters map[string]interface{} `json:"parameters,omitempty"`
}

// Channel represents the Channel schema from the OpenAPI specification
type Channel struct {
	TypeField string `json:"type,omitempty"` // The type of delivery mechanism used for this channel. Valid values are "web_hook" (or "webhook"). Both values refer to a channel where Http requests are used to deliver messages.
	Address string `json:"address,omitempty"` // The address where notifications are delivered for this channel.
	Payload bool `json:"payload,omitempty"` // A Boolean value to indicate whether payload is wanted. Optional.
	Resourceid string `json:"resourceId,omitempty"` // An opaque ID that identifies the resource being watched on this channel. Stable across different API versions.
	Expiration string `json:"expiration,omitempty"` // Date and time of notification channel expiration, expressed as a Unix timestamp, in milliseconds. Optional.
	Params map[string]interface{} `json:"params,omitempty"` // Additional parameters controlling delivery channel behavior. Optional.
	Resourceuri string `json:"resourceUri,omitempty"` // A version-specific identifier for the watched resource.
	Token string `json:"token,omitempty"` // An arbitrary string delivered to the target address with each notification delivered over this channel. Optional.
	Id string `json:"id,omitempty"` // A UUID or similar unique string that identifies this channel.
	Kind string `json:"kind,omitempty"` // Identifies this as a notification channel used to watch for changes to a resource, which is "api#channel".
}

// Calendar represents the Calendar schema from the OpenAPI specification
type Calendar struct {
	Location string `json:"location,omitempty"` // Geographic location of the calendar as free-form text. Optional.
	Summary string `json:"summary,omitempty"` // Title of the calendar.
	Timezone string `json:"timeZone,omitempty"` // The time zone of the calendar. (Formatted as an IANA Time Zone Database name, e.g. "Europe/Zurich".) Optional.
	Conferenceproperties ConferenceProperties `json:"conferenceProperties,omitempty"`
	Description string `json:"description,omitempty"` // Description of the calendar. Optional.
	Etag string `json:"etag,omitempty"` // ETag of the resource.
	Id string `json:"id,omitempty"` // Identifier of the calendar. To retrieve IDs call the calendarList.list() method.
	Kind string `json:"kind,omitempty"` // Type of the resource ("calendar#calendar").
}

// EventAttendee represents the EventAttendee schema from the OpenAPI specification
type EventAttendee struct {
	Comment string `json:"comment,omitempty"` // The attendee's response comment. Optional.
	Id string `json:"id,omitempty"` // The attendee's Profile ID, if available.
	Optional bool `json:"optional,omitempty"` // Whether this is an optional attendee. Optional. The default is False.
	Responsestatus string `json:"responseStatus,omitempty"` // The attendee's response status. Possible values are: - "needsAction" - The attendee has not responded to the invitation (recommended for new events). - "declined" - The attendee has declined the invitation. - "tentative" - The attendee has tentatively accepted the invitation. - "accepted" - The attendee has accepted the invitation. Warning: If you add an event using the values declined, tentative, or accepted, attendees with the "Add invitations to my calendar" setting set to "When I respond to invitation in email" won't see an event on their calendar unless they choose to change their invitation response in the event invitation email.
	Additionalguests int `json:"additionalGuests,omitempty"` // Number of additional guests. Optional. The default is 0.
	Email string `json:"email,omitempty"` // The attendee's email address, if available. This field must be present when adding an attendee. It must be a valid email address as per RFC5322. Required when adding an attendee.
	Organizer bool `json:"organizer,omitempty"` // Whether the attendee is the organizer of the event. Read-only. The default is False.
	Self bool `json:"self,omitempty"` // Whether this entry represents the calendar on which this copy of the event appears. Read-only. The default is False.
	Resource bool `json:"resource,omitempty"` // Whether the attendee is a resource. Can only be set when the attendee is added to the event for the first time. Subsequent modifications are ignored. Optional. The default is False.
	Displayname string `json:"displayName,omitempty"` // The attendee's name, if available. Optional.
}

// AclRule represents the AclRule schema from the OpenAPI specification
type AclRule struct {
	Role string `json:"role,omitempty"` // The role assigned to the scope. Possible values are: - "none" - Provides no access. - "freeBusyReader" - Provides read access to free/busy information. - "reader" - Provides read access to the calendar. Private events will appear to users with reader access, but event details will be hidden. - "writer" - Provides read and write access to the calendar. Private events will appear to users with writer access, and event details will be visible. - "owner" - Provides ownership of the calendar. This role has all of the permissions of the writer role with the additional ability to see and manipulate ACLs.
	Scope map[string]interface{} `json:"scope,omitempty"` // The extent to which calendar access is granted by this ACL rule.
	Etag string `json:"etag,omitempty"` // ETag of the resource.
	Id string `json:"id,omitempty"` // Identifier of the Access Control List (ACL) rule. See Sharing calendars.
	Kind string `json:"kind,omitempty"` // Type of the resource ("calendar#aclRule").
}

// ConferenceData represents the ConferenceData schema from the OpenAPI specification
type ConferenceData struct {
	Parameters ConferenceParameters `json:"parameters,omitempty"`
	Signature string `json:"signature,omitempty"` // The signature of the conference data. Generated on server side. Unset for a conference with a failed create request. Optional for a conference with a pending create request.
	Conferenceid string `json:"conferenceId,omitempty"` // The ID of the conference. Can be used by developers to keep track of conferences, should not be displayed to users. The ID value is formed differently for each conference solution type: - eventHangout: ID is not set. (This conference type is deprecated.) - eventNamedHangout: ID is the name of the Hangout. (This conference type is deprecated.) - hangoutsMeet: ID is the 10-letter meeting code, for example aaa-bbbb-ccc. - addOn: ID is defined by the third-party provider. Optional.
	Conferencesolution ConferenceSolution `json:"conferenceSolution,omitempty"`
	Createrequest CreateConferenceRequest `json:"createRequest,omitempty"`
	Entrypoints []EntryPoint `json:"entryPoints,omitempty"` // Information about individual conference entry points, such as URLs or phone numbers. All of them must belong to the same conference. Either conferenceSolution and at least one entryPoint, or createRequest is required.
	Notes string `json:"notes,omitempty"` // Additional notes (such as instructions from the domain administrator, legal notices) to display to the user. Can contain HTML. The maximum length is 2048 characters. Optional.
}

// Error represents the Error schema from the OpenAPI specification
type Error struct {
	Domain string `json:"domain,omitempty"` // Domain, or broad category, of the error.
	Reason string `json:"reason,omitempty"` // Specific reason for the error. Some of the possible values are: - "groupTooBig" - The group of users requested is too large for a single query. - "tooManyCalendarsRequested" - The number of calendars requested is too large for a single query. - "notFound" - The requested resource was not found. - "internalError" - The API service has encountered an internal error. Additional error types may be added in the future, so clients should gracefully handle additional error statuses not included in this list.
}

// ColorDefinition represents the ColorDefinition schema from the OpenAPI specification
type ColorDefinition struct {
	Background string `json:"background,omitempty"` // The background color associated with this color definition.
	Foreground string `json:"foreground,omitempty"` // The foreground color that can be used to write on top of a background with 'background' color.
}

// CalendarNotification represents the CalendarNotification schema from the OpenAPI specification
type CalendarNotification struct {
	Method string `json:"method,omitempty"` // The method used to deliver the notification. The possible value is: - "email" - Notifications are sent via email. Required when adding a notification.
	TypeField string `json:"type,omitempty"` // The type of notification. Possible values are: - "eventCreation" - Notification sent when a new event is put on the calendar. - "eventChange" - Notification sent when an event is changed. - "eventCancellation" - Notification sent when an event is cancelled. - "eventResponse" - Notification sent when an attendee responds to the event invitation. - "agenda" - An agenda with the events of the day (sent out in the morning). Required when adding a notification.
}

// ConferenceSolutionKey represents the ConferenceSolutionKey schema from the OpenAPI specification
type ConferenceSolutionKey struct {
	TypeField string `json:"type,omitempty"` // The conference solution type. If a client encounters an unfamiliar or empty type, it should still be able to display the entry points. However, it should disallow modifications. The possible values are: - "eventHangout" for Hangouts for consumers (deprecated; existing events may show this conference solution type but new conferences cannot be created) - "eventNamedHangout" for classic Hangouts for Google Workspace users (deprecated; existing events may show this conference solution type but new conferences cannot be created) - "hangoutsMeet" for Google Meet (http://meet.google.com) - "addOn" for 3P conference providers
}

// EntryPoint represents the EntryPoint schema from the OpenAPI specification
type EntryPoint struct {
	Password string `json:"password,omitempty"` // The password to access the conference. The maximum length is 128 characters. When creating new conference data, populate only the subset of {meetingCode, accessCode, passcode, password, pin} fields that match the terminology that the conference provider uses. Only the populated fields should be displayed. Optional.
	Uri string `json:"uri,omitempty"` // The URI of the entry point. The maximum length is 1300 characters. Format: - for video, http: or https: schema is required. - for phone, tel: schema is required. The URI should include the entire dial sequence (e.g., tel:+12345678900,,,123456789;1234). - for sip, sip: schema is required, e.g., sip:12345678@myprovider.com. - for more, http: or https: schema is required.
	Entrypointfeatures []string `json:"entryPointFeatures,omitempty"` // Features of the entry point, such as being toll or toll-free. One entry point can have multiple features. However, toll and toll-free cannot be both set on the same entry point.
	Label string `json:"label,omitempty"` // The label for the URI. Visible to end users. Not localized. The maximum length is 512 characters. Examples: - for video: meet.google.com/aaa-bbbb-ccc - for phone: +1 123 268 2601 - for sip: 12345678@altostrat.com - for more: should not be filled Optional.
	Meetingcode string `json:"meetingCode,omitempty"` // The meeting code to access the conference. The maximum length is 128 characters. When creating new conference data, populate only the subset of {meetingCode, accessCode, passcode, password, pin} fields that match the terminology that the conference provider uses. Only the populated fields should be displayed. Optional.
	Regioncode string `json:"regionCode,omitempty"` // The CLDR/ISO 3166 region code for the country associated with this phone access. Example: "SE" for Sweden. Calendar backend will populate this field only for EntryPointType.PHONE.
	Accesscode string `json:"accessCode,omitempty"` // The access code to access the conference. The maximum length is 128 characters. When creating new conference data, populate only the subset of {meetingCode, accessCode, passcode, password, pin} fields that match the terminology that the conference provider uses. Only the populated fields should be displayed. Optional.
	Passcode string `json:"passcode,omitempty"` // The passcode to access the conference. The maximum length is 128 characters. When creating new conference data, populate only the subset of {meetingCode, accessCode, passcode, password, pin} fields that match the terminology that the conference provider uses. Only the populated fields should be displayed.
	Pin string `json:"pin,omitempty"` // The PIN to access the conference. The maximum length is 128 characters. When creating new conference data, populate only the subset of {meetingCode, accessCode, passcode, password, pin} fields that match the terminology that the conference provider uses. Only the populated fields should be displayed. Optional.
	Entrypointtype string `json:"entryPointType,omitempty"` // The type of the conference entry point. Possible values are: - "video" - joining a conference over HTTP. A conference can have zero or one video entry point. - "phone" - joining a conference by dialing a phone number. A conference can have zero or more phone entry points. - "sip" - joining a conference over SIP. A conference can have zero or one sip entry point. - "more" - further conference joining instructions, for example additional phone numbers. A conference can have zero or one more entry point. A conference with only a more entry point is not a valid conference.
}

// FreeBusyResponse represents the FreeBusyResponse schema from the OpenAPI specification
type FreeBusyResponse struct {
	Calendars map[string]interface{} `json:"calendars,omitempty"` // List of free/busy information for calendars.
	Groups map[string]interface{} `json:"groups,omitempty"` // Expansion of groups.
	Kind string `json:"kind,omitempty"` // Type of the resource ("calendar#freeBusy").
	Timemax string `json:"timeMax,omitempty"` // The end of the interval.
	Timemin string `json:"timeMin,omitempty"` // The start of the interval.
}

// ConferenceProperties represents the ConferenceProperties schema from the OpenAPI specification
type ConferenceProperties struct {
	Allowedconferencesolutiontypes []string `json:"allowedConferenceSolutionTypes,omitempty"` // The types of conference solutions that are supported for this calendar. The possible values are: - "eventHangout" - "eventNamedHangout" - "hangoutsMeet" Optional.
}

// FreeBusyRequestItem represents the FreeBusyRequestItem schema from the OpenAPI specification
type FreeBusyRequestItem struct {
	Id string `json:"id,omitempty"` // The identifier of a calendar or a group.
}

// TimePeriod represents the TimePeriod schema from the OpenAPI specification
type TimePeriod struct {
	Start string `json:"start,omitempty"` // The (inclusive) start of the time period.
	End string `json:"end,omitempty"` // The (exclusive) end of the time period.
}

// Acl represents the Acl schema from the OpenAPI specification
type Acl struct {
	Items []AclRule `json:"items,omitempty"` // List of rules on the access control list.
	Kind string `json:"kind,omitempty"` // Type of the collection ("calendar#acl").
	Nextpagetoken string `json:"nextPageToken,omitempty"` // Token used to access the next page of this result. Omitted if no further results are available, in which case nextSyncToken is provided.
	Nextsynctoken string `json:"nextSyncToken,omitempty"` // Token used at a later point in time to retrieve only the entries that have changed since this result was returned. Omitted if further results are available, in which case nextPageToken is provided.
	Etag string `json:"etag,omitempty"` // ETag of the collection.
}

// EventAttachment represents the EventAttachment schema from the OpenAPI specification
type EventAttachment struct {
	Fileid string `json:"fileId,omitempty"` // ID of the attached file. Read-only. For Google Drive files, this is the ID of the corresponding Files resource entry in the Drive API.
	Fileurl string `json:"fileUrl,omitempty"` // URL link to the attachment. For adding Google Drive file attachments use the same format as in alternateLink property of the Files resource in the Drive API. Required when adding an attachment.
	Iconlink string `json:"iconLink,omitempty"` // URL link to the attachment's icon. This field can only be modified for custom third-party attachments.
	Mimetype string `json:"mimeType,omitempty"` // Internet media type (MIME type) of the attachment.
	Title string `json:"title,omitempty"` // Attachment title.
}

// ConferenceSolution represents the ConferenceSolution schema from the OpenAPI specification
type ConferenceSolution struct {
	Key ConferenceSolutionKey `json:"key,omitempty"`
	Name string `json:"name,omitempty"` // The user-visible name of this solution. Not localized.
	Iconuri string `json:"iconUri,omitempty"` // The user-visible icon for this solution.
}

// CreateConferenceRequest represents the CreateConferenceRequest schema from the OpenAPI specification
type CreateConferenceRequest struct {
	Conferencesolutionkey ConferenceSolutionKey `json:"conferenceSolutionKey,omitempty"`
	Requestid string `json:"requestId,omitempty"` // The client-generated unique ID for this request. Clients should regenerate this ID for every new request. If an ID provided is the same as for the previous request, the request is ignored.
	Status ConferenceRequestStatus `json:"status,omitempty"`
}

// EventFocusTimeProperties represents the EventFocusTimeProperties schema from the OpenAPI specification
type EventFocusTimeProperties struct {
	Chatstatus string `json:"chatStatus,omitempty"` // The status to mark the user in Chat and related products. This can be available or doNotDisturb.
	Declinemessage string `json:"declineMessage,omitempty"` // Response message to set if an existing event or new invitation is automatically declined by Calendar.
	Autodeclinemode string `json:"autoDeclineMode,omitempty"` // Whether to decline meeting invitations which overlap Focus Time events. Valid values are declineNone, meaning that no meeting invitations are declined; declineAllConflictingInvitations, meaning that all conflicting meeting invitations that conflict with the event are declined; and declineOnlyNewConflictingInvitations, meaning that only new conflicting meeting invitations which arrive while the Focus Time event is present are to be declined.
}
