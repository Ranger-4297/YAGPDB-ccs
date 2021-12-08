{{/*
        Made by Ranger (765316548516380732)

    Trigger Type: `Command`
    Trigger: `viewSettings`
©️ Ranger 2021
MIT License
*/}}

{{$a := ""}}
{{$min := ""}}
{{$max := ""}}
{{$failRate := ""}}
{{$symbol := ""}}
{{$startBalance := ""}}
{{$embedColor := 0x00ff8b}}
{{$settingsEmbed := (cembed
            "author" (sdict "name" $.User.Username "icon_url" ($.User.AvatarURL "128"))
            "description" (print "No settings data found")
            "timestamp" currentTime
            "color" $embedColor
            )}}
{{with (dbGet 0 "EconomySettings")}}
    {{$a = sdict .Value}}
    {{$min = $a.min}}
    {{$max = $a.max}}
    {{$failRate = $a.failRate}}
    {{$symbol = $a.symbol}}
    {{$startBalance = $a.startBalance}}
    {{$settingsEmbed = (cembed
            "author" (sdict "name" $.User.Username "icon_url" ($.User.AvatarURL "128"))
            "description" (print "Min: `" $min "`\nMax: `" $max "`\nfailRate: `" $failRate "`\nEconomy symbol: `" $symbol "`\nstartBalance: `" $symbol $startBalance "`")
            "timestamp" currentTime
            "color" $embedColor
            )}}
{{end}}
{{sendMessage nil $settingsEmbed}}