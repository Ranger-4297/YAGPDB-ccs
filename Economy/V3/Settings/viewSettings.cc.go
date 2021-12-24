{{/*
        Made by Ranger (765316548516380732)

    Trigger Type: `Regex`
    Trigger: `\A(-|<@!?204255221017214977>\s*)(view-?settings?|settingview)(\s+|\z)`

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

{{/*
If there is no setting values
You'll be asked to set it up with default values
You can change these later
*/}}

{{/* Response */}}
{{$embed := sdict}}
{{$embed.Set "author" (sdict "name" $.User.Username "icon_url" ($.User.AvatarURL "1024"))}}
{{$embed.Set "timestamp" currentTime}}
{{with (dbGet 0 "EconomySettings")}}
    {{$a := sdict .Value}}
    {{$min := $a.min}}
    {{$max := $a.max}}
    {{$symbol := $a.symbol}}
    {{$crimeCooldown := $a.crimeCooldown}}
    {{$robCooldown := $a.robCooldown}}
    {{$startBalance := $a.startBalance}}
    {{if (reFind `(<a?:[A-z+]+\:\d{17,19}>)` $symbol)}}
        {{$symbol = $symbol}}
    {{else}}
        {{$symbol = (print "`" $symbol "`")}}
    {{end}}
    {{$embed.Set "description" (print "Min: `" $min "`\nMax: `" $max "`\nSymbol: " $symbol "\nstartBalance: " $symbol $startBalance "\ncrimeCooldown: " $crimeCooldown "\nrobCooldown: " $robCooldown)}}
    {{$embed.Set "color" $successColor}}
{{else}}
    {{$embed.Set "description" (print "No `Settings` database found.\nPlease set it up with the default values using `" $prefix "set default`")}}
    {{$embed.Set "color" $errorColor}}
{{end}}
{{sendMessage nil (cembed $embed)}}