{{/*
		Made by ranger_4297 (765316548516380732)

	Trigger Type: `Regex`
	Trigger: `\A(-|<@!?204255221017214977>\s*)((add-?)?response(s)?)(\s+|\z)`

	©️ RhykerWells 2020-Present
	GNU, GPLV3 License
	Repository: https://github.com/Ranger-4297/YAGPDB-ccs
*/}}


{{/* Only edit below if you know what you're doing (: rawr */}}

{{/* Initiates variables */}}
{{$userID := .User.ID}}
{{$successColor := 0x00ff7b}}
{{$errorColor := 0xFF0000}}
{{$prefix := .ServerPrefix}}

{{/* Add response */}}

{{/* Response */}}
{{$embed := sdict "author" (sdict "name" .User.Username "icon_url" (.User.AvatarURL "1024")) "timestamp" currentTime "color" $errorColor}}
{{$perms := split (index (split (exec "viewperms") "\n") 2) ", "}}
{{if not (or (in $perms "Administrator") (in $perms "ManageServer"))}}
	{{$embed.Set "description" (print "Insufficient permissions.\nTo use this command you need to have either `Administrator` or `ManageServer` permissions")}}
	{{sendMessage nil (cembed $embed)}}
	{{return}}
{{end}}
{{$economySettings := (dbGet 0 "EconomySettings").Value}}
{{if not $economySettings}}
	{{$embed.Set "description" (print "No `Settings` database found.\nPlease set it up with the default values using `" .ServerPrefix "server-set default`")}}
	{{sendMessage nil (cembed $embed)}}
	{{return}}
{{end}}
{{$symbol := $economySettings.symbol}}
{{$responses := or ($economySettings.responses) (sdict "crime" cslice "work" cslice)}}
{{if not .CmdArgs}}
	{{if reFind `add-?responses?` .Cmd}}
		{{$embed.Set "description" (print "No `type` argument passed.\nSyntax is: `" .Cmd " <Type:Work/Crime> <Reponse>`")}}
	{{else}}
		{{$embed.Set "description" (print "No `CommandType` argument passed.\nSyntax is: `" .Cmd " <CommandType:List> <Type:Work/Crime>`")}}
	{{end}}
	{{sendMessage nil (cembed $embed)}}
	{{return}}
{{end}}
{{$type := (index .CmdArgs 0) | lower}}
{{if and (reFind `add-?responses?` .Cmd) (not (eq $type "work" "crime"))}}
	{{$embed.Set "description" (print "Invalid `type` argument passed.\nSyntax is: `" .Cmd " <Type:Work/Crime> <Reponse>`")}}
	{{sendMessage nil (cembed $embed)}}
	{{return}}
{{else if and (not (reFind `add-?responses?` .Cmd)) (not (eq $type "list"))}}
	{{$embed.Set "description" (print "Invalid `CommandType` argument passed.\nSyntax is: `" .Cmd " <CommandType:List> <Type:Work/Crime>`")}}
	{{sendMessage nil (cembed $embed)}}
	{{return}}
{{end}}
{{if (reFind `add-?responses?` .Cmd)}}
	{{if lt (len .CmdArgs) 2}}
		{{$embed.Set "description" (print "No `Response` argument passed.\nSyntax is: `" .Cmd " <Type:Work/Crime> <Reponse>`")}}
		{{sendMessage nil (cembed $embed)}}
		{{return}}
	{{end}}
	{{$response := (joinStr " " (slice .CmdArgs 1))}}
	{{if not (reFind `\(amount\)` $response)}}
		{{$embed.Set "description" (print "Please include the exact string `(amount)` as a placeholder for where the amount goes.")}}
		{{sendMessage nil (cembed $embed)}}
		{{return}}
	{{end}}
	{{$currentResponses := $responses.Get $type}}
	{{$currentResponses = $currentResponses.Append $response}}
	{{$responses.Set $type $currentResponses}}
	{{$embed.Set "description" (print "Successfully added \"" (reReplace `\(amount\)` $response "`(amount)`") "\" to the list of responses for `" $type "`.\nTo view the list of responses in it's entirey. Please use the command `" $prefix "responses list`")}}
	{{$embed.Set "color" $successColor}}
	{{$economySettings.Set "responses" $responses}}
{{else}}
	{{if lt (len .CmdArgs) 2}}
		{{$embed.Set "description" (print "No `Type` argument passed.\nSyntax is: `" .Cmd " <CommandType:List> <Type:Work/Crime>`")}}
		{{sendMessage nil (cembed $embed)}}
		{{return}}
	{{end}}
	{{$type = index .CmdArgs 1 | lower}}
	{{if not (eq $type "work" "crime")}}
		{{$embed.Set "description" (print "Invalid `type` argument passed.\nSyntax is: `" .Cmd " <CommandType:List> <Type:Work/Crime>`")}}
		{{sendMessage nil (cembed $embed)}}
		{{return}}
	{{end}}
	{{$typeResponses := $responses.Get $type}}
	{{if not $typeResponses}}
		{{$embed.Set "description" (print "No responses for this type have been set.\nAdd some with `" $prefix "add-response <Type:Work/Crime> <Response>`")}}
		{{sendMessage nil (cembed $embed)}}
		{{return}}
	{{end}}
	{{$responseNumber := 1}}
	{{$fields := cslice}}
	{{range $typeResponses}}
		{{- $fields = $fields.Append (sdict "Name" (print "Response: " $responseNumber) "value" (reReplace `\(amount\)` . "`(amount)`") "inline" false) -}}
		{{$responseNumber = $responseNumber | add 1}}
	{{end}}
	{{$embed.Set "description" (print "All responses for " $type "-based payouts are\nPlease note. This doesn't support pages yet.")}}
	{{$embed.Set "fields" $fields}}
	{{$embed.Set "color" $successColor}}
{{end}}
{{dbSet 0 "EconomySettings" $economySettings}}
{{sendMessage nil (cembed $embed)}}