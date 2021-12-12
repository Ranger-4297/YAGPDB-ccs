{{/*
        Made by Ranger (765316548516380732)

    Trigger Type: `Regex`
    Trigger: `\A(-|<@!?204255221017214977>\s*)(balance|bal|wallet|money)(\s+|\z)`
©️ Ranger 2021
MIT License
*/}}

{{/* Only edit below if you know what you're doing (: rawr */}}

{{/* Initiates variables */}}
{{$user := .User}}
{{$successColor := 0x00ff7b}}
{{$errorColor := 0xFF0000}}
{{$prefix := index (reFindAllSubmatches `.*?: \x60(.*)\x60\z` (execAdmin "Prefix")) 0 1 }}
{{with .CmdArgs}}
    {{$newUser := userArg (index . 0)}}
    {{if $newUser}}
        {{$user = $newUser}}
    {{end}}
{{end}}


{{/* Output balance */}}

{{/*
If the user isn't in the economy database 
It'll automatically add them
--
If there is no setting values
You'll be asked to set it up with default values
You can change these later
*/}}

{{/* Response */}}
{{with (dbGet 0 "EconomySettings")}}
	{{$a := sdict .Value}}
	{{$symbol := $a.symbol}}
    {{if not (dbGet $user.ID "EconomyInfo")}}
        {{dbSet $user.ID "EconomyInfo" (sdict "cash" 200 "bank" 0)}}
    {{end}}
	{{with (dbGet $user.ID "EconomyInfo")}}
        {{$a := sdict .Value}}
		{{$cash := (toInt $a.cash)}}
		{{$bank := (toInt $a.bank)}}
		{{sendMessage nil (cembed
            "author" (sdict "name" $user.Username "icon_url" ($user.AvatarURL "128"))
            "description" (print $user.Mention "'s balance")
            "fields" (cslice 
                (sdict "name" "Cash" "value" (print $symbol $cash) "inline" true)
                (sdict "name" "Bank" "value" (print $symbol $bank) "inline" true)
                (sdict "name" "Networth" "value" (print $symbol (toString $cash | add $bank)) "inline" true))
            "color" $successColor
            "timestamp" currentTime
            )}}
	{{end}}
{{else}}
    {{sendMessage nil (cembed
            "author" (sdict "name" $.User.Username "icon_url" ($.User.AvatarURL "128"))
            "description" (print "No `Settings` database found.\nPlease set it up with the default values using `" $prefix "set default`")
            "color" $errorColor
            "timestamp" currentTime
            )}}
{{end}}