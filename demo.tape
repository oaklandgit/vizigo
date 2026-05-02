Output demo.gif

Set FontSize 22
Set Width 1200
Set Height 800
Set TypingSpeed 0.1s

Type "go run ."
Sleep 500ms
Enter
Sleep 1s

# Row 1: headers
Type "Category"
Tab
Sleep 150ms
Type "Budget"
Tab
Sleep 150ms
Type "Actual"
Enter
Sleep 150ms
Left 2
Down

# Row 3: Rent
Type "Rent"
Tab
Sleep 250ms
Type@150ms "1500"
Tab
Sleep 250ms
Type@150ms "1500"
Enter
Sleep 250ms
Left 2

# Row 4: Groceries
Type "Groceries"
Tab
Sleep 250ms
Type@150ms "400"
Tab
Sleep 250ms
Type@150ms "423"
Enter
Sleep 150ms
Left 2

# Row 5: Utilities
Type "Utilities"
Tab
Sleep 150ms
Type@150ms "120"
Tab
Sleep 150ms
Type@150ms "134"
Enter
Sleep 250ms
Left 2

# Row 6: Transport
Type "Transport"
Tab
Sleep 150ms
Type@120ms "80"
Tab
Sleep 150ms
Type@120ms "65"
Enter
Sleep 150ms
Left 2

# Rows 7-8 are empty — skip to row 9
Down
Down

# Row 9: Totals
Type "Total"
Tab
Sleep 150ms
Type@120ms "=SUM(B3:B6)"
Tab
Sleep 500ms
Type@120ms "=SUM(C3:C6)"
Enter
Sleep 500ms
Left 2
Down

# Row 11: Over budget?
Type "Over budget?"
Tab
Sleep 150ms
Type@160ms `=C9>B9 ? "yes" : "no"`
Enter
Sleep 500ms

# Land on the result
Up
Right
Sleep 4s

# VHS documentation
#
# Output:
#   Output <path>.gif               Create a GIF output at the given <path>
#   Output <path>.mp4               Create an MP4 output at the given <path>
#   Output <path>.webm              Create a WebM output at the given <path>
#
# Require:
#   Require <string>                Ensure a program is on the $PATH to proceed
#
# Settings:
#   Set FontSize <number>           Set the font size of the terminal
#   Set FontFamily <string>         Set the font family of the terminal
#   Set Height <number>             Set the height of the terminal
#   Set Width <number>              Set the width of the terminal
#   Set LetterSpacing <float>       Set the font letter spacing (tracking)
#   Set LineHeight <float>          Set the font line height
#   Set LoopOffset <float>%         Set the starting frame offset for the GIF loop
#   Set Theme <json|string>         Set the theme of the terminal
#   Set Padding <number>            Set the padding of the terminal
#   Set Framerate <number>          Set the framerate of the recording
#   Set PlaybackSpeed <float>       Set the playback speed of the recording
#   Set MarginFill <file|#000000>   Set the file or color the margin will be filled with.
#   Set Margin <number>             Set the size of the margin. Has no effect if MarginFill isn't set.
#   Set BorderRadius <number>       Set terminal border radius, in pixels.
#   Set WindowBar <string>          Set window bar type. (one of: Rings, RingsRight, Colorful, ColorfulRight)
#   Set WindowBarSize <number>      Set window bar size, in pixels. Default is 40.
#   Set TypingSpeed <time>          Set the typing speed of the terminal. Default is 50ms.
#
# Sleep:
#   Sleep <time>                    Sleep for a set amount of <time> in seconds
#
# Type:
#   Type[@<time>] "<characters>"    Type <characters> into the terminal with a
#                                   <time> delay between each character
#
# Keys:
#   Escape[@<time>] [number]        Press the Escape key
#   Backspace[@<time>] [number]     Press the Backspace key
#   Delete[@<time>] [number]        Press the Delete key
#   Insert[@<time>] [number]        Press the Insert key
#   Down[@<time>] [number]          Press the Down key
#   Enter[@<time>] [number]         Press the Enter key
#   Space[@<time>] [number]         Press the Space key
#   Tab[@<time>] [number]           Press the Tab key
#   Left[@<time>] [number]          Press the Left Arrow key
#   Right[@<time>] [number]         Press the Right Arrow key
#   Up[@<time>] [number]            Press the Up Arrow key
#   Down[@<time>] [number]          Press the Down Arrow key
#   PageUp[@<time>] [number]        Press the Page Up key
#   PageDown[@<time>] [number]      Press the Page Down key
#   Ctrl+<key>                      Press the Control key + <key> (e.g. Ctrl+C)
#
# Display:
#   Hide                            Hide the subsequent commands from the output
#   Show                            Show the subsequent commands in the output
#   ScrollUp[@<time>] [number]      Scroll terminal viewport up by rows
#   ScrollDown[@<time>] [number]    Scroll terminal viewport down by rows
