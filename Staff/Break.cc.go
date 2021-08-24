{{/*
        Made by Ranger (765316548516380732)
        Credits: Devonte (622146791659405313)

    Trigger Type: `Command`
    Trigger: `Break`
    Usage: `-Break <time>`
©️ Dynamic 2021
MIT License
*/}}


{{/* Configuration values start */}}
{{$rankID := 784132357002625047}} {{/* Staff rank ID | e.g moderators or administration */}}
{{$separatorID := 784132355379036196}} {{/* Staff role seperator ID | e.g a role named "Staff roles" | remove if none */}}
{{/* Configuration values end */}}

{{/* Only edit below if you know what you're doing (: rawr */}}

{{$args := parseArgs 1 "" (carg "duration" "time")}}
{{if not .ExecData}}
    {{removeRoleID $rankID}} {{/*Staff Role*/}}
    {{removeRoleID $separatorID}} {{/* Remove if none */}}
    You are now on a staff break for {{$args.Get 0}}
    {{execCC .CCID nil (toInt ($args.Get 0).Seconds) true}}
{{else}}
    {{addRoleID $rankID}}
    {{addRoleID $separatorID}} {{/* Remove if none*/}}
{{end}}
