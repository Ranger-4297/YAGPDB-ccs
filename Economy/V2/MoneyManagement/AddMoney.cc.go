{{/*
		Made by Ranger (765316548516380732)

	Trigger Type: `Regex`
	Trigger: `\A(-|<@!?204255221017214977>\s*)(add-?money|inc-?money|money-?add)(\s+|\z)`
©️ Ranger 2021
MIT License
*/}}

{{/* Only edit below if you know what you're doing (: rawr */}}

{{/* Initiates variables */}}
{{$successColor := 0x00ff7b}}
{{$errorColor := 0xFF0000}}
{{$prefix := index (reFindAllSubmatches `.*?: \x60(.*)\x60\z` (execAdmin "Prefix")) 0 1 }}

{{/* Adds money to given user */}}

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
        {{if index . 0}}
            {{if index . 0 | getMember}}
                {{$user := getMember (index . 0)}}
                {{$receivingUser := $user.User.ID}}
                {{if not (dbGet $receivingUser "EconomyInfo")}}
                    {{dbSet $receivingUser "EconomyInfo" (sdict "cash" 200 "bank" 0)}}
                {{end}}
                {{if gt (len $.CmdArgs) 1}}
                    {{$moneyDestination := (lower (index . 1))}}
                    {{if eq $moneyDestination "cash" "bank"}}
                        {{with (dbGet $receivingUser "EconomyInfo")}}
                            {{$a = sdict .Value}}
                            {{$receivingBalance := ""}}
                            {{if eq $moneyDestination "cash"}}
                                {{$receivingBalance = $a.cash}}
                            {{else if eq $moneyDestination "bank"}}
                                {{$receivingBalance = $a.bank}}
                            {{end}}
                            {{if gt (len $.CmdArgs) 2}}
                                {{$amount := (index $.CmdArgs 2)}}
                                {{if (toInt $amount)}}
                                    {{$receivingNewBalance := $receivingBalance | add $amount}}
                                    {{sendMessage nil (cembed
                                            "author" (sdict "name" $.User.Username "icon_url" ($.User.AvatarURL "128"))
                                            "description" (print "You added " $symbol $amount " to <@!" $receivingUser ">'s " $moneyDestination "\nThey now have " $symbol $receivingNewBalance " in their " $moneyDestination "!")
                                            "color" $successColor
                                            "timestamp" currentTime
                                            )}}
                                    {{$sdict := (dbGet $receivingUser "EconomyInfo").Value}}
                                    {{$sdict.Set $moneyDestination $receivingNewBalance}}
                                    {{dbSet $receivingUser "EconomyInfo" $sdict}}
                                {{else}}
                                    {{sendMessage nil (cembed
                                            "author" (sdict "name" $.User.Username "icon_url" ($.User.AvatarURL "128"))
                                            "description" (print "You're unable to add this value, check that you used a valid number above 1")
                                            "color" $errorColor
                                            "timestamp" currentTime
                                            )}}
                                {{end}}
                            {{else}}
                                {{sendMessage nil (cembed
                                            "author" (sdict "name" $.User.Username "icon_url" ($.User.AvatarURL "128"))
                                            "description" (print "No `Amount` argument passed.\nSyntax is: `" $.Cmd " <Member:Mention/ID> <Destination:Cash/Bank> <Amount:Amount>`")
                                            "color" $errorColor
                                            "timestamp" currentTime
                                            )}}
                            {{end}}
                        {{end}}
                    {{end}}
                {{else}}
                    {{sendMessage nil (cembed
                            "author" (sdict "name" $.User.Username "icon_url" ($.User.AvatarURL "128"))
                            "description" (print "Invalid `Destination` argument provided.\nSyntax is `" $.Cmd " <User:Mention/ID> <Destination:Cash/Bank> <Amount:Amount>`")
                            "color" $errorColor
                            "timestamp" currentTime
                            )}}
                {{end}}
            {{else}}
                {{sendMessage nil (cembed
                            "author" (sdict "name" $.User.Username "icon_url" ($.User.AvatarURL "128"))
                            "description" (print "No `Destination` argument provided.\nSyntax is `" $.Cmd " <User:Mention/ID> <Destination:Cash/Bank> <Amount:Amount>`")
                            "color" $errorColor
                            "timestamp" currentTime
                            )}}
            {{end}}
        {{else}}
            {{sendMessage nil (cembed
            "author" (sdict "name" $.User.Username "icon_url" ($.User.AvatarURL "128"))
            "description" (print "Invalid `User` argument provided.\nSyntax is `" $.Cmd " <User:Mention/ID> <Destination:Cash/Bank> <Amount:Amount>`")
            "color" $errorColor
            "timestamp" currentTime
            )}}
        {{end}}
    {{else}}
        {{sendMessage nil (cembed
            "author" (sdict "name" $.User.Username "icon_url" ($.User.AvatarURL "128"))
            "description" (print "No `User` argument provided.\nSyntax is `" $.Cmd " <User:Mention/ID> <Destination:Cash/Bank> <Amount:Amount>`")
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