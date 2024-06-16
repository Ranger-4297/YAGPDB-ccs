{{/*
		Made by ranger_4297 (765316548516380732)

	Trigger Type: `Regex`
	Trigger: `\A(-|<@!?204255221017214977>\s*)(l(eader)?-?b(oard)?|top)(\s+|\z)`

	©️ RhykerWells 2020-Present
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
{{$embed := sdict "author" (sdict "name" .User.Username "icon_url" (.User.AvatarURL "1024")) "timestamp" currentTime}}
{{$economySettings := (dbGet 0 "EconomySettings").Value}}
{{if not $economySettings}}
	{{$embed.Set "description" (print "No `Settings` database found.\nPlease set it up with the default values using `" .ServerPrefix "server-set default`")}}
	{{sendMessage nil (cembed $embed)}}
	{{return}}
{{end}}
{{$symbol := $economySettings.symbol}}
{{$page := 1}}
{{with .CmdArgs}}
	{{if $page = (index . 0) | toInt}}{{end}}
{{end}}
{{$rank := mult (sub $page 1) 10}}
{{$users := dbTopEntries "cash" 10 $rank}}
{{if not (len $users)}}
	{{$embed.Set "description" "There were no users on this page"}}
	{{$embed.Set "color" $errorColor}}
	{{sendMessage nil (cembed $embed)}}
	{{return}}
{{end}}
{{$pos := dict 1 "🥇" 2 "🥈" 3 "🥉"}}
{{$display := ""}}
{{$dRank := $rank}}
{{range $users}}
	{{$leaderboard := or (toString (dbGet .User.ID "userEconData").Value.settings.leaderboard) "yes"}}
	{{if eq $leaderboard "yes"}}
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
{{end}}
{{$embed.Set "description" $display}}
{{$embed.Set "footer" (sdict "text" (joinStr "" "Page " $page))}}
{{$embed.Set "color" $successColor}}
{{sendMessage nil (cembed $embed)}}