{{/*
		Made by ranger_4297 (765316548516380732)

	Trigger Type: `Regex`
	Trigger: `\A(-|<@!?204255221017214977>\s*)((reset|remove)-?user)(\s+|\z)`

	©️ RhykerWells 2020-Present
	GNU, GPLV3 License
	Repository: https://github.com/Ranger-4297/YAGPDB-ccs
*/}}


{{/* Only edit below if you know what you're doing (: rawr */}}

{{/* Initiates variables */}}
{{$successColor := 0x00ff7b}}
{{$errorColor := 0xFF0000}}
{{$prefix := .ServerPrefix }}

{{/* Reset user */}}

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
{{if not .CmdArgs}}
	{{$embed.Set "description" (print "No `User` argument provided.\nSyntax is `" .Cmd " <User:Mention/ID> <Destination:Cash/Bank> <Amount:Amount>`")}}
	{{sendMessage nil (cembed $embed)}}
	{{return}}
{{end}}
{{$user := index .CmdArgs 0}}
{{if not (getMember $user)}}
	{{$embed.Set "description" (print "Invalid `User` argument provided.\nSyntax is `" .Cmd " <User:Mention/ID> <Destination:Cash/Bank> <Amount:Amount>`")}}
	{{sendMessage nil (cembed $embed)}}
	{{return}}
{{end}}
{{$user = (getMember $user).User.ID}}
{{if (dbGet $user "cash")}}{{dbDel $user "cash"}}{{end}}
{{if (dbGet $user "userEconData")}}{{dbDel $user "userEconData"}}{{end}}
{{$bank := (dbGet 0 "bank").Value}}
{{if $bank}}{{$bank.Del (toString $user)}}{{dbSet 0 "bank" $bank}}{{end}}
{{$embed.Set "description" (print "Successfully removed the user from the economy system.")}}
{{$embed.Set "color" $successColor}}
{{sendMessage nil (cembed $embed)}}