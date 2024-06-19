{{/*
		Made by ranger_4297 (765316548516380732)

	Trigger Type: `Regex`
	Trigger: `\A(-|<@!?204255221017214977>\s*)(account(s?)(-)?)(\s+|\z)`

	©️ RhykerWells 2020-Present
	GNU, GPLV3 License
	Repository: https://github.com/Ranger-4297/YAGPDB-ccs
*/}}

{{/* Only edit below if you know what you're doing (: rawr */}}

{{/* Initiates variables */}}
{{$uID := str .User.ID}}
{{$oA := "\nAvailable options are `create`, `set`, `list`, `balance`, `withdraw`, `deposit`, `view`"}}
{{$oB := "\nAvailable options are `withdrawlimit`, `whitelist`"}}

{{/* Accounts */}}

{{/* Response */}}
{{$e := sdict "author" (sdict "name" .User.Username "icon_url" (.User.AvatarURL "1024")) "timestamp" currentTime "color" 0xFF0000}}
{{$s := dbGet 0 "EconomySettings"}}
{{if not $s}}
	{{$e.Set "description" (print "No `Settings` database found.\nPlease set it up with the default values using `" .ServerPrefix "server-set default`")}}
	{{sendMessage nil (cembed $e)}}
	{{return}}
{{end}}
{{$aDB := (dbGet 0 "accounts").Value}}
{{$cS := $s.Value.symbol}}
{{if not .CmdArgs}}
	{{$e.Set "description" (print "No option argument passed" $oA)}}
	{{sendMessage nil (cembed $e)}}
	{{return}}
{{end}}
{{$o := lower (index .CmdArgs 0)}}
{{$fAL := cslice}}
{{range $k,$v:= $aDB}}
	{{- $fAL = $fAL.Append $k -}}
{{end}}
{{if eq $o "create"}}
	{{if $aDB.Get $uID}}
		{{$e.Set "description" (print "You already have an account")}}
		{{sendMessage nil (cembed $e)}}
		{{return}}
	{{end}}
	{{$aDB.Set $uID (sdict "accountSettings" (sdict "whitelistedUsers" (cslice) "withdrawLimit" 500) "accountBalance" 0 "balanceHistory" sdict "historyCounter" 0)}}
	{{dbSet 0 "accounts" $aDB}}
	{{$e.Set "description" (print "Account created")}}
	{{$e.Set "color" 0x00ff7b}}
{{else if eq $o "delete"}}
	{{if not ($aDB.Get $uID)}}
		{{sendMessage nil (cembed $e)}}
		{{return}}
	{{end}}
	{{$aDB.Del $uID}}
	{{dbSet 0 "accounts" $aDB}}
	{{$e.Set "description" (print "Account deleted")}}
	{{$e.Set "color" 0x00ff7b}}
{{else if eq $o "set"}}
	{{$cA := $aDB.Get $uID}}
	{{if not $cA}}
		{{$e.Set "description" (print "You have no account. Open one with `" .Cmd " create` or wait to be added to a whitelist")}}
		{{sendMessage nil (cembed $e)}}
		{{return}}
	{{end}}
	{{if lt (len .CmdArgs) 2}}
		{{$e.Set "description" (print "No option argument passed.\nSyntax is: `" .Cmd " set <option> <values>`" $oB)}}
		{{sendMessage nil (cembed $e)}}
		{{return}}
	{{end}}
	{{$o = lower (index .CmdArgs 1)}}
	{{$aS := $cA.accountSettings}}
	{{if eq $o "whitelist" }}
		{{if lt (len .CmdArgs) 3}}
			{{$e.Set "description" (print "No user(s) provided.\nPlease mention or use user IDs\nOptionally; reset your whitelist with `" .Cmd " " $o " reset`")}}
			{{sendMessage nil (cembed $e)}}
			{{return}}
		{{end}}
		{{if eq (index .CmdArgs 2) "reset"}}
			{{$aS.Set "whitelistedUsers" cslice}}
			{{$e.Set "description" "Whitelist reset"}}
		{{else}}
			{{$ID := joinStr " " (slice .CmdArgs 2)}}
			{{$nW := cslice}}
			{{range reFindAll `\d{17,19}` $ID}}
				{{- if eq (toInt .) (toInt $uID)}}
					{{- $e.Set "description" (print "You cannot add yourself to the whitelist")}}
					{{- sendMessage nil (cembed $e)}}
					{{- return}}
				{{- end}}
				{{- $nW = $nW.Append . -}}
			{{end}}
			{{if not $nW}}
				{{$e.Set "description" (print "No user(s) added to the whitelist")}}
				{{sendMessage nil (cembed $e)}}
				{{return}}
			{{end}}
			{{$aS.Set "whitelistedUsers" $nW}}
			{{$e.Set "description" (print "User(s) now added to whitelist")}}
		{{end}}
	{{else if eq $o "withdrawlimit"}}
		{{if lt (len .CmdArgs) 3}}
			{{$e.Set "description" (print "No amount provided")}}
			{{sendMessage nil (cembed $e)}}
			{{return}}
		{{end}}
		{{if le (toInt (index .CmdArgs 2)) 0}}
			{{$e.Set "description" (print "Invalid amount provided")}}
			{{sendMessage nil (cembed $e)}}
			{{return}}
		{{end}}
		{{$aS.Set "withdrawLimit" (toInt (index .CmdArgs 2))}}
		{{$e.Set "description" "New withdraw limit added"}}
	{{else}}
		{{$e.Set "description" (print "Invalid option provided" $oB)}}
		{{sendMessage nil (cembed $e)}}
		{{return}}
	{{end}}
	{{$e.Set "color" 0x00ff7b}}
	{{dbSet 0 "accounts" $aDB}}
{{else if eq $o "list"}}
	{{$aAL := cslice}}
	{{if $aDB.Get $uID}}
		{{$aAL = $aAL.Append $uID}}
	{{end}}
	{{range $fAL}}
		{{- range ($aDB.Get .).accountSettings.whitelistedUsers}}
			{{- if in . $uID}}
				{{- $aAL = $aAL.Append .}}
			{{- end}}
		{{- end -}}
	{{end}}
	{{if not $aAL}}
		{{$e.Set "description" (print "You are whitelisted in no accounts. Open one with `" .Cmd " create` or wait to be added to a whitelist")}}
		{{sendMessage nil (cembed $e)}}
		{{return}}
	{{end}}
	{{$fields := cslice}}
	{{range $aAL}}
		{{- $fields = $fields.Append (sdict "Name" " Account " "Value"  (print "`" . "`") "inline" false) -}}
	{{end}}
	{{$e.Set "fields" $fields}}
	{{$e.Set "color" 0x00ff7b}}
{{else if reFind `(bal(ance)?|with(draw)?|dep(osit)|view?)` $o}}
	{{if lt (len .CmdArgs) 2}}
		{{$e.Set "description" (print "No account provided\nView your current account(s) with `" .Cmd " list`")}}
		{{sendMessage nil (cembed $e)}}
		{{return}}
	{{end}}
	{{$aID := toInt (index .CmdArgs 1)}}
	{{if not $aID}}
		{{$e.Set "description" (print "Invalid account type provided")}}
		{{sendMessage nil (cembed $e)}}
		{{return}}
	{{end}}
	{{if not (in $fAL (toString $aID))}}
		{{$e.Set "description" (print "This account does not exist")}}
		{{sendMessage nil (cembed $e)}}
		{{return}}
	{{end}}
	{{$c := false}}
	{{if $aDB.Get $uID}}
		{{- $c = true -}}
	{{end}}
	{{range (($aDB.Get (toString $aID)).accountSettings).whitelistedUsers}}
		{{- if in . $uID}}
			{{- $c = true}}
		{{- end -}}
	{{end}}
	{{$cA := $aDB.Get (toString $aID)}}
	{{$aB := toInt $cA.accountBalance}}
	{{$aS := $cA.accountSettings}}
	{{$wL := $aS.withdrawLimit}}
	{{$bH := $cA.balanceHistory}}
	{{if and (reFind `bal(ance)?` $o) $c}}
		{{$e.Set "description" (print "The account `" $aID "` has a balance of " $cS (humanizeThousands $aB))}}
		{{$e.Set "color" 0x00ff7b}}
	{{else if and (eq "view" $o) $c}}
		{{$p := 1}}
		{{if and (index .CmdArgs 2) (toInt (index .CmdArgs 2)) (ge (toInt (index .CmdArgs 2)) 1)}}{{$p = (toInt (index .CmdArgs 2))}}{{end}}
		{{$d := print "### Account holder\n<@!" $aID ">\n### Account balance\n" $cS (humanizeThousands $aB) "\n### Account history (page " $p ")\n"}}
		{{range (seq (toInt (print (sub $p 1) "1")) (toInt (print $p "1")))}}
			{{if $aH := $bH.Get (str .)}}
				{{if eq $aH.action "w"}}{{$s = "ACCOUNT WITHDRAWN"}}{{else}}{{- $s = "ACCOUNT DEPOSIT"}}{{end}}
				{{$d = print $d "<t:" $aH.timestamp ":f> - <@!" $aH.user "> **" $s "**\n> **Amount:** £" (humanizeThousands $aH.amount) " | **Previous balance:** £" (humanizeThousands $aH.before) " | **New balance:** £" (humanizeThousands $aH.after) "\n"}}
			{{end}}
		{{end}}
		{{if not (in $d "Amount")}}{{$d = print $d "No accounts found on this page"}}{{end}}
		{{$e.Set "description" $d}}
		{{$e.Set "color" 0x00ff7b}}
		{{$e.Set "footer" (sdict "text" (print print "page " $p "/" (roundCeil (fdiv (len $bH) 10))))}}
	{{else if and (reFind `with(draw)?|dep(osit)?` $o) $c}}
		{{if lt (len .CmdArgs) 3}}
			{{$e.Set "description" (print "No `amount` argument provided")}}
			{{sendMessage nil (cembed $e)}}
			{{return}}
		{{end}}
		{{$mA := toInt (index .CmdArgs 2)}}
		{{if le $mA 0}}
			{{$e.Set "description" (print "Invalid `amount` argument provided")}}
			{{sendMessage nil (cembed $e)}}
			{{return}}
		{{end}}
		{{$cash := toInt (dbGet (toInt $uID) "cash").Value}}
		{{$nAB := 0}}
		{{$cM := ""}}
		{{$k := ""}}
		{{if and (not $cA.historyCounter) (not $cA.balanceHistory) }}
            {{$cA.Set "historyCounter" 0}}
            {{$cA.Set "balanceHistory" sdict}}
        {{end}}
		{{$counter := $cA.historyCounter}}
        {{$balanceHistory := $cA.balanceHistory}}
		{{if eq $o "with" "withdraw"}}
			{{if gt $mA $aB}}
				{{$e.Set "description" (print "You can't withdraw more than the account has")}}
				{{sendMessage nil (cembed $e)}}
				{{return}}
			{{end}}
			{{if gt $mA $wL}}
				{{$e.Set "description" (print "You can't withdraw more than " $cS $wL " at once")}}
				{{sendMessage nil (cembed $e)}}
				{{return}}
			{{end}}
			{{$cM = "withdrawn"}}{{$k = "w"}}
			{{$nAB = sub $aB $mA}}
			{{$cash = add $cash $mA}}
		{{else if eq $o "dep" "deposit"}}
			{{if gt $mA $cash}}
				{{$e.Set "description" (print "You can't deposit more than " $cS $cash)}}
				{{sendMessage nil (cembed $e)}}
				{{return}}
			{{end}}
			{{$cM = "deposited"}}{{$k = "d"}}
			{{$nAB = add $aB $mA}}
			{{$cash = sub $cash $mA}}
		{{end}}
		{{$counter := add $counter 1}}
		{{$cA.Set "historyCounter" $counter}}
		{{$balanceHistory.Set (str $counter) (sdict "timestamp" currentTime.Unix "user" .User.ID "action" $k "amount" $mA "before" $aB "after" $nAB)}}
		{{$cA.Set "balanceHistory" $balanceHistory}}
		{{$cA.Set "accountBalance" $nAB}}
		{{$e.Set "description" (print "You've just " $cM " " $cS $mA)}}
		{{$e.Set "color" 0x00ff7b}}
		{{dbSet 0 "accounts" $aDB}}
		{{dbSet (toInt $uID) "cash" $cash}}
	{{else}}
		{{$e.Set "description" (print "You do not have access to this account")}}
	{{end}}
{{else}}
	{{$e.Set "description" (print "Invalid `option` argument passed" $oA)}}
{{end}}
{{sendMessage nil (cembed $e)}}