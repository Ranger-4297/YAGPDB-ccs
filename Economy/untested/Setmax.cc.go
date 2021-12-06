{{/*
        Made by Ranger (765316548516380732)

    Trigger Type: `Command`
    Trigger: `setMax`
©️ Ranger 2021
MIT License
*/}}

{{/* Only edit below if you know what you're doing (: rawr */}}

{{$args := parseArgs 1 "setMax <value>" (carg "int" "Value")}}
{{$max := $args.Get 0}}

{{$min := ""}}
{{with (dbGet 0 "EconomySettings")}}
    {{$a = sdict .Value}}
    {{$min = $a.min}}
    {{if lt $max $min}}
        {{$errorEmbed := (cembed
            "author" (sdict "name" $.User.Username "icon_url" ($.User.AvatarURL "1020"))
            "description" (print "You cannot set `max` to a value below `min`.\nYour min is set to `£" $min "`")
            "color" 0x00ff8b
            "timestamp" currentTime
            )}}
        {{sendMessage nil $errorEmbed}}
    {{else}}
        {{$updateEmbed := (cembed
            "author" (sdict "name" $.User.Username "icon_url" ($.User.AvatarURL "1020"))
            "description" (print "Successfully set `min` to  " $min)
            "color" 0x00ff8b
            "timestamp" currentTime
            )}}
        {{sendMessage nil $errorEmbed}}
        {{$sdict := (dbGet 0 "EconomySettings").Value}}
        {{$sdict.Set "max" $max}}
        {{dbSet 0 "EconomySettings" $sdict}}
    {{end}}
{{end}}
