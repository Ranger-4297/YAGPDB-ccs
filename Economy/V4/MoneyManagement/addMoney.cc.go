    {{/*
            Made by Ranger (765316548516380732)

        Trigger Type: `Regex`
        Trigger: `\A(-|<@!?204255221017214977>\s*)(add-?money|inc(crease)?-?money)(\s+|\z)`

        ©️ Ranger 2020-Present
        GNU, GPLV3 License
        Repository: https://github.com/Ranger-4297/YAGPDB-ccs
    */}}


    {{/* Only edit below if you know what you're doing (: rawr */}}

{{/* Initiates variables */}}
{{$successColor := 0x00ff7b}}
{{$errorColor := 0xFF0000}}
{{$prefix := index (reFindAllSubmatches `.*?: \x60(.*)\x60\z` (execAdmin "Prefix")) 0 1 }}

{{/* Adds money to given user */}}

{{/* Response */}}
{{$embed := sdict}}
{{$embed.Set "author" (sdict "name" $.User.Username "icon_url" ($.User.AvatarURL "1024"))}}
{{$embed.Set "timestamp" currentTime}}
{{$perms := split (index (split (exec "viewperms" ) "\n" ) 2) ", "}}
{{if or (in $perms "Administrator") (in $perms "ManageServer")}}
    {{with (dbGet 0 "EconomySettings")}}
        {{$a := sdict .Value}}
        {{$symbol := $a.symbol}}
        {{with $.CmdArgs}}
            {{if (index . 0)}}
                {{if (index . 0) | getMember}}
                    {{$user := (index . 0) | getMember}}
                    {{$receivingUser := $user.User.ID}}
                    {{if not (dbGet $receivingUser "EconomyInfo")}}
                        {{dbSet $receivingUser "EconomyInfo" (sdict "cash" 200 "bank" 0)}}
                    {{end}}
                    {{if gt (len $.CmdArgs) 1}}
                        {{$moneyDestination := (index . 1) | lower}}
                        {{if eq $moneyDestination "cash" "bank"}}
                            {{with (dbGet $receivingUser "EconomyInfo")}}
                                {{$a = sdict .Value}}
                                {{$receivingBalance := $a.Get $moneyDestination}}
                                {{if gt (len $.CmdArgs) 2}}
                                    {{$amount := (index $.CmdArgs 2)}}
                                    {{if (toInt $amount)}}
                                        {{if gt (toInt $amount) 0}}
                                            {{$receivingNewBalance := $receivingBalance | add $amount}}
                                            {{$embed.Set "description" (print "You added " $symbol (humanizeThousands $amount) " to <@!" $receivingUser ">'s " $moneyDestination "\nThey now have " $symbol (humanizeThousands $receivingNewBalance) " in their " $moneyDestination "!")}}
                                            {{$embed.Set "color" $successColor}}
                                            {{$a.Set $moneyDestination $receivingNewBalance}}
                                            {{dbSet $receivingUser "EconomyInfo" $a}}
                                        {{else}}
                                            {{$embed.Set "description" (print "You're unable to add this value, check that you used a valid number above 1")}}
                                            {{$embed.Set "color" $errorColor}}
                                        {{end}}
                                    {{else}}
                                        {{$embed.Set "description" (print "You're unable to add this value, check that you used a valid number")}}
                                        {{$embed.Set "color" $errorColor}}
                                    {{end}}
                                {{else}}
                                    {{$embed.Set "description" (print "No `Amount` argument passed.\nSyntax is: `" $.Cmd " <Member:Mention/ID> <Destination:Cash/Bank> <Amount:Amount>`")}}
                                    {{$embed.Set "color" $errorColor}}
                                {{end}}
                            {{end}}
                        {{else}}
                            {{$embed.Set "description" (print "Invalid `Destination` argument provided.\nSyntax is `" $.Cmd " <User:Mention/ID> <Destination:Cash/Bank> <Amount:Amount>`")}}
                            {{$embed.Set "color" $errorColor}}
                        {{end}}
                    {{else}}
                        {{$embed.Set "description" (print "No `Destination` argument provided.\nSyntax is `" $.Cmd " <User:Mention/ID> <Destination:Cash/Bank> <Amount:Amount>`")}}
                        {{$embed.Set "color" $errorColor}}
                    {{end}}
                {{else}}
                    {{$embed.Set "description" (print "Invalid `User` argument provided.\nSyntax is `" $.Cmd " <User:Mention/ID> <Destination:Cash/Bank> <Amount:Amount>`")}}
                    {{$embed.Set "color" $errorColor}}
                {{end}}
            {{end}}
        {{else}}
            {{$embed.Set "description" (print "No `User` argument provided.\nSyntax is `" $.Cmd " <User:Mention/ID> <Destination:Cash/Bank> <Amount:Amount>`")}}
            {{$embed.Set "color" $errorColor}}
        {{end}}
    {{else}}
        {{$embed.Set "description" (print "No `Settings` database found.\nPlease set it up with the default values using `" $prefix "set default`")}}
        {{$embed.Set "color" $errorColor}}
    {{end}}
{{else}}
    {{$embed.Set "description" (print "Insufficient permissions.\nTo use this command you need to have either `Administrator` or `ManageServer` permissions")}}
    {{$embed.Set "color" $errorColor}}
{{end}}
{{sendMessage nil (cembed $embed)}}