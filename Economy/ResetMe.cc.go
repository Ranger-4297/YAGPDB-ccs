{{/*
        Made by Ranger (765316548516380732)

    Trigger Type: `Command`
    Trigger: `ResetMe`
©️ Ranger 2021
MIT License
*/}}

{{/* Only edit below if you know what you're doing (: rawr */}}

{{$depositEmbed := (cembed
        "author" (sdict "name" $.User.Username "icon_url" ($.User.AvatarURL "1024"))
        "description" (print "You reset your balance")
        "color" 0x00ff7b
        "timestamp" currentTime
        )}}
{{sendMessage nil $depositEmbed}}
{{$sdict := (dbGet .User.ID "EconomyInfo").Value}}
{{$sdict.Set "bank" (toInt "0")}}
{{dbSet .User.ID "EconomyInfo" $sdict}}
{{$sdict.Set "cash" (toInt "0")}}
{{dbSet .User.ID "EconomyInfo" $sdict}}