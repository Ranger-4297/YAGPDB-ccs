{{/*
		Made by ranger_4297 (765316548516380732)

	Trigger Type: `Regex`
	Trigger: `\A(-|<@!?204255221017214977>\s*)(bal(ance)?|wallet|money)(\s+|\z)`

	©️ RhykerWells 2020-Present
	GNU, GPLV3 License
	Repository: https://github.com/Ranger-4297/YAGPDB-ccs
*/}}


{{/* Only edit below if you know what you're doing (: rawr */}}

{{/* Initiates variables */}}
{{$user := .User}}
{{$successColor := 0x00ff7b}}
{{$errorColor := 0xFF0000}}
{{$prefix := .ServerPrefix}}

{{/* Balance */}}

{{/* Response */}}
{{$embed := sdict "author" (sdict "name" .User.Username "icon_url" (.User.AvatarURL "1024")) "timestamp" currentTime}}
{{$economySettings := (dbGet 0 "EconomySettings").Value}}
{{if not $economySettings}}
	{{$embed.Set "description" (print "No `Settings` database found.\nPlease set it up with the default values using `" .ServerPrefix "server-set default`")}}
	{{sendMessage nil (cembed $embed)}}
	{{return}}
{{end}}
{{$symbol := $economySettings.symbol}}
{{with .CmdArgs}}
	{{if $newUser := (index . 0) | userArg}}
		{{$user = $newUser}}
	{{end}}
{{end}}
{{$bank := or (((dbGet 0 "bank").Value).Get (toString $user.ID)) 0 | toInt}}
{{$cash := or (dbGet $user.ID "cash").Value 0 | toInt}}
{{$net := humanizeThousands ($cash | add $bank)}}
{{$rank := ""}}
{{$pos := dict 1 "🥇" 2 "🥈" 3 "🥉"}}
{{$rank = dbRank (sdict "pattern" "cash") $user.ID "cash"}}
{{if not $rank}}
	{{$rank = "None"}}
{{else}}
	{{if in (cslice 1 2 3) $rank }}
		{{$rank = $pos.Get $rank}}
	{{else}}
		{{$ord := "th"}}
		{{$cent := toInt (mod $rank 100)}}
		{{$dec := toInt (mod $rank 10)}}
		{{if not (and (ge $cent 10) (le $cent 19))}}
			{{if eq $dec 1}}{{$ord = "st"}}{{else if eq $dec 2}}{{$ord = "nd"}}{{else if eq $dec 3}}{{$ord = "rd"}}{{end}}
		{{end}}
		{{$rank = print $rank $ord "."}}
	{{end}}
{{end}}
{{$embed.Set "description" (print $user.Mention "'s balance\nLeaderboard rank: " $rank)}}
{{$embed.Set "fields" (cslice (sdict "name" "Cash" "value" (print $symbol (humanizeThousands $cash)) "inline" true) (sdict "name" "Bank" "value" (print $symbol (humanizeThousands $bank)) "inline" true) (sdict "name" "Networth" "value" (print $symbol $net) "inline" true))}}
{{$embed.Set "color" $successColor}}
{{if not (and (eq .User.ID $user.ID) (eq $balStatus "no"))}}
	{{sendMessage nil (cembed $embed)}}
	{{return}}
{{end}}
{{sendDM (cembed $embed)}}
{{$embed.Set "description" "Sent this to your DM as your inventory is on private"}}
{{$embed.Del "footer"}}
{{$embed.Del "fields"}}
{{sendMessage nil (cembed $embed)}}