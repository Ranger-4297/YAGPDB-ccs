{{/*
		Made by ranger_4297 (765316548516380732)

	Trigger Type: `Regex`
	Trigger: `\A(-|<@!?204255221017214977>\s*)(give-?money|loan-?money|pay)(\s+|\z)`

	©️ RhykerWells 2020-Present
	GNU, GPLV3 License
	Repository: https://github.com/Ranger-4297/YAGPDB-ccs
*/}}


{{/* Only edit below if you know what you're doing (: rawr */}}

{{/* Initiates variables */}}
{{$userID := .User.ID}}
{{$successColor := 0x00ff7b}}
{{$errorColor := 0xFF0000}}
{{$prefix := .ServerPrefix }}

{{/* Give money */}}

{{/* Response */}}
{{$embed := sdict "author" (sdict "name" User.Username "icon_url" (User.AvatarURL "1024")) "timestamp" currentTime "color" $errorColor}}
{{$economySettings := (dbGet 0 "EconomySettings").Value}}
{{if not $economySettings}}
	{{$embed.Set "description" (print "No `Settings` database found.\nPlease set it up with the default values using `" .ServerPrefix "server-set default`")}}
	{{sendMessage nil (cembed $embed)}}
	{{return}}
{{end}}
{{$symbol := $economySettings.symbol}}
{{$cash := or (dbGet $userID "cash").Value 0 | toInt}}
{{if not .CmdArgs}}
	{{$embed.Set "description" (print "No `User` argument provided.\nSyntax is `" .Cmd " <User:Mention/ID> <Amount:Amount>`")}}
	{{sendMessage nil (cembed $embed)}}
	{{return}}
{{end}}
{{$user := index .CmdArgs 0}}
{{if not (getMember $user)}}
	{{$embed.Set "description" (print "Invalid `User` argument provided.\nSyntax is `" .Cmd " <User:Mention/ID> <Amount:Amount>`")}}
	{{sendMessage nil (cembed $embed)}}
	{{return}}
{{end}}
{{$receivingUser := (getMember $user).User.ID}}
{{if eq $receivingUser $userID}}
	{{$embed.Set "description" (print "You can't rob yourself.")}}
	{{sendMessage nil (cembed $embed)}}
	{{return}}
{{end}}
{{if not (gt (len .CmdArgs) 1)}}
	{{$embed.Set "description" (print "No `Amount` argument provided.\nSyntax is `" .Cmd " <User:Mention/ID> <Amount:Amount>`")}}
	{{sendMessage nil (cembed $embed)}}
	{{return}}
{{end}}
{{$amount := index .CmdArgs 1}}
{{if not (or (toInt $amount) (eq $amount "all"))}}
	{{$embed.Set "description" (print "Invalid `Amount` argument provided.\nSyntax is `" .Cmd " <User:Mention/ID> <Amount:Amount>`")}}
	{{sendMessage nil (cembed $embed)}}
	{{return}}
{{end}}
{{if eq $amount "all"}}{{$amount = $cash}}{{else}}{{$amount = toInt $amount}}{{end}}
{{if gt $amount $cash}}
	{{$embed.Set "description" (print "You're unable to give more than you have on hand.\nYou currently have " $symbol (humanizeThousands $cash) " on you.")}}
	{{sendMessage nil (cembed $embed)}}
	{{return}}
{{end}}	
{{$embed.Set "description" (print "You gave " $symbol (humanizeThousands $amount) " to <@!" $receivingUser ">\nYou now have " $symbol (humanizeThousands $cash) " in cash!")}}
{{$embed.Set "color" $successColor}}
{{$receivingCash := or (dbGet $receivingUser "cash").Value 0 | toInt}}
{{$receivingCash = add $receivingCash $amount}}
{{$cash = sub $cash $amount}}
{{dbSet $receivingUser "cash" $receivingCash}}
{{dbSet $userID "cash" $cash}}
{{sendMessage nil (cembed $embed)}}