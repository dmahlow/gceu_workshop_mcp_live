package main

import (
	"context"
	"fmt"
	"log"

	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"

	"github.com/dmahlow/desktop-automation-mcp/internal/automation"
)

func main() {
	// Create a new MCP server
	s := server.NewMCPServer(
		"Desktop Automation Server",
		"1.0.0",
		server.WithToolCapabilities(true),
		server.WithRecovery(),
	)

	// Add mouse click tool
	clickTool := mcp.NewTool("click",
		mcp.WithDescription("Click at specified coordinates"),
		mcp.WithNumber("x",
			mcp.Required(),
			mcp.Description("X coordinate for click"),
		),
		mcp.WithNumber("y",
			mcp.Required(),
			mcp.Description("Y coordinate for click"),
		),
	)

	s.AddTool(clickTool, func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		x, err := request.RequireFloat("x")
		if err != nil {
			return mcp.NewToolResultError(err.Error()), nil
		}

		y, err := request.RequireFloat("y")
		if err != nil {
			return mcp.NewToolResultError(err.Error()), nil
		}

		err = automation.Click(int(x), int(y))
		if err != nil {
			return mcp.NewToolResultError(fmt.Sprintf("Click failed: %v", err)), nil
		}

		return mcp.NewToolResultText(fmt.Sprintf("Clicked at (%d, %d)", int(x), int(y))), nil
	})

	// Add type text tool
	typeTextTool := mcp.NewTool("type_text",
		mcp.WithDescription("Type text at current cursor position"),
		mcp.WithString("text",
			mcp.Required(),
			mcp.Description("Text to type"),
		),
		mcp.WithNumber("delay",
			mcp.Description("Delay between characters in milliseconds (optional)"),
		),
	)

	s.AddTool(typeTextTool, func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		text, err := request.RequireString("text")
		if err != nil {
			return mcp.NewToolResultError(err.Error()), nil
		}

		args := request.GetArguments()
		delay, hasDelay := args["delay"]

		if hasDelay && delay != nil {
			delayVal := delay.(float64)
			err = automation.TypeStringWithDelay(text, int(delayVal))
		} else {
			err = automation.TypeString(text)
		}

		if err != nil {
			return mcp.NewToolResultError(fmt.Sprintf("Type text failed: %v", err)), nil
		}

		return mcp.NewToolResultText(fmt.Sprintf("Typed: %s", text)), nil
	})

	// Add press key tool
	pressKeyTool := mcp.NewTool("press_key",
		mcp.WithDescription("Press a key or key combination"),
		mcp.WithString("key",
			mcp.Required(),
			mcp.Description("Key to press (e.g., 'enter', 'space', 'ctrl')"),
		),
		mcp.WithArray("modifiers",
			mcp.Description("Modifier keys (e.g., ['ctrl', 'shift'])"),
		),
	)

	s.AddTool(pressKeyTool, func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		key, err := request.RequireString("key")
		if err != nil {
			return mcp.NewToolResultError(err.Error()), nil
		}

		args := request.GetArguments()
		modifiersRaw, hasModifiers := args["modifiers"]

		if hasModifiers && modifiersRaw != nil {
			modifiersArray := modifiersRaw.([]interface{})
			modifiers := make([]string, len(modifiersArray))
			for i, mod := range modifiersArray {
				modifiers[i] = mod.(string)
			}
			keys := append(modifiers, key)
			err = automation.PressKeyCombo(keys...)

			if err != nil {
				return mcp.NewToolResultError(fmt.Sprintf("Press key failed: %v", err)), nil
			}

			return mcp.NewToolResultText(fmt.Sprintf("Pressed key combination: %v + %s", modifiers, key)), nil
		} else {
			err = automation.PressKey(key)

			if err != nil {
				return mcp.NewToolResultError(fmt.Sprintf("Press key failed: %v", err)), nil
			}
		}
		return mcp.NewToolResultText(fmt.Sprintf("Pressed key: %s", key)), nil
	})

	// Add mouse move tool
	moveMouseTool := mcp.NewTool("move_mouse",
		mcp.WithDescription("Move mouse to specified coordinates"),
		mcp.WithNumber("x",
			mcp.Required(),
			mcp.Description("X coordinate to move to"),
		),
		mcp.WithNumber("y",
			mcp.Required(),
			mcp.Description("Y coordinate to move to"),
		),
		mcp.WithBoolean("smooth",
			mcp.Description("Use smooth movement (default: false)"),
		),
		mcp.WithNumber("duration",
			mcp.Description("Duration for smooth movement in seconds (default: 1.0)"),
		),
	)

	s.AddTool(moveMouseTool, func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		x, err := request.RequireFloat("x")
		if err != nil {
			return mcp.NewToolResultError(err.Error()), nil
		}

		y, err := request.RequireFloat("y")
		if err != nil {
			return mcp.NewToolResultError(err.Error()), nil
		}

		args := request.GetArguments()
		smoothRaw, _ := args["smooth"]
		durationRaw, hasDuration := args["duration"]

		smooth := false
		if smoothRaw != nil {
			smooth = smoothRaw.(bool)
		}

		duration := 1.0
		if hasDuration && durationRaw != nil {
			duration = durationRaw.(float64)
		}

		if smooth {
			err = automation.SmoothMove(int(x), int(y), duration)
		} else {
			err = automation.Move(int(x), int(y))
		}

		if err != nil {
			return mcp.NewToolResultError(fmt.Sprintf("Move mouse failed: %v", err)), nil
		}

		return mcp.NewToolResultText(fmt.Sprintf("Moved mouse to (%d, %d)", int(x), int(y))), nil
	})

	// Add get mouse position tool
	getMousePosTool := mcp.NewTool("get_mouse_position",
		mcp.WithDescription("Get current mouse cursor position"),
	)

	s.AddTool(getMousePosTool, func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		x, y := automation.GetPosition()
		return mcp.NewToolResultText(fmt.Sprintf("Mouse position: (%d, %d)", x, y)), nil
	})

	// Add right click tool
	rightClickTool := mcp.NewTool("right_click",
		mcp.WithDescription("Right click at specified coordinates"),
		mcp.WithNumber("x",
			mcp.Required(),
			mcp.Description("X coordinate for right click"),
		),
		mcp.WithNumber("y",
			mcp.Required(),
			mcp.Description("Y coordinate for right click"),
		),
	)

	s.AddTool(rightClickTool, func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		x, err := request.RequireFloat("x")
		if err != nil {
			return mcp.NewToolResultError(err.Error()), nil
		}

		y, err := request.RequireFloat("y")
		if err != nil {
			return mcp.NewToolResultError(err.Error()), nil
		}

		err = automation.RightClick(int(x), int(y))
		if err != nil {
			return mcp.NewToolResultError(fmt.Sprintf("Right click failed: %v", err)), nil
		}

		return mcp.NewToolResultText(fmt.Sprintf("Right clicked at (%d, %d)", int(x), int(y))), nil
	})

	// Add double click tool
	doubleClickTool := mcp.NewTool("double_click",
		mcp.WithDescription("Double click at specified coordinates"),
		mcp.WithNumber("x",
			mcp.Required(),
			mcp.Description("X coordinate for double click"),
		),
		mcp.WithNumber("y",
			mcp.Required(),
			mcp.Description("Y coordinate for double click"),
		),
	)

	s.AddTool(doubleClickTool, func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		x, err := request.RequireFloat("x")
		if err != nil {
			return mcp.NewToolResultError(err.Error()), nil
		}

		y, err := request.RequireFloat("y")
		if err != nil {
			return mcp.NewToolResultError(err.Error()), nil
		}

		err = automation.DoubleClick(int(x), int(y))
		if err != nil {
			return mcp.NewToolResultError(fmt.Sprintf("Double click failed: %v", err)), nil
		}

		return mcp.NewToolResultText(fmt.Sprintf("Double clicked at (%d, %d)", int(x), int(y))), nil
	})

	// Start the stdio server
	log.Println("Starting Desktop Automation MCP Server...")
	if err := server.ServeStdio(s); err != nil {
		log.Fatalf("Server error: %v", err)
	}
}
