{{/*
        Made by Ranger (765316548516380732)

    Trigger Type: `Command`
    Trigger: `setSymbol`
©️ Ranger 2021
MIT License
*/}}

{{/* Only edit below if you know what you're doing (: rawr */}}

{{$args := parseArgs 1 "setSymbol Symbol" (carg "string" "Value")}}
{{$newSymbol := $args.Get 0}}
{{$EconomySymbol := ""}}
{{$a := ""}}
{{with (dbGet 0 "EconomySettings")}}
    {{$a = sdict .Value}}
    {{$EconomySymbol = $a.EconomySymbol}}
    {{$updateEmbed := (cembed
            "author" (sdict "name" $.User.Username "icon_url" ($.User.AvatarURL "128"))
            "description" (print "You set the server currency symbol to " $newSymbol " from " $EconomySymbol)
            "color" 0x00ff8b
            "timestamp" currentTime
            )}}
    {{sendMessage nil $updateEmbed}}
    {{$sdict := (dbGet 0 "EconomySettings").Value}}
    {{$sdict.Set "EconomySymbol" $newSymbol}}
    {{dbSet 0 "EconomySettings" $sdict}}
{{end}}