{{/*
        Made by Ranger (765316548516380732)

    Trigger Type: `Command`
    Trigger: `AddMoney`
©️ Ranger 2021
MIT License
*/}}

{{$args := parseArgs 0 "" (carg "member" "Member giving money to") (carg "string" "Cash or bank") (carg "int" "money adding")}}
{{$errorColor := 0xFF0000}}
{{$successColor := 0x00ff8b}}
{{$prefix := index (reFindAllSubmatches `.*?: \x60(.*)\x60\z` (execAdmin "Prefix")) 0 1 }}
{{$Receiving := ""}}
{{$receivingUser := ""}}
{{$b := .User.ID}}
{{if not ($args.Get 0)}}
    {{$errorEmbed := (cembed
            "description" (print "No user argument passed.\nSyntax is: `" $prefix "giveMoney <Member:Mention/ID> <Destination:Cash/Bank> <Amount:Amount>`")
            "color" $errorColor
            )}}
    {{sendMessage nil $errorEmbed}}
{{else if ($args.Get 0).User}}
    {{$Receiving = ($args.Get 0).User}}
    {{$receivingUser = $Receiving}}
    {{$Receiving = $Receiving.ID}}
{{else}}
    {{$errorEmbed := (cembed
            "description" (print "Invalid user argument passed.\nSyntax is: `" $prefix "GiveMoney <Member:Mention/ID> <Destination:Cash/Bank> <Amount:Amount>`")
            "color" $errorColor
            )}}
    {{sendMessage nil $errorEmbed}}
{{end}}

{{with and $Receiving}}
    {{if not ($args.Get 1)}}
        {{$errorEmbed := (cembed
            "description" (print "No destination argument passed.\nSyntax is: `" $prefix "GiveMoney <Member:Mention/ID> <Destination:Cash/Bank> <Amount:Amount>`")
            "color" $errorColor
        )}}
        {{sendMessage nil $errorEmbed}}
    {{else if ($args.Get 1)}}
        {{$moneyDestination := (lower ($args.Get 1))}}
        {{if eq $moneyDestination  "cash" "bank"}}
            {{with (dbGet 0 "EconomySettings")}}
                {{$a := sdict .Value}}
                {{$symbol := $a.symbol}}
                {{with (dbGet $Receiving "EconomyInfo")}}
                    {{$a = sdict .Value}}
                    {{$receivingMoney := ""}}
                    {{if eq $moneyDestination "cash"}}
                        {{$receivingMoney = $a.cash}}
                    {{else if eq $moneyDestination "bank"}}
                        {{$receivingMoney = $a.bank}}
                    {{end}}
                    {{if not ($args.Get 2)}}
                        {{$errorEmbed := (cembed
                            "description" (print "No amount argument passed.\nSyntax is: `" $prefix "GiveMoney <Member:Mention/ID> <Destination:Cash/Bank> <Amount:Amount>`")
                            "color" $errorColor
                            )}}
                        {{sendMessage nil $errorEmbed}}
                    {{else}}
                        {{$amount := ($args.Get 2)}}
                        {{if (toInt $amount)}}
                            {{with (dbGet $b "EconomyInfo")}}
                                {{$a = sdict .Value}}
                                {{$yourMoney := $a.cash}}
                                {{if gt $amount $yourMoney}}
                                    {{$errorEmbed := (cembed
                                            "description" (print "You do not have enough to give this. You currently have " $symbol $yourMoney " on hand.")
                                            "color" $errorColor
                                            )}}
                                {{else}}
                                    {{$yourNewAmount := (sub $yourMoney $amount)}}
                                    {{$newAmount := (add $receivingMoney $amount)}}
                                    {{$sdict := (dbGet $Receiving "EconomyInfo").Value}}
                                    {{$sdict.Set $moneyDestination $newAmount}}
                                    {{dbSet $Receiving "EconomyInfo" $sdict}}
                                    {{$giveMoneyEmbed := (cembed
                                            "author" (sdict "name" $.User.Username "icon_url" ($.User.AvatarURL "128"))
                                            "description" (print "You gave " $symbol $amount " to <@!" $Receiving ">'s " $moneyDestination " They now have " $symbol $newAmount "!\nYou lost " $symbol $amount " You now have " $symbol $yourNewAmount " in cash!")
                                            "color" $successColor
                                            "timestamp" currentTime
                                            )}}
                                    {{sendMessage nil $giveMoneyEmbed}}
                                    {{$sdict = (dbGet $b "EconomyInfo").Value}}
                                    {{$sdict.Set "cash" $yourNewAmount}}
                                    {{dbSet $b "EconomyInfo" $sdict}}
                                {{end}}
                            {{end}}
                        {{else}}
                            {{$errorEmbed := (cembed
                            "author" (sdict "name" $.User.Username "icon_url" ($.User.AvatarURL "128"))
                            "description" (print "You're unable to add this value, check that you used a valid number above 1")
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