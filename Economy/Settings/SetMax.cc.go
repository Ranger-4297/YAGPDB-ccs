{{/*
        Made by Ranger (765316548516380732)

    Trigger Type: `Command`
    Trigger: `setMax`
©️ Ranger 2021
MIT License
*/}}

{{/* Only edit below if you know what you're doing (: rawr */}}

{{$args := parseArgs 1 "setMax <value>" (carg "int" "Value")}}
{{$newMax := $args.Get 0}}

{{$EconomySymbol := ""}}
{{$min := ""}}
{{$a := ""}}
{{with (dbGet 0 "EconomySettings")}}
    {{$a = sdict .Value}}
    {{$min = $a.min}}
    {{$EconomySymbol = $a.EconomySymbol}}
    {{if lt $newMax $min}}
        {{$errorEmbed := (cembed
            "author" (sdict "name" $.User.Username "icon_url" ($.User.AvatarURL "128"))
            "description" (print "You cannot set `max` to a value below `min`.\nYour min is set to `" $EconomySymbol $min "`")
            "color" 0x00ff8b
            "timestamp" currentTime
            )}}
        {{sendMessage nil $errorEmbed}}
    {{else}}
        {{$updateEmbed := (cembed
            "author" (sdict "name" $.User.Username "icon_url" ($.User.AvatarURL "128"))
            "description" (print "Successfully set `max` to  " $EconomySymbol $newMax)
            "color" 0x00ff8b
            "timestamp" currentTime
            )}}
        {{sendMessage nil $updateEmbed}}
        {{$sdict := (dbGet 0 "EconomySettings").Value}}
        {{$sdict.Set "max" $newMax}}
        {{dbSet 0 "EconomySettings" $sdict}}
    {{end}}
{{end}}