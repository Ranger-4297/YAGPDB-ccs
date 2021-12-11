{{/*
        Made by Ranger (765316548516380732)

    Trigger Type: `Command`
    Trigger: `AddMoney`
©️ Ranger 2021
MIT License
*/}}

{{$args := parseArgs 0 "" (carg "member" "Member adding money to") (carg "string" "Cash or bank") (carg "int" "money adding")}}
{{$errorColor := 0xFF0000}}
{{$successColor := 0x00ff8b}}
{{$prefix := index (reFindAllSubmatches `.*?: \x60(.*)\x60\z` (execAdmin "Prefix")) 0 1 }}
{{$Receiving := ""}}
{{$ReceivingUser := ""}}
{{if not ($args.Get 0)}}
    {{$errorEmbed := (cembed
            "description" (print "No user argument passed.\nSyntax is: `" $prefix "AddMoney <Member:Mention/ID> <Destination:Cash/Bank> <Amount:Amount>`")
            "color" $errorColor
            )}}
    {{sendMessage nil $errorEmbed}}
{{else if ($args.Get 0).User}}
    {{$Receiving = ($args.Get 0).User}}
    {{$Receiving = $Receiving.ID}}
{{else}}
    {{$errorEmbed := (cembed
            "description" (print "Invalid user argument passed.\nSyntax is: `" $prefix "AddMoney <Member:Mention/ID> <Destination:Cash/Bank> <Amount:Amount>`")
            "color" $errorColor
            )}}
{{end}}

{{with and $Receiving}}
    {{if not ($args.Get 1)}}
        {{$errorEmbed := (cembed
            "description" (print "No destination argument passed.\nSyntax is: `" $prefix "AddMoney <Member:Mention/ID> <Destination:Cash/Bank> <Amount:Amount>`")
            "color" $errorColor
        )}}
        {{sendMessage nil $errorEmbed}}
    {{else if ($args.Get 1)}}
        {{$MoneyDestination := (lower ($args.Get 1))}}
        {{if eq $MoneyDestination  "cash" "bank"}}
            {{with (dbGet 0 "EconomySettings")}}
                {{$a := sdict .Value}}
                {{$symbol := $a.symbol}}
                {{with (dbGet $Receiving "EconomyInfo")}}
                    {{$a = sdict .Value}}
                    {{$ReceivingMoney := ""}}
                    {{if eq $MoneyDestination "cash"}}
                        {{$ReceivingMoney = $a.cash}}
                    {{else if eq $MoneyDestination "bank"}}
                        {{$ReceivingMoney = $a.bank}}
                    {{end}}
                    {{if not ($args.Get 2)}}
                        {{$errorEmbed := (cembed
                            "description" (print "No amount argument passed.\nSyntax is: `" $prefix "AddMoney <Member:Mention/ID> <Destination:Cash/Bank> <Amount:Amount>`")
                            "color" $errorColor
                            )}}
                        {{sendMessage nil $errorEmbed}}
                    {{else}}
                        {{$amount := ($args.Get 2)}}
                        {{if (toInt $amount)}}
                            {{$newAmount := (add $ReceivingMoney $amount)}}
                            {{$addMoneyEmbed := (cembed
                            "author" (sdict "name" $.User.Username "icon_url" ($.User.AvatarURL "128"))
                            "description" (print "You added " $symbol $amount " to <@!" $Receiving ">'s " $MoneyDestination "\nThey now have " $symbol $newAmount "!")
                            "color" $successColor
                            "timestamp" currentTime
                            )}}
                            {{sendMessage nil $addMoneyEmbed}}
                            {{$sdict := (dbGet $Receiving "EconomyInfo").Value}}
                            {{$sdict.Set $MoneyDestination $newAmount}}
                            {{dbSet $Receiving "EconomyInfo" $sdict}}
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