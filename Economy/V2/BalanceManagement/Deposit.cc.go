{{/*
        Made by Ranger (765316548516380732)

    Trigger Type: `Regex`
    Trigger: `\A(-|<@!?204255221017214977>\s*)(deposit|dep|pledge|load)(\s+|\z)`
©️ Ranger 2021
MIT License
*/}}

{{/* Only edit below if you know what you're doing (: rawr */}}

{{/* Initiates variables */}}
{{$userID := .User.ID}}
{{$successColor := 0x00ff7b}}
{{$prefix := index (reFindAllSubmatches `.*?: \x60(.*)\x60\z` (execAdmin "Prefix")) 0 1 }}
{{$errorColor := 0xFF0000}}

{{/* Deposits money */}}

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
    {{with .CmdArgs}}
        {{if not (dbGet $userID "EconomyInfo")}}
            {{dbSet $userID "EconomyInfo" (sdict "cash" 200 "bank" 0)}}
        {{end}}
        {{with (dbGet $userID "EconomyInfo")}}
            {{$a = sdict .Value}}
            {{$cash := (toInt $a.cash)}}
            {{$bank := (toInt $a.bank)}}
            {{$amount := (index $.CmdArgs 0)}}
            {{if (toInt $amount)}}
                {{if gt (toInt $amount) (toInt $cash)}}
                    {{sendMessage nil (cembed
                            "author" (sdict "name" $.User.Username "icon_url" ($.User.AvatarURL "128"))
                            "description" (print "You're unable to deposit more than you have on hand.\nYou currently have " $symbol $cash " on you.")
                            "color" $errorColor
                            "timestamp" currentTime
                            )}}
                {{else}}
                    {{$newCashBalance := $amount | sub $cash}}
                    {{$newBankBalance := $amount | add $bank}}
                    {{sendMessage nil (cembed
                            "author" (sdict "name" $.User.Username "icon_url" ($.User.AvatarURL "128"))
                            "description" (print "You deposited " $symbol $amount " into your bank!")
                            "color" $successColor
                            "timestamp" currentTime
                            )}}
                    {{dbSet $userID "EconomyInfo" (sdict "cash" $newCashBalance "bank" $newBankBalance)}}
                {{end}}
            {{else if eq (lower (toString $amount)) "all"}}
                {{$newCashBalance := (toInt 0)}}
                {{$newBankBalance := $bank | add $cash}}
                {{sendMessage nil (cembed
                            "author" (sdict "name" $.User.Username "icon_url" ($.User.AvatarURL "128"))
                            "description" (print "You deposited " $symbol $cash " into your bank!")
                            "color" $successColor
                            "timestamp" currentTime
                            )}}
                {{dbSet $userID "EconomyInfo" (sdict "cash" $newCashBalance "bank" $newBankBalance)}}
            {{else}}
                {{sendMessage nil (cembed
                            "author" (sdict "name" $.User.Username "icon_url" ($.User.AvatarURL "128"))
                            "description" (print "Invalid amount.\nCommand syntax is `" $.Cmd " <Amount:Amount>`")
                            "color" $errorColor
                            "timestamp" currentTime
                            )}}
            {{end}}
        {{end}}
    {{else}}
        {{sendMessage nil (cembed
            "author" (sdict "name" $.User.Username "icon_url" ($.User.AvatarURL "128"))
            "description" (print "Invalid amount.\nCommand syntax is `" .Cmd " <Amount:Amount>`")
            "color" $errorColor
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