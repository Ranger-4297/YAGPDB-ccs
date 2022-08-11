{{/*
		Made by Ranger (765316548516380732)

	Trigger Type: `Regex`
	Trigger: `\A(-|<@!?204255221017214977>\s*)(give-?money|loan-?money|money-?give|money-?loan)(\s+|\z)`

    ©️ Ranger 2020-Present
    GNU, GPLV3 License
    Repository: https://github.com/Ranger-4297/YAGPDB-ccs
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
{{$embed := sdict}}
{{$embed.Set "author" (sdict "name" $.User.Username "icon_url" ($.User.AvatarURL "1024"))}}
{{$embed.Set "timestamp" currentTime}}
{{with (dbGet 0 "EconomySettings")}}
    {{$a := sdict .Value}}
    {{$symbol := $a.symbol}}
    {{if (dbGet $userID "EconomyInfo")}}
        {{with $.CmdArgs}}
            {{if index . 0}}
                {{if index . 0 | getMember}}
                    {{$user := getMember (index . 0)}}
                    {{$receivingUser := $user.User.ID}}
                    {{if eq $receivingUser $.User.ID}}
                        {{$embed.Set "description" (print "You cannot give money to yourself.")}}
                        {{$embed.Set "color" $errorColor}}
                    {{else}}
                        {{if not (dbGet $receivingUser "EconomyInfo")}}
                            {{dbSet $receivingUser "EconomyInfo" (sdict "cash" 200 "bank" 0)}}
                        {{end}}
                        {{if gt (len $.CmdArgs) 1}}
                            {{$amount := (index $.CmdArgs 1)}}
                            {{if (toInt $amount)}}
                                {{if gt (toInt $amount) 0}}
                                    {{with (dbGet $receivingUser "EconomyInfo")}}
                                        {{$a = sdict .Value}}
                                        {{$receivingBalance := $a.cash}}
                                        {{with (dbGet $userID "EconomyInfo")}}
                                            {{$a = sdict .Value}}
                                            {{$yourBalance := $a.cash}}
                                            {{if gt (toInt $amount) (toInt $yourBalance)}}
                                                {{$embed.Set "description" (print "You cannot give more than you have.")}}
                                                {{$embed.Set "color" $errorColor}}
                                            {{else}}
                                                {{$receivingNewBalance := $receivingBalance | add $amount}}
                                                {{$yourNewBalance := $amount |sub $yourBalance}}
                                                    {{$embed.Set "description" (print "You gave " $symbol $amount " to <@!" $receivingUser ">\nThey now have " $symbol $receivingNewBalance " in cash!")}}
                                                    {{$embed.Set "color" $successColor}}
                                                {{$sdict := (dbGet $receivingUser "EconomyInfo").Value}}
                                                {{$sdict.Set "cash" $receivingNewBalance}}
                                                {{dbSet $receivingUser "EconomyInfo" $sdict}}
                                                {{$sdict := (dbGet $userID "EconomyInfo").Value}}
                                                {{$sdict.Set "cash" $yourNewBalance}}
                                                {{dbSet $userID "EconomyInfo" $sdict}}
                                            {{end}}
                                        {{end}}
                                    {{end}}
                                {{else}}
                                    {{$embed.Set "description" (print "You're unable to give this value, check that you used a valid number above 1")}}
                                    {{$embed.Set "color" $errorColor}}
                                {{end}}
                            {{else}}
                                {{$embed.Set "description" (print "Invalid `Amount argument passed.\nCheck that you used a valid number above 1")}}
                                {{$embed.Set "color" $errorColor}}
                            {{end}}
                        {{else}}
                            {{$embed.Set "description" (print "No `Amount` argument passed.\nSyntax is: `" $.Cmd " <Member:Mention/ID> <Amount:Amount>`")}}
                            {{$embed.Set "color" $errorColor}}
                        {{end}}
                    {{end}}
                {{else}}
                    {{$embed.Set "description" (print "Invalid `user` argument provided.\nSyntax is `" $.Cmd " <User:Mention/ID> <Amount:Amount>`")}}
                    {{$embed.Set "color" $errorColor}}
                {{end}}
            {{end}}
        {{else}}
            {{$embed.Set "description" (print "No `User` argument provided.\nSyntax is `" $.Cmd " <User:Mention/ID> <Amount:Amount>`")}}
            {{$embed.Set "color" $errorColor}}
        {{end}}
    {{else}}
        {{dbSet $userID "EconomyInfo" (sdict "cash" 200 "bank" 0)}}
        {{$embed.Set "description" (print "You were not in the economy database....adding you now\nPlease try again")}}
        {{$embed.Set "color" $errorColor}}
    {{end}}
{{else}}
    {{$embed.Set "description" (print "No `Settings` database found.\nPlease set it up with the default values using `" $prefix "set default`")}}
    {{$embed.Set "color" $errorColor}}
{{end}}
{{sendMessage nil (cembed $embed)}}