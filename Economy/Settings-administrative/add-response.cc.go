{{/*
		Made by ranger_4297 (765316548516380732)

	Trigger Type: `Regex`
	Trigger: `\A(-|<@!?204255221017214977>\s*)((add-?)?response(s)?)(\s+|\z)`

	©️ Ranger 2020-Present
	GNU, GPLV3 License
	Repository: https://github.com/Ranger-4297/YAGPDB-ccs
*/}}

{{/* Only edit below if you know what you're doing (: rawr */}}

{{/* Initiates variables */}}
{{$userID := .User.ID}}
{{$successColor := 0x00ff7b}}
{{$errorColor := 0xFF0000}}
{{$prefix := .ServerPrefix}}

{{/* Set */}}

{{/* Response */}}
{{$embed := sdict}}
{{$embed.Set "author" (sdict "name" $.User.Username "icon_url" ($.User.AvatarURL "1024"))}}
{{$embed.Set "timestamp" currentTime}}
{{$db := or (dbGet 0 "EconomySettings").Value (sdict "min" 200 "max" 500 "betMax" 5000 "symbol" "£" "startBalance" 200 "incomeCooldown" 300 "workCooldown" 7200 "crimeCooldown" 14400 "robCooldown" 21600 "responses" (sdict "crime" cslice "work" cslice))}}
{{$perms := split (index (split (exec "viewperms") "\n") 2) ", "}}
{{if or (in $perms "Administrator") (in $perms "ManageServer")}}
	{{with $db}}
		{{$responses := or ($db.responses) (sdict "crime" cslice "work" cslice)}}
		{{with $.CmdArgs}}
			{{$type := (index $.CmdArgs 0) | lower}}
			{{if eq $type "work" "crime" "list"}}
				{{if reFind `add-?responses?` $.Cmd}}
					{{if eq $type "work" "crime"}}
						{{if gt (len $.CmdArgs) 1}}
							{{$response := (joinStr " " (slice $.CmdArgs 1))}}
							{{if (reFind `\(amount\)` $response)}}
								{{$currentResponses := $responses.Get $type}}
								{{$currentResponses = $currentResponses.Append $response}}
								{{$responses.Set $type $currentResponses}}
								{{$embed.Set "description" (print "Successfully added \"" (reReplace `\(amount\)` $response "`(amount)`") "\" to the list of responses for `" $type "`.\nTo view the list of responses in it's entirey. Please use the command `" $prefix "responses list`")}}
								{{$embed.Set "color" $successColor}}
								{{$db.Set "responses" $responses}}
							{{else}}
								{{$embed.Set "description" (print "Please include the exact string `(amount)` as a placeholder for where the amount goes.")}}
								{{$embed.Set "color" $errorColor}}
							{{end}}
						{{else}}
							{{$embed.Set "description" (print "No `Response` argument passed.\nSyntax is: `" $.Cmd " " $type " <Reponse>`")}}
							{{$embed.Set "color" $errorColor}}
						{{end}}
					{{end}}
				{{else if reFind `responses?` $.Cmd}}
					{{if gt (len $.CmdArgs) 1}}
						{{$type := (index $.CmdArgs 1) | lower}}
						{{if eq $type "work" "crime"}}
							{{$typeResponses := $responses.Get $type}}
							{{if $typeResponses}}
								{{$responseNumber := 1}}
								{{$fields := cslice}}
								{{range $typeResponses}}
									{{- $fields = $fields.Append (sdict "Name" (print "Response: " $responseNumber) "value" (reReplace `\(amount\)` . "`(amount)`") "inline" false) -}}
									{{$responseNumber = $responseNumber | add 1}}
								{{end}}
								{{$embed.Set "description" (print "All responses for " $type "-based payouts are\nPlease note. This doesn't support pages yet.")}}
								{{$embed.Set "fields" $fields}}
								{{$embed.Set "color" $successColor}}
							{{else}}
								{{$embed.Set "description" (print "No responses for this type have been set.\nAdd some with `" $prefix "add-response <Type:Work/Crime> <Response>`")}}
								{{$embed.Set "color" $errorColor}}
							{{end}}
						{{else}}
							{{$embed.Set "description" (print "Invalid `type` argument passed.\nSyntax is: `" $.Cmd " list <Type:Work/Crime>`")}}
							{{$embed.Set "color" $errorColor}}
						{{end}}
					{{else}}
						{{$embed.Set "description" (print "No `type` argument passed.\nSyntax is: `" $.Cmd " list <Type:Work/Crime>`")}}
						{{$embed.Set "color" $errorColor}}
					{{end}}
				{{end}}
			{{else}}
				{{if reFind `add-?responses?` $.Cmd}}
					{{$embed.Set "description" (print "Invalid `type` argument passed.\nSyntax is: `" $.Cmd " <Type:Work/Crime> <Reponse>`")}}
				{{else if reFind `responses?` $.Cmd}}
					{{$embed.Set "description" (print "Invalid `CommandType` argument passed.\nSyntax is: `" $.Cmd " <CommandType:List> <Type:Work/Crime>`")}}
				{{end}}
				{{$embed.Set "color" $errorColor}}
			{{end}}
		{{else}}
			{{if reFind `add-?responses?` $.Cmd}}
				{{$embed.Set "description" (print "No `type` argument passed.\nSyntax is: `" $.Cmd " <Type:Work/Crime> <Reponse>`")}}
			{{else if reFind `responses?` $.Cmd}}
				{{$embed.Set "description" (print "No `CommandType` argument passed.\nSyntax is: `" $.Cmd " <CommandType:List> <Type:Work/Crime>`")}}
			{{end}}
			{{$embed.Set "color" $errorColor}}
		{{end}}
	{{end}}
{{else}}
	{{$embed.Set "description" (print "Insufficient permissions.\nTo use this command you need to have either `Administrator` or `ManageServer` permissions")}}
	{{$embed.Set "color" $errorColor}}
{{end}}
{{dbSet 0 "EconomySettings" $db}}
{{sendMessage nil (cembed $embed)}}