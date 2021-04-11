{{/*
        Made by Ranger (765316548516380732)

    Trigger Type: `Command`
    Trigger: `Modmyself`
©️ Dynamic 2021
MIT License
*/}}

{{/* Configuration values start */}}
{{$serverLink := "https://discord.gg/ekMQH384KC"}} {{/* Your servers invite link */}}
{{/* Configuration values end */}}

{{/* Only edit below if you know what you're doing except for the 21st line. Add your staff application link at the end. (: rawr */}}

{{$random := randInt 4}}
{{if le $random 0}}
    {{execAdmin "warn" .User.ID "you fookin donut, you modded yourself"}}
    {{else if le $random 1}}
        {{execAdmin "mute" .User.ID (randInt 61) "you fookin donut, you modded yourself"}}
        {{else if le $random 2}}
            {{execAdmin "kick" .User.ID (print "you fookin donut, you modded yourself :/ rejoin at [" .Server.Name "](" $serverLink ")")}}
            {{else if le $random 3}}
                {{execAdmin "ban" .User.ID (print "you fookin donut, you modded yourself\n:/rejoin when your ban has timedout. (Yes. You have to wait) at [" .Server.Name "](" $serverLink ")") "-d" (print (randInt 61) "mins")}}
        {{end}}
    {{end}}
{{end}}
