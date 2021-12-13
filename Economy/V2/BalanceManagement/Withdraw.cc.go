{{/*
        Made by Ranger (765316548516380732)

    Trigger Type: `Regex`
    Trigger: `\A(-|<@!?204255221017214977>\s*)(withdraw|with|extract|unload)(\s+|\z)`

    ©️ Ranger 2020-Present
    GNU, GPLV3 License
    Repository: https://github.com/Ranger-4297/YAGPDB-ccs
*/}}

{{/* Only edit below if you know what you're doing (: rawr */}}

{{/* Initiates variables */}}
{{$userID := .User.ID}}
{{$successColor := 0x00ff7b}}
{{$prefix := index (reFindAllSubmatches `.*?: \x60(.*)\x60\z` (execAdmin "Prefix")) 0 1 }}
{{$errorColor := 0xFF0000}}

{{/* Withdraws money */}}

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
    {{with $.CmdArgs}}
        {{if not (dbGet $userID "EconomyInfo")}}
            {{dbSet $userID "EconomyInfo" (sdict "cash" 200 "bank" 0)}}
        {{end}}
        {{with (dbGet $userID "EconomyInfo")}}
            {{$a = sdict .Value}}
            {{$cash := (toInt $a.cash)}}
            {{$bank := (toInt $a.bank)}}
            {{$amount := (index $.CmdArgs 0)}}
            {{if (toInt $amount)}}
                {{if gt (toInt $amount) (toInt $bank)}}
                    {{sendMessage nil (cembed
                            "author" (sdict "name" $.User.Username "icon_url" ($.User.AvatarURL "128"))
                            "description" (print "You're unable to withdraw more than you have in your bank.\nYou currently have " $symbol $bank " in your bank.")
                            "color" $errorColor
                            "timestamp" currentTime
                            )}}
                {{else}}
                    {{$newCashBalance := $amount | add $cash}}
                    {{$newBankBalance := $amount | sub $bank}}
                    {{sendMessage nil (cembed
                            "author" (sdict "name" $.User.Username "icon_url" ($.User.AvatarURL "128"))
                            "description" (print "You withdrew " $symbol $amount " from your bank!")
                            "color" $successColor
                            "timestamp" currentTime
                            )}}
                    {{dbSet $userID "EconomyInfo" (sdict "cash" $newCashBalance "bank" $newBankBalance)}}
                {{end}}
            {{else if eq (lower (toString $amount)) "all"}}
                {{$newCashBalance := $bank | add $cash}}
                {{$newBankBalance := (toInt 0)}}
                {{sendMessage nil (cembed
                            "author" (sdict "name" $.User.Username "icon_url" ($.User.AvatarURL "128"))
                            "description" (print "You withdrew " $symbol $bank " from your bank!")
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