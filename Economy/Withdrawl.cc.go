{{/*
        Made by Ranger (765316548516380732)

    Trigger Type: `Command`
    Trigger: `Withdraw`
©️ Ranger 2021
MIT License
*/}}

{{/* Only edit below if you know what you're doing (: rawr */}}

{{$a := ""}}
{{$cash := ""}}
{{$bank := ""}}
{{$b := .User.ID}}
{{if not (dbGet $b "EconomyInfo")}}
    {{dbSet .User.ID "EconomyInfo" (sdict "cash" 0 "bank" 0)}}
{{end}}
{{with (dbGet $b "EconomyInfo")}}
    {{$a = sdict .Value}}
    {{$cash = $a.cash}}
    {{$bank = $a.bank}}
    {{$withdrawEmbed := (cembed
            "author" (sdict "name" $.User.Username "icon_url" ($.User.AvatarURL "1024"))
            "description" (print "You withdrew £" $bank " from your bank!")
            "color" 0x00ff7b
            "timestamp" currentTime
            )}}
    {{sendMessage nil $withdrawEmbed}}
    {{$sdict := (dbGet .User.ID "EconomyInfo").Value}}
    {{$sdict.Set "bank" (toInt "0")}}
    {{dbSet $b "EconomyInfo" $sdict}}
    {{$sdict.Set "cash" (add (toInt $bank) (toInt $cash))}}
    {{dbSet $b "EconomyInfo" $sdict}}
{{end}}