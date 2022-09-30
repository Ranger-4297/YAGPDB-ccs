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
{{$prefix := index (reFindAllSubmatches `.*?: \x60(.*)\x60\z` (execAdmin "Prefix")) 0 1}}
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
    {{$leaderboard := dbBottomEntries "EconomyInfo" 10 $skip}}
    {{if (len $leaderboard)}}
        {{$rank := $skip}}
        {{$display := ""}}
        {{range $leaderboard}}
            {{$cash := humanizeThousands (toInt .Value.cash)}}
            {{$rank = add $rank 1}}
            {{$display = (print $display "**" $rank ".** " .User.String  " **•** " $symbol $cash "\n")}}
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