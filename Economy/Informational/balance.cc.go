{{/*
		Made by Ranger (765316548516380732)

	Trigger Type: `Regex`
	Trigger: `\A(-|<@!?204255221017214977>\s*)(bal(ance)?|wallet|money)(\s+|\z)`

	©️ Ranger 2020-Present
	GNU, GPLV3 License
	Repository: https://github.com/Ranger-4297/YAGPDB-ccs
*/}}


{{/* Only edit below if you know what you're doing (: rawr */}}

{{/* Initiates variables */}}
{{$user := .User}}
{{$successColor := 0x00ff7b}}
{{$errorColor := 0xFF0000}}
{{/* $prefix := index (reFindAllSubmatches `.*?: \x60(.*)\x60\z` (execAdmin "Prefix")) 0 1 */}}
{{$prefix := .ServerPrefix}}

{{/* Balance */}}

{{/* Response */}}
{{$embed := sdict "author" (sdict "name" .User.Username "icon_url" (.User.AvatarURL "1024")) "timestamp" currentTime}}
{{/* $embed.Set "author" (sdict "name" .User.Username "icon_url" (.User.AvatarURL "1024"))}}
{{$embed.Set "timestamp" currentTime */}}
{{with (dbGet 0 "EconomySettings")}}
	{{$a := sdict .Value}}
	{{$symbol := $a.symbol}}
	{{with $.CmdArgs}}
		{{$newUser := (index . 0) | userArg}}
		{{if $newUser}}
			{{$user = $newUser}}
		{{end}}
	{{end}}
	{{if not (dbGet $user.ID "EconomyInfo")}}
		{{dbSet $user.ID "EconomyInfo" (sdict "cash" 200 "bank" 0)}}
	{{end}}
	{{with (dbGet $user.ID "EconomyInfo")}}
		{{$a = sdict .Value}}
		{{$cash := ($a.cash | toInt)}}
		{{$bank := ($a.bank | toInt)}}
		{{$net := humanizeThousands ($cash | add $bank)}}
		{{$embed.Set "author" (sdict "name" $user.Username "icon_url" ($user.AvatarURL "128"))}}
		{{$embed.Set "description" (print $user.Mention "'s balance")}}
		{{$embed.Set "fields" (cslice 
			(sdict "name" "Cash" "value" (print $symbol (humanizeThousands $cash)) "inline" true)
			(sdict "name" "Bank" "value" (print $symbol (humanizeThousands $bank)) "inline" true)
			(sdict "name" "Networth" "value" (print $symbol $net) "inline" true))}}
		{{$embed.Set "color" $successColor}}
		{{/* $embed.Set "timestamp" currentTime */}}
	{{end}}
{{else}}
	{{$embed.Set "description" (print "No database found.\nPlease set it up with the default values using `" $prefix "set default`")}}
	{{$embed.Set "color" $errorColor}}
{{end}}
{{sendMessage nil (cembed $embed)}}
