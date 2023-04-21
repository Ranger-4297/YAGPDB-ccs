{{/*
		Made by Ranger (765316548516380732)

	Trigger Type: `Regex`
	Trigger: `\A(-|<@!?204255221017214977>\s*)(l(eader)?-?b(oard)?|top)(\s+|\z)`

	©️ Ranger 2020-Present
	GNU, GPLV3 License
	Repository: https://github.com/Ranger-4297/YAGPDB-ccs
*/}}


{{/* Only edit below if you know what you're doing (: rawr */}}

{{/* Initiates variables */}}
{{$successColor := 0x00ff7b}}
{{$errorColor := 0xFF0000}}
{{$prefix := .ServerPrefix}}
{{$ex := or (and (reFind "a_" .Guild.Icon) "gif") "png"}}
{{$icon := print "https://cdn.discordapp.com/icons/" .Guild.ID "/" .Guild.Icon "." $ex "?size=1024"}}

{{/* Leaderboard */}}

{{/* Response */}}
{{$embed := sdict}}
{{$embed.Set "author" (sdict "name" (print .Guild.Name " leaderboard") "icon_url" $icon)}}
{{$embed.Set "timestamp" currentTime}}
{{with (dbGet 0 "EconomySettings")}}
	{{$a := sdict .Value}}
	{{$symbol := $a.symbol}}
	{{$page := 1}}
	{{if $.CmdArgs}}
		{{if (index $.CmdArgs 0) | toInt}}
			{{$page = (index $.CmdArgs 0)}}
		{{end}}
	{{end}}
	{{$skip := mult (sub $page 1) 10}}
	{{$pos := dict 1 "🥇" 2 "🥈" 3 "🥉"}}
	{{$users := dbTopEntries "cash" 10 $skip}}
	{{if (len $users)}}
		{{$rank := $skip}}
		{{$display := ""}}
		{{$dRank := $rank}}
		{{range $users}}
			{{$cash := humanizeThousands (toInt .Value)}}
			{{$rank = add $rank 1}}
			{{$dRank = $rank}}
			{{if in (cslice 1 2 3) $rank}}
				{{- $dRank = $pos.Get $rank -}}
			{{else}}
				{{$dRank = print "  " $rank "."}}
			{{end}}
			{{$display = (print $display "**" $dRank "** " .User.String  " **•** " $symbol $cash "\n")}}
		{{end}}
		{{$embed.Set "description" $display}}
		{{$embed.Set "footer" (sdict "text" (joinStr "" "Page " $page))}}
		{{$embed.Set "color" $successColor}}
	{{else}}
		{{$embed.Set "description" "There were no users on this page"}}
		{{$embed.Set "color" $errorColor}}
	{{end}}
{{else}}
	{{$embed.Set "description" (print "No `Settings` database found.\nPlease set it up with the default values using `" $prefix "set default`")}}
	{{$embed.Set "color" $errorColor}}
{{end}}
{{sendMessage nil (cembed $embed)}}