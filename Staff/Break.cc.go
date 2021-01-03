{{/*
        Made by Rhyker/Ranger (779096217853886504)
        Helped by Devonte (622146791659405313) (or it might've been Dev who made it but def for me, since I see parseArgs and they confuse me)

    Trigger Type: `Command`
    Trigger: `Break` | Syntax: `-Break <time>` | e.g `-Break 1y3mo8w6d7h9m1s`
©️ Dynamic 2021
MIT License
*/}}


{{/* Configuration values start */}}
{{$rankID := 784132357002625047}} {{/* Staff rank ID | e.g moderators or administratio */}}
{{$separatorID := 784132355379036196}} {{/* Staff role seperator ID | e.g a role named "Staff roles" | remove if none */}}
{{/* Configuration values end */}}

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
