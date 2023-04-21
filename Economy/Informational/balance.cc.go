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
{{$prefix := index (reFindAllSubmatches `.*?: \x60(.*)\x60\z` (execAdmin "Prefix")) 0 1}}

{{/* Balance */}}

{{/* Response */}}
{{$embed := sdict}}
{{$embed.Set "author" (sdict "name" .User.Username "icon_url" (.User.AvatarURL "1024"))}}
{{$embed.Set "timestamp" currentTime}}
{{with (dbGet 0 "EconomySettings")}}
	{{$a := sdict .Value}}
	{{$symbol := $a.symbol}}
	{{with $.CmdArgs}}
		{{$newUser := (index . 0) | userArg}}
		{{if $newUser}}
			{{$user = $newUser}}
		{{end}}
	{{end}}
	{{$bank := or (((dbGet 0 "bank").Value).Get (toString $user.ID)) 0 | toInt}}
	{{$cash := or (dbGet $user.ID "cash").Value 0 | toInt}}
	{{$net := humanizeThousands ($cash | add $bank)}}
	{{$pos := dict 1 "🥇" 2 "🥈" 3 "🥉"}}
	{{$rank := dbRank (sdict "pattern" "cash") $user.ID "cash"}}
	{{if in (cslice 1 2 3) $rank }}
		{{$rank = $pos.Get $rank}}
	{{else}}
	{{$ord := "th"}}
	{{$cent := toInt (mod $rank 100)}}
	{{$dec := toInt (mod $rank 10)}}
	{{if not (and (ge $cent 10) (le $cent 19))}}
		{{if eq $dec 1}}{{$ord = "st"}}
		{{else if eq $dec 2}}{{$ord = "nd"}}
		{{else if eq $dec 3}}{{$ord = "rd"}}
		{{end}}
	{{end}}
		{{$rank = print $rank $ord "."}}
	{{end}}
	{{$embed.Set "author" (sdict "name" $user.Username "icon_url" ($user.AvatarURL "128"))}}
	{{$embed.Set "description" (print $user.Mention "'s balance\nLeaderboard rank: " $rank)}}
	{{$embed.Set "fields" (cslice 
		(sdict "name" "Cash" "value" (print $symbol (humanizeThousands $cash)) "inline" true)
		(sdict "name" "Bank" "value" (print $symbol (humanizeThousands $bank)) "inline" true)
		(sdict "name" "Networth" "value" (print $symbol $net) "inline" true))}}
	{{$embed.Set "color" $successColor}}
	{{$embed.Set "timestamp" currentTime}}
{{else}}
	{{$embed.Set "description" (print "No database found.\nPlease set it up with the default values using `" $prefix "set default`")}}
	{{$embed.Set "color" $errorColor}}
{{end}}
{{sendMessage nil (cembed $embed)}}