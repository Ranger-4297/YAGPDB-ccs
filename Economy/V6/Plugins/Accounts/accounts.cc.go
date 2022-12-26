{{/*
        Made by Ranger (765316548516380732)

        Trigger Type: `Regex`
        Trigger: ``

    ©️ Ranger 2020-Present
    GNU, GPLV3 License
    Repository: https://github.com/Ranger-4297/YAGPDB-ccs
*/}}

{{/* Only edit below if you know what you're doing (: rawr */}}

{{/* Initiates variables */}}
{{$userID := .User.ID}}
{{$sucCol := 0x00ff7b}}
{{$errCol := 0xFF0000}}
{{$prefix := index (reFindAllSubmatches `.*?: \x60(.*)\x60\z` (execAdmin "Prefix")) 0 1 }}

{{/* Buy item */}}

{{/* Response */}}
{{$embed := sdict}}
{{$embed.Set "author" (sdict "name" $.User.Username "icon_url" ($.User.AvatarURL "1024"))}}
{{$embed.Set "timestamp" currentTime}}
{{with dbGet 0 "EconomySettings"}}
    {{with $db := (dbGet 0 "accounts")}}
        {{$a := sdict .Value}}
        {{$crtAccts := cslice}}
        {{if $db}}
            {{range $k,$v:= $a}}{{- $crtAccts = $crtAccts.Append $k -}}{{end}}
        {{end}}
        {{with .CmdArgs}}
            {{if index $.CmdArgs 0}}
                {{$option := (index $.CmdArgs 0) | lower}}
                {{if eq $option "create"}}
                    {{dbSet "accounts" (sdict (toString $.User.ID) (sdict "accountSettings" (sdict "whitelistedUsers" (cslice) "withdrawLimit" 500) "accountBalance" 500))}}
                {{else if eq $option "set"}}
                    {{if gt (len $.CmdArgs) 1}}
                        {{$val := (lower (toString (index $.CmdArgs)))}}
                        {{if eq $val "wl" "whitelist"}}
                            {{if gt (len $.CmdArgs) 2}}
                                {{$ID := (joinStr " " (slice $.CmdArgs 2))}}
                                {{$users := cslice}}
                                {{range (reFindAll `\d{17,19}` $ID)}}{{- if not eq $ID $.User.ID -}}{{- $users = $users.Append .}}{{else}}{{$embed.Set "description" (print "You cannot add yourself to the whitelist")}}{{$embed.Set "color" $errCol}}{{end}}{{- end}}
                                {{if $users}}
                                    {{$acct := $a.Get (toString $.User.ID)}}
                                    {{$stngs := $acct.Get "accountSettings"}}
                                    {{$wList := $stngs.Get "whitelistedUsers"}}
                                    {{$wList = $wList.AppendSlice $users}}
                                    {{$stngs.Set "whitelistedUsers" $wList}}
                                    {{$acct.Set "accountSettings" $stngs}}
                                    {{$a.Set (toString $.User.ID) $acct}}
                                    {{dbSet 0 "accounts" $a}}
                                    {{$embed.Set "description" (print "User(s) now added to whitelist.")}}
                                    {{$embed.Set "color" $sucCol}}
                                {{end}}
                            {{else}}
                                {{$embed.Set "description" (print "No user provided.\nPlease ensure that you use UserIDs")}}
                                {{$embed.Set "color" $errCol}}
                            {{end}}
                        {{else if eq $val "withdrawlimit" "withlimit" "wl"}}
                            {{if gt (len $.CmdArgs) 2}}
                                {{if (toInt $.CmdArgs 2)}}
                                    {{$acct := $a.Get (toString $.User.ID)}}
                                    {{$stngs := $acct.Get "accountSettings"}}
                                    {{$stngs.Set "withdrawLimit" (toInt $.CmdArgs 2)}}
                                    {{$acct.Set "accountSettings" $stngs}}
                                    {{$a.Set (toString $.User.ID) $acct}}
                                    {{dbSet 0 "accounts" $a}}
                                    {{$embed.Set "description" (print "New withdraw limit added.")}}
                                    {{$embed.Set "color" $sucCol}}
                                {{else}}
                                    {{$embed.Set "description" (print "Invalid amount provided.")}}
                                    {{$embed.Set "color" $errCol}}
                                {{end}}
                            {{else}}
                                {{$embed.Set "description" (print "No amount provided.")}}
                                {{$embed.Set "color" $errCol}}
                            {{end}}
                        {{else}}
                            {{$embed.Set "description" (print "Invalid option provided.")}}
                            {{$embed.Set "color" $errCol}}
                        {{end}}
                    {{else}}
                        {{$msg.Set "description" (print "No option argument passed.\nSyntax is: `" $.Cmd " <Option> <Values>`")}}
                        {{$embed.Set "color" $errCol}}
                    {{end}}
                {{else if eq $option "list"}}
                    {{$acctList := cslice}}
                    {{if $a.Get (toString $.User.ID)}}
                        {{$acctList = $acctList.Append $.User.ID}}
                    {{end}}
                    {{range $crtAccts}}{{- $acct := . -}}{{- $wList := (($a.Get .).accountSettings).whitelistedUsers -}}{{range $wList}}{{- if in . (toString $.User.ID) -}}{{- $acctList = $acctList.Append $acct -}}{{end}}{{end}}{{end}}
                    {{if $acctList}}
                        {{$fields := cslice}}
                        {{range $acctList}}{{- $fields = $fields.Append (sdict "Name" (print " Account ") "value"  (print .) "inline" false) -}}{{end}}
                        {{$embed.Set "fields" $fields}}
                        {{$embed.Set "color" $sucCol}}
                    {{else}}
                        {{$msg.Set "description" (print "You have no accounts. Open one with `" $.Cmd "create` or wait to be added to a whitelist.")}}
                        {{$embed.Set "color" $errCol}}
                    {{end}}
                {{else if eq $option "balance" "withdraw" "dep" "deposit"}}
                    {{if gt (len $.CmdArgs) 1}}
                        {{$acct := (index $.CmdArgs 1)}}
                        {{if (toInt $acct)}}
                            {{if in $crtAccts (toString $acct)}}
                                {{$continue := false}}
                                {{- $wList := (($a.Get (toString $acct)).accountSettings).whitelistedUsers -}}
                                {{range $wList}}{{- if in . (toString $.User.ID) -}}{{- $continue = true -}}{{end}}{{end}}
                                {{$acctBal := (($a.Get (toString $acct)).accountBalance)}}
                                {{if and (eq $option "balance") $continue}}
                                    {{$embed.Set "description" (print "The account `" $acct "` has a balance of " $acctBal)}}
                                {{else if and (eq $option "withdraw" "with" "dep" "with") $continue}}
                                    {{if gt (len $.CmdArgs) 2}}
                                        {{$amount := (index $.CmdArgs 2)}}
                                        {{$limit := (toInt ($a.get (toString $account)).accountSettings.withdrawLimit)}}
                                        {{if (toInt $amount)}}
                                            {{if not (dbGet $user.ID "EconomyInfo")}}
                                                {{dbSet $user.ID "EconomyInfo" (sdict "cash" 200 "bank" 0)}}
                                            {{end}}
                                            {{with (dbGet $user.ID "EconomyInfo")}}
                                                {{$uDb := sdict .Value}}
                                                {{$cash (toInt $uDB.cash)}}
                                                {{$bank := (toInt $uDB.bank)}}
                                                {{$newUsrBal := ""}}
                                                {{$newAccBal := ""}}
                                                {{if lt (toInt $amount) $limit}}
                                                    {{if eq $option "with" "withdraw"}}
                                                        {{if lt (toInt $amount) $accBal}}
                                                            {{$newUsrBal = $amount | add $cash}}
                                                            {{$newAccBal = $amount | sub $acctBal}}
                                                            {{$a.Set  ($a.Get (toString $acct)).Get "accountBalance" $newAccBal}}
                                                        {{else}}
                                                            {{$embed.Set "description" (print "You cannot withdraw more than the account has")}}
                                                            {{$embed.Set "color" $errCol}}
                                                        {{end}}
                                                    {{else if eq $option "dep" "deposit"}}
                                                        {{if gt (toInt $amount) $cash}}
                                                            {{$newUsrBal = $amount | sub $cash}}
                                                            {{$newAccBal = $amount | add $acctBal}}
                                                            {{$a.Set  ($a.Get (toString $acct)).Get "accountBalance" $newAccBal}}
                                                        {{else}}
                                                            {{$embed.Set "description" (print "You cannot deposit more than the you have")}}
                                                            {{$embed.Set "color" $errCol}}
                                                        {{end}}
                                                    {{end}}
                                                    {{dbSet $user.ID "EconomyInfo" (sdict "cash" $newUsrBal "bank" $bank)}}
                                                    {{dbSet 0 "accounts" $a}}
                                                {{else}}
                                                    {{$embed.Set "description" (print "You cannot " $option " more than " $limit " at one time")}}
                                                    {{$embed.Set "color" $errCol}}
                                                {{end}}
                                            {{else}}
                                                {{$embed.Set "description" (print "Invalid `amount` argument provided.\nCommand syntax is `" $.Cmd " <Amount:Amount>`")}}
                                                {{$embed.Set "color" $errCol}}
                                            {{end}}
                                        {{end}}
                                    {{else}}
                                        {{$embed.Set "description" (print "No `amount` argument provided.\nCommand syntax is `" $.Cmd " <Amount:Amount>`")}}
                                        {{$embed.Set "color" $errCol}}
                                    {{end}}
                                {{else}}
                                    {{$msg.Set "description" (print "You do not have access to this account")}}
                                    {{$embed.Set "color" $errCol}}
                                {{end}}
                            {{else}}
                                {{$msg.Set "description" (print "This account does not exist")}}
                                {{$embed.Set "color" $errCol}}
                            {{end}}
                        {{else}}
                            {{$msg.Set "description" (print "Invalid account type provided")}}
                            {{$embed.Set "color" $errCol}}
                        {{end}}
                    {{else}}
                        {{$msg.Set "description" (print "No account provided")}}
                        {{$embed.Set "color" $errCol}}
                    {{end}}
                {{else}}
                    {{$msg.Set "description" (print "Invalid option argument passed.\nSyntax is: `" $.Cmd " <Option> <Values>`")}}
                    {{$embed.Set "color" $errCol}}
                {{end}}
            {{end}}
        {{else}}
            {{$msg.Set "description" (print "No option argument passed.\nSyntax is: `" $.Cmd " <Option> <Values>`")}}
            {{$embed.Set "color" $errCol}}
        {{end}}
    {{end}}
{{else}}
	{{$embed.Set "description" (print "No `Settings` database found.\nPlease set it up with the default values using `" $prefix "set default`")}}
	{{$embed.Set "color" $errCol}}
{{end}}
{{sendMessage nil (cembed $embed)}}