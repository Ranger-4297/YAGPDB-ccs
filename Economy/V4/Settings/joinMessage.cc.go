{{/*
        Made by Ranger (765316548516380732)

    Trigger Type: `Join message in channel`

    ©️ Ranger 2020-Present
    GNU, GPLV3 License
    Repository: https://github.com/Ranger-4297/YAGPDB-ccs
*/}}

{{/* 
Use this in conjunction with leaveMessage
*/}}

{{/* Only edit below if you know what you're doing (: rawr */}}

{{/* Response */}}
{{with (dbGet 0 "EconomySettings")}}
    {{$a := sdict .Value}}
    {{$startBalance := (toInt $a.startBalance)}}
    {{dbSet .User.ID "EconomyInfo" (sdict "cash" $startBalance "bank" 0)}}
{{end}}