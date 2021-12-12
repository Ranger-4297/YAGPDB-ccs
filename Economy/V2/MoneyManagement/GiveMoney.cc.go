{{/*
		Made by Ranger (765316548516380732)

	Trigger Type: `Regex`
	Trigger: `\A(-|<@!?204255221017214977>\s*)(give-?money|loan-?money|money-?give|money-?loan)(\s+|\z)`
©️ Ranger 2021
MIT License
*/}}

{{/* Only edit below if you know what you're doing (: rawr */}}

{{/* Initiates variables */}}
{{$userID := .User.ID}}
{{$successColor := 0x00ff7b}}
{{$errorColor := 0xFF0000}}
{{$prefix := index (reFindAllSubmatches `.*?: \x60(.*)\x60\z` (execAdmin "Prefix")) 0 1 }}

{{/* Gives money from your balance to the given user */}}

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
        {{if index . 0}}
            {{if index . 0 | getMember}}
                {{$user := getMember (index . 0)}}
                {{$receivingUser := $user.User.ID}}
                {{if not (dbGet $receivingUser "EconomyInfo")}}
                    {{dbSet $receivingUser "EconomyInfo" (sdict "cash" 200 "bank" 0)}}
                {{end}}
                {{if gt (len $.CmdArgs) 1}}
                    {{$amount := (index $.CmdArgs 1)}}
                    {{if (toInt $amount)}}
                        {{with (dbGet $receivingUser "EconomyInfo")}}
                            {{$a = sdict .Value}}
                            {{$receivingBalance := $a.cash}}
                            {{if not (dbGet $userID "EconomyInfo")}}
                                {{dbSet $receivingUser "EconomyInfo" (sdict "cash" 200 "bank" 0)}}
                            {{end}}
                            {{with (dbGet $userID "EconomyInfo")}}
                                {{$a = sdict .Value}}
                                {{$yourBalance := $a.cash}}
                                {{if gt (toInt $amount) (toInt $yourBalance)}}
                                    {{sendMessage nil (cembed
                                            "author" (sdict "name" $.User.Username "icon_url" ($.User.AvatarURL "128"))
                                            "description" (print "You cannot give more than you have.")
                                            "color" $errorColor
                                            "timestamp" currentTime
                                            )}}
                                {{else}}
                                    {{$receivingNewBalance := $receivingBalance | add $amount}}
                                    {{$yourNewBalance := $amount |sub $yourBalance}}
                                    {{sendMessage nil (cembed
                                            "author" (sdict "name" $.User.Username "icon_url" ($.User.AvatarURL "128"))
                                            "description" (print "You gave " $symbol $amount " to <@!" $receivingUser ">\nThey now have " $symbol $receivingNewBalance " in cash!")
                                            "color" $successColor
                                            "timestamp" currentTime
                                            )}}
                                    {{$sdict := (dbGet $receivingUser "EconomyInfo").Value}}
                                    {{$sdict.Set "cash" $receivingNewBalance}}
                                    {{dbSet $receivingUser "EconomyInfo" $sdict}}
                                {{end}}
                            {{end}}
                        {{end}}
                    {{else}}
                        {{sendMessage nil (cembed
                                    "author" (sdict "name" $.User.Username "icon_url" ($.User.AvatarURL "128"))
                                    "description" (print "No `Amount` argument passed.\nSyntax is: `" $.Cmd " <Member:Mention/ID> <Amount:Amount>`")
                                    "color" $errorColor
                                    "timestamp" currentTime
                                    )}}
                    {{end}}
                {{else}}
                    {{sendMessage nil (cembed
                            "author" (sdict "name" $.User.Username "icon_url" ($.User.AvatarURL "128"))
                            "description" (print "You're unable to give this value, check that you used a valid number above 1")
                            "color" $errorColor
                            "timestamp" currentTime
                            )}}
                {{end}}
            {{else}}
                {{sendMessage nil (cembed
                            "author" (sdict "name" $.User.Username "icon_url" ($.User.AvatarURL "128"))
                            "description" (print "Invalid `Amount` argument provided.\nSyntax is `" $.Cmd " <User:Mention/ID> <Amount:Amount>`")
                            "color" $errorColor
                            "timestamp" currentTime
                            )}}
            {{end}}
        {{else}}
            {{sendMessage nil (cembed
            "author" (sdict "name" $.User.Username "icon_url" ($.User.AvatarURL "128"))
            "description" (print "Invalid `User` argument provided.\nSyntax is `" $.Cmd " <User:Mention/ID> <Amount:Amount>`")
            "color" $errorColor
            "timestamp" currentTime
            )}}
        {{end}}
    {{else}}
        {{sendMessage nil (cembed
            "author" (sdict "name" $.User.Username "icon_url" ($.User.AvatarURL "128"))
            "description" (print "No `User` argument provided.\nSyntax is `" $.Cmd " <User:Mention/ID> <Amount:Amount>`")
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