{{/*
        Made by Ranger (765316548516380732)

    Trigger Type: `Join message in channel`
©️ Ranger 2021
MIT License
*/}}

{{/* Only edit below if you know what you're doing (: rawr */}}

{{/* Response */}}
{{with (dbGet 0 "EconomySettings")}}
    {{$a := sdict .Value}}
    {{$startBalance := (toInt $a.startBalance)}}
    {{dbSet .User.ID "EconomyInfo" (sdict "cash" $startBalance "bank" 0)}}
{{end}}