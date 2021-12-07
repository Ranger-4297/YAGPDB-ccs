{{/*
        Made by Ranger (765316548516380732)

    Trigger Type: `Command`
    Trigger: `setStartBalance`
©️ Ranger 2021
MIT License
*/}}

{{/* Only edit below if you know what you're doing (: rawr */}}

{{$args := parseArgs 1 "setStartBalance Symbol" (carg "string" "Value")}}
{{$newStartBalance := (toInt $args.Get 0)}}
{{$startBalance := ""}}
{{$economySymbol := ""}}
{{$a := ""}}
{{with (dbGet 0 "EconomySettings")}}
    {{$a = sdict .Value}}
    {{$startBalance = $a.startBalance}}
    {{$economySymbol = $a.economySymbol}}
    {{$updateEmbed := (cembed
            "author" (sdict "name" $.User.Username "icon_url" ($.User.AvatarURL "128"))
            "description" (print "You set the server start-balance to " $newStartBalance " from " $startBalance)
            "color" 0x00ff8b
            "timestamp" currentTime
            )}}
    {{sendMessage nil $updateEmbed}}
    {{$sdict := (dbGet 0 "EconomySettings").Value}}
    {{$sdict.Set "startBalance" $newStartBalance}}
    {{dbSet 0 "EconomySettings" $sdict}}
{{end}}
