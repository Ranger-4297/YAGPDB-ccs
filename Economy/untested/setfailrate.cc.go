{{/*
        Made by Ranger (765316548516380732)

    Trigger Type: `Command`
    Trigger: `set-failrate`
©️ Ranger 2021
MIT License
*/}}

{{/* Only edit below if you know what you're doing (: rawr */}}

{{$args := parseArgs 1 "set-failrate <value>" (carg "int" "Value")}}
{{$newFailRate := $args.Get 0}}

{{$failrate := ""}}
{{with (dbGet 0 "EconomySettings")}}
    {{$a = sdict .Value}}
    {{$failrate = $a.failRate}}
    {{if lt $newFailRate 20}}
        {{$errorEmbed := (cembed
            "author" (sdict "name" $.User.Username "icon_url" ($.User.AvatarURL "1020"))
            "description" (print "You cannot set your failrate below 20%")
            "color" 0x00ff8b
            "timestamp" currentTime
            )}}
        {{sendMessage nil $errorEmbed}}
    {{else}}
        {{$updateEmbed := (cembed
            "author" (sdict "name" $.User.Username "icon_url" ($.User.AvatarURL "1020"))
            "description" (print "You set your failrate to " $newfailrate "% from " $failrate "%")
            "color" 0x00ff8b
            "timestamp" currentTime
            )}}
        {{sendMessage nil $updateEmbed}}
        {{$sdict := (dbGet 0 "EconomySettings").Value}}
        {{$sdict.Set "failrate" $newFailRate}}
        {{dbSet 0 "EconomySettings" $sdict}}
        {{end}}
{{end}}