package main

import (
	"github.com/calendar-api/mcp-server/config"
	"github.com/calendar-api/mcp-server/models"
	tools_calendars "github.com/calendar-api/mcp-server/tools/calendars"
	tools_acl "github.com/calendar-api/mcp-server/tools/acl"
	tools_colors "github.com/calendar-api/mcp-server/tools/colors"
	tools_freebusy "github.com/calendar-api/mcp-server/tools/freebusy"
	tools_calendarlist "github.com/calendar-api/mcp-server/tools/calendarlist"
	tools_settings "github.com/calendar-api/mcp-server/tools/settings"
	tools_events "github.com/calendar-api/mcp-server/tools/events"
	tools_channels "github.com/calendar-api/mcp-server/tools/channels"
)

func GetAll(cfg *config.APIConfig) []models.Tool {
	return []models.Tool{
		tools_calendars.CreateCalendar_calendars_insertTool(cfg),
		tools_acl.CreateCalendar_acl_watchTool(cfg),
		tools_calendars.CreateCalendar_calendars_clearTool(cfg),
		tools_colors.CreateCalendar_colors_getTool(cfg),
		tools_freebusy.CreateCalendar_freebusy_queryTool(cfg),
		tools_calendarlist.CreateCalendar_calendarlist_listTool(cfg),
		tools_calendarlist.CreateCalendar_calendarlist_insertTool(cfg),
		tools_calendarlist.CreateCalendar_calendarlist_watchTool(cfg),
		tools_settings.CreateCalendar_settings_getTool(cfg),
		tools_calendarlist.CreateCalendar_calendarlist_deleteTool(cfg),
		tools_calendarlist.CreateCalendar_calendarlist_getTool(cfg),
		tools_calendarlist.CreateCalendar_calendarlist_patchTool(cfg),
		tools_calendarlist.CreateCalendar_calendarlist_updateTool(cfg),
		tools_acl.CreateCalendar_acl_deleteTool(cfg),
		tools_acl.CreateCalendar_acl_getTool(cfg),
		tools_acl.CreateCalendar_acl_patchTool(cfg),
		tools_acl.CreateCalendar_acl_updateTool(cfg),
		tools_events.CreateCalendar_events_quickaddTool(cfg),
		tools_events.CreateCalendar_events_deleteTool(cfg),
		tools_events.CreateCalendar_events_getTool(cfg),
		tools_events.CreateCalendar_events_patchTool(cfg),
		tools_events.CreateCalendar_events_updateTool(cfg),
		tools_channels.CreateCalendar_channels_stopTool(cfg),
		tools_acl.CreateCalendar_acl_listTool(cfg),
		tools_acl.CreateCalendar_acl_insertTool(cfg),
		tools_events.CreateCalendar_events_moveTool(cfg),
		tools_events.CreateCalendar_events_instancesTool(cfg),
		tools_settings.CreateCalendar_settings_watchTool(cfg),
		tools_events.CreateCalendar_events_listTool(cfg),
		tools_events.CreateCalendar_events_insertTool(cfg),
		tools_events.CreateCalendar_events_importTool(cfg),
		tools_events.CreateCalendar_events_watchTool(cfg),
		tools_calendars.CreateCalendar_calendars_getTool(cfg),
		tools_calendars.CreateCalendar_calendars_patchTool(cfg),
		tools_calendars.CreateCalendar_calendars_updateTool(cfg),
		tools_calendars.CreateCalendar_calendars_deleteTool(cfg),
		tools_settings.CreateCalendar_settings_listTool(cfg),
	}
}
