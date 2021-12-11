{{/*
        Made by Ranger (765316548516380732)

    Trigger Type: `Command`
    Trigger: `RemoveMoney
©️ Ranger 2021
MIT License
*/}}

{{$args := parseArgs 0 "" (carg "member" "Member taking money from") (carg "string" "Cash or bank") (carg "int" "money removing")}}
{{$errorColor := 0xFF0000}}
{{$successColor := 0x00ff8b}}
{{$prefix := index (reFindAllSubmatches `.*?: \x60(.*)\x60\z` (execAdmin "Prefix")) 0 1 }}
{{$Removing := ""}}
{{$RemovingUser := ""}}
{{if not ($args.Get 0)}}
    {{$errorEmbed := (cembed
            "description" (print "No user argument passed.\nSyntax is: `" $prefix "RemoveMoney <Member:Mention/ID> <Destination:Cash/Bank> <Amount:Amount>`")
            "color" $errorColor
            )}}
    {{sendMessage nil $errorEmbed}}
{{else if ($args.Get 0).User}}
    {{$Removing = ($args.Get 0).User}}
    {{$RemovingUser = $Removing}}
    {{$Removing = $Removing.ID}}
{{else}}
    {{$errorEmbed := (cembed
            "description" (print "Invalid user argument passed.\nSyntax is: `" $prefix "RemoveMoney <Member:Mention/ID> <Destination:Cash/Bank> <Amount:Amount>`")
            "color" $errorColor
            )}}
{{end}}

{{with and $Removing}}
    {{if not ($args.Get 1)}}
        {{$errorEmbed := (cembed
            "description" (print "No destination argument passed.\nSyntax is: `" $prefix "RemoveMoney <Member:Mention/ID> <Destination:Cash/Bank> <Amount:Amount>`")
            "color" $errorColor
        )}}
        {{sendMessage nil $errorEmbed}}
    {{else if ($args.Get 1)}}
        {{$MoneyDestination := (lower ($args.Get 1))}}
        {{if eq $MoneyDestination  "cash" "bank"}}
            {{with (dbGet 0 "EconomySettings")}}
                {{$a := sdict .Value}}
                {{$symbol := $a.symbol}}
                {{with (dbGet $Removing "EconomyInfo")}}
                    {{$a = sdict .Value}}
                    {{$RemovingMoney := ""}}
                    {{if eq $MoneyDestination "cash"}}
                        {{$RemovingMoney = $a.cash}}
                    {{else if eq $MoneyDestination "bank"}}
                        {{$RemovingMoney = $a.bank}}
                    {{end}}
                    {{if not ($args.Get 2)}}
                        {{$errorEmbed := (cembed
                            "description" (print "No amount argument passed.\nSyntax is: `" $prefix "RemoveMoney <Member:Mention/ID> <Destination:Cash/Bank> <Amount:Amount>`")
                            "color" $errorColor
                            )}}
                        {{sendMessage nil $errorEmbed}}
                    {{else}}
                        {{$amount := ($args.Get 2)}}
                        {{if (toInt $amount)}}
                            {{$newAmount := (sub $RemovingMoney $amount)}}
                            {{$addMoneyEmbed := (cembed
                            "author" (sdict "name" $.User.Username "icon_url" ($.User.AvatarURL "128"))
                            "description" (print "You removed " $symbol $amount " from <@!" $Removing ">'s " $MoneyDestination "\nThey now have " $symbol $newAmount "!")
                            "color" $successColor
                            "timestamp" currentTime
                            )}}
                            {{sendMessage nil $addMoneyEmbed}}
                            {{$sdict := (dbGet $Removing "EconomyInfo").Value}}
                            {{$sdict.Set $MoneyDestination $newAmount}}
                            {{dbSet $Removing "EconomyInfo" $sdict}}
                        {{else}}
                            {{$errorEmbed := (cembed
                            "author" (sdict "name" $.User.Username "icon_url" ($.User.AvatarURL "128"))
                            "description" (print "You're unable to remove this value, check that you used a valid number above 1")
                            "color" $errorColor
                            "timestamp" currentTime
                            )}}
                            {{sendMessage nil $errorEmbed}}
                        {{end}}
                    {{end}}
                {{end}}
            {{end}}
        {{else}}
            {{$errorEmbed := (cembed
            "description" (print "No destination argument passed.\nSyntax is: `" $prefix "AddMoney <Member:Mention/ID> <Destination:Cash/Bank> <Amount:Amount>`")
            "color" $errorColor
            )}}
        {{sendMessage nil $errorEmbed}}
        {{end}}
    {{end}}
{{end}}