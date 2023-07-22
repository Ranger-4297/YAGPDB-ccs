{{/*
		Made by Ranger (765316548516380732)

	Trigger Type: `Regex`
	Trigger: `\A(-|<@!?204255221017214977>\s*)(set-?bal(ance)?)(\s+|\z)`

	©️ Ranger 2020-Present
	GNU, GPLV3 License
	Repository: https://github.com/Ranger-4297/YAGPDB-ccs
*/}}


{{/* Only edit below if you know what you're doing (: rawr */}}

{{/* Initiates variables */}}
{{$successColor := 0x00ff7b}}
{{$errorColor := 0xFF0000}}
{{$prefix := .ServerPrefix }}

{{/* Sets balance */}}

{{/* Response */}}
{{$embed := sdict}}
{{$embed.Set "author" (sdict "name" $.User.Username "icon_url" ($.User.AvatarURL "1024"))}}
{{$embed.Set "timestamp" currentTime}}
{{$perms := split (index (split (exec "viewperms") "\n") 2) ", "}}
{{if or (in $perms "Administrator") (in $perms "ManageServer")}}
	{{with (dbGet 0 "EconomySettings")}}
		{{$a := sdict .Value}}
		{{$symbol := $a.symbol}}
		{{with $.CmdArgs}}
            {{if index . 0 | getMember}}
                {{$user := getMember (index . 0)}}
				{{$user = $user.User}}
                {{if gt (len $.CmdArgs) 1}}
                    {{$moneyDestination := (lower (index . 1))}}
                    {{if eq $moneyDestination "cash" "bank"}}
                        {{$bankDB := or (dbGet 0 "bank").Value sdict}}
						{{$bankUser := or ($bankDB.Get (toString $user.ID)) 0 | toInt}}
						{{$cash := or (dbGet $user.ID "cash").Value 0 | toInt}}
                        {{if gt (len $.CmdArgs) 2}}
                            {{$amount := (index $.CmdArgs 2)}}
							{{if (toInt $amount)}}
                                {{if eq $moneyDestination "bank"}}
                                    {{$bankDB.Set (toString $user.ID) $amount}}
                                {{else}}
                                    {{$cash = $amount}}
                                {{end}}
                                {{dbSet 0 "bank" $bankDB}}
                                {{dbSet $user.ID "cash" $cash}}
                                {{$embed.Set "description" (print $user.Mention "'s " $moneyDestination " set to " $amount)}}
                                {{$embed.Set "color" $successColor}}
                            {{else}}
                                {{$embed.Set "description" (print "Invalid `Amount` argument passed.\nSyntax is: `" $.Cmd " <Member:Mention/ID> <Destination:Cash/Bank> <Amount:Amount>`")}}
                                {{$embed.Set "color" $errorColor}}
                            {{end}}
                        {{else}}
                            {{$embed.Set "description" (print "No `Amount` argument passed.\nSyntax is: `" $.Cmd " <Member:Mention/ID> <Destination:Cash/Bank> <Amount:Amount>`")}}
                            {{$embed.Set "color" $errorColor}}
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
        {{else}}
			{{$embed.Set "description" (print "No `User` argument provided.\nSyntax is `" $.Cmd " <User:Mention/ID> <Destination:Cash/Bank> <Amount:Amount>`")}}
			{{$embed.Set "color" $errorColor}}
		{{end}}
    {{else}}
        {{$embed.Set "description" (print "No `Settings` database found.\nPlease set it up with the default values using `" $prefix "server-set default`")}}
        {{$embed.Set "color" $errorColor}}
    {{end}}
{{else}}
	{{$embed.Set "description" (print "Insufficient permissions.\nTo use this command you need to have either `Administrator` or `ManageServer` permissions")}}
	{{$embed.Set "color" $errorColor}}
{{end}}
{{sendMessage nil (cembed $embed)}}