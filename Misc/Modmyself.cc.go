{{/*
        Made by Ranger (765316548516380732)

    Trigger Type: `RegEx`
    Trigger: `\A[^a-zA-Z0-9]*\-[​‏]*?m[​‏]*o[​‏]*d[​‏]*m[​‏]*y[​‏]*s[​‏]*e[​‏]*[lI][​‏]*f`
©️ Ranger 2021
MIT License
*/}}

{{/* Configuration values start */}}
{{$serverLink := "https://discord.gg/ekMQH384KC"}} {{/* Your servers invite link */}}
{{/* Configuration values end */}}

{{/* Only edit below if you know what you're doing (: rawr */}}

{{$random := randInt 4}}
{{if le $random 0}}
    {{execAdmin "warn" .User.ID "you fuckin' donut, you modded yourself"}}
{{else if le $random 1}}
    {{execAdmin "mute" .User.ID (randInt 61) "you fuckin' donut, you modded yourself"}}
{{else if le $random 2}}
    {{execAdmin "kick" .User.ID (print "you fuckin' donut, you modded yourself :/ rejoin at [" .Server.Name "](" $serverLink ")")}}
{{else if le $random 3}}
    {{execAdmin "ban" .User.ID (print "you fuckin' donut, you modded yourself :/ rejoin at [" .Server.Name "](" $serverLink ")") "-d" (print (randInt 61) "mins")}}
{{end}}