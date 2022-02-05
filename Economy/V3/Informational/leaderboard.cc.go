{{/*
        Made by Ranger (765316548516380732)

    Trigger Type: `Regex`
    Trigger: `\A(-|<@!?204255221017214977>\s*)(l(eader)?-?b(oard)?|top)(\s+|\z)`

    ©️ Ranger 2020-Present
    GNU, GPLV3 License
    Repository: https://github.com/Ranger-4297/YAGPDB-ccs
*/}}

{{range (dbTopEntries "EconomyInfo" 10 0)}}
    {{dbSet .User.ID "Leaderboard" ((dbGet .User.ID "EconomyInfo").Value.cash)}}
{{end}}
{{$page := 1}}
{{with reFind `\d+` (joinStr " " .CmdArgs)}}
    {{$page = . | toInt}}
{{end}}
{{$skip := mult (sub $page 1) 10}}
{{$leaderboard := dbTopEntries "Leaderboard" 10 $skip}}
{{$rank := $skip}}
{{$display := ""}}
{{range $leaderboard}}
    {{$cash := toInt .Value }}
    {{$rank = add $rank 1 }}
    {{$display = (printf "%s**%d.** %s **||** %d \n" $display $rank .User.String $cash)}}
    {{dbDel .User.ID "Leaderboard"}}
{{end }}
{{sendMessage nil (cembed
        "title" "Cash leaderboard"
        "description" $display
        "color" 0x00ff7b
        "footer" (sdict "text" (joinStr "" "Page " $page))
)}}