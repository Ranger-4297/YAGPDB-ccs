{{/*
        Made by Ranger (765316548516380732)

    Trigger Type: `Regex`
    Trigger: `\A(-|<@!?204255221017214977>\s*)(view-?settings?)(\s+|\z)`

    ©️ Ranger 2020-Present
    GNU, GPLV3 License
    Repository: https://github.com/Ranger-4297/YAGPDB-ccs
*/}}


{{/* Only edit below if you know what you're doing (: rawr */}}

{{/* Initiates variables */}}
{{$successColor := 0x00ff7b}}
{{$errorColor := 0xFF0000}}
{{$prefix := index (reFindAllSubmatches `.*?: \x60(.*)\x60\z` (execAdmin "Prefix")) 0 1 }}

{{/* Configures economy settings */}}

{{/* Response */}}
{{$embed := sdict}}
{{$embed.Set "author" (sdict "name" $.User.Username "icon_url" ($.User.AvatarURL "1024"))}}
{{$embed.Set "timestamp" currentTime}}
{{with (dbGet 0 "EconomySettings")}}
    {{$a := sdict .Value}}
    {{$min := (humanizeThousands $a.min)}}
    {{$max := (humanizeThousands $a.max)}}
    {{$symbol := $a.symbol}}
    {{$workCooldown := humanizeDurationSeconds (mult $.TimeSecond $a.workCooldown | toDuration)}}
    {{$crimeCooldown := humanizeDurationSeconds (mult $.TimeSecond $a.crimeCooldown | toDuration)}}
    {{$robCooldown := humanizeDurationSeconds (mult $.TimeSecond $a.robCooldown | toDuration)}}
    {{$startBalance := (humanizeThousands $a.startBalance)}}
    {{if (reFind `(<a?:[A-z+]+\:\d{17,19}>)` $symbol)}}
        {{$symbol = $symbol}}
    {{else}}
        {{$symbol = (print $symbol)}}
    {{end}}
    {{$embed.Set "description" (print "Min: `" $min "`\nMax: `" $max "`\nSymbol: " $symbol "\nstartBalance: `" $startBalance "`\nworkCooldown: `" $workCooldown "`\ncrimeCooldown: `" $crimeCooldown "`\nrobCooldown: `" $robCooldown "`")}}
    {{$embed.Set "color" $successColor}}
{{else}}
    {{$embed.Set "description" (print "No `Settings` database found.\nPlease set it up with the default values using `" $prefix "set default`")}}
    {{$embed.Set "color" $errorColor}}
{{end}}
{{sendMessage nil (cembed $embed)}}