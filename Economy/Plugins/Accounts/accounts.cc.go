{{/*
		Made by ranger_4297 (765316548516380732)

	Trigger Type: `Regex`
	Trigger: `\A(-|<@!?204255221017214977>\s*)(account(s?)(-)?)(\s+|\z)`

	©️ RhykerWells 2020-Present
	GNU, GPLV3 License
	Repository: https://github.com/Ranger-4297/YAGPDB-ccs
*/}}

{{/* Only edit below if you know what you're doing (: rawr */}}

{{/* Initiates variBles */}}
{{$uID := str .User.ID}}
{{$oA := "\nAvailable options are `create`, `set`, `list`, `balance`, `withdraw`, `deposit`"}}
{{$oB := "\nAvailable options are `withdrawlimit`, `whitelist`"}}

{{/* Accounts */}}

{{/* Response */}}
{{$e := sdict}}
{{$e.Set "author" (sdict "name" $.User.Username "icon_url" ($.User.AvatarURL "1024"))}}
{{$e.Set "timestamp" currentTime}}
{{$e.Set "color" 0xFF0000}}
{{with dbGet 0 "EconomySettings"}}
	{{$sb := (sdict .Value).symbol}}
	{{with $.CmdArgs}}
		{{if index $.CmdArgs 0}}
			{{$o := (index $.CmdArgs 0) | lower}}
			{{with (dbGet 0 "accounts")}}
				{{$a := sdict .Value}}
				{{$cAs := cslice}}
				{{range $k,$v:= $a}}
					{{- $cAs = $cAs.Append $k -}}
				{{end}}
				{{if eq $o "create"}}
					{{if not ($a.Get $uID)}}
						{{$a.Set $uID (sdict "accountSettings" (sdict "whitelistedUsers" (cslice) "withdrawLimit" 500) "accountBalance" 500)}}
						{{dbSet 0 "accounts" $a}}
						{{$e.Set "description" (print "Account created")}}
						{{$e.Set "color" 0x00ff7b}}
					{{else}}
						{{$e.Set "description" (print "You already have an account")}}
					{{end}}
				{{else if eq $o "delete"}}
					{{if ($a.Get $uID)}}
						{{$a.Del $uID}}
						{{dbSet 0 "accounts" $a}}
						{{$e.Set "description" (print "Account deleted")}}
						{{$e.Set "color" 0x00ff7b}}
					{{else}}
						{{$e.Set "description" (print "You do not have an account. Open one with `" $.Cmd " create`" )}}
					{{end}}
				{{else if eq $o "set"}}
					{{$ct := false}}
					{{if $a.Get $uID}}
						{{if gt (len $.CmdArgs) 1}}
							{{$o = (lower (toString (index $.CmdArgs 1)))}}
							{{$cA := sdict}}
							{{$aS := sdict}}
							{{if eq $o "whitelist" }}
								{{if gt (len $.CmdArgs) 2}}
									{{if eq (index $.CmdArgs 2) "reset"}}
										{{$ct = true}}
										{{$cA = $a.Get $uID}}
										{{$aS = $cA.Get "accountSettings"}}
										{{$aS.Set "whitelistedUsers" cslice}}
										{{$cA.Set "accountSettings" $aS}}
										{{$a.Set $uID $cA}}
										{{$e.Set "description" (print "Whitelist reset")}}
									{{else}}
										{{$ID := (joinStr " " (slice $.CmdArgs 2))}}
										{{$nW := cslice}}
										{{range (reFindAll `\d{17,19}` $ID)}}
											{{- if not (eq (toInt .) (toInt $uID)) -}}
												{{- $nW = $nW.Append .}}
											{{else}}
												{{$e.Set "description" (print "You cannot add yourself to the whitelist")}}
											{{end}}
										{{- end}}
										{{if $nW}}
											{{$ct = true}}
											{{$cA = $a.Get $uID}}
											{{$aS = $cA.Get "accountSettings"}}
											{{$aW := $aS.Get "whitelistedUsers"}}
											{{$aS.Set "whitelistedUsers" ($aW.AppendSlice $nW)}}
											{{$cA.Set "accountSettings" $aS}}
											{{$a.Set $uID $cA}}
											{{$e.Set "description" (print "User(s) now added to whitelist")}}
										{{else}}
											{{$e.Set "description" (print "No user(s) provided")}}
										{{end}}
									{{end}}
								{{else}}
									{{$e.Set "description" (print "No user(s) provided.\nPlease mention or use user IDs\nOptionally; reset your whitelist with `" $.Cmd " " $o " reset`")}}
								{{end}}
							{{else if eq $o "withdrawlimit"}}
								{{if gt (len $.CmdArgs) 2}}
									{{if (ge (toInt (index $.CmdArgs 2)) 0)}}
										{{$ct = true}}
										{{$cA = $a.Get $uID}}
										{{$aS = $cA.Get "accountSettings"}}
										{{$aS.Set "withdrawLimit" (index $.CmdArgs 2)}}
										{{$e.Set "description" (print "New withdraw limit added")}}
										{{$cA.Set "accountSettings" $aS}}
										{{$a.Set $uID $cA}}
									{{else}}
										{{$e.Set "description" (print "Invalid amount provided")}}
									{{end}}
								{{else}}
									{{$e.Set "description" (print "No amount provided")}}
								{{end}}
							{{else}}
								{{$e.Set "description" (print "Invalid option provided" $oB)}}
							{{end}}
						{{else}}
							{{$e.Set "description" (print "No option argument passed.\nSyntax is: `" $.Cmd " set <option> <values>`" $oB)}}
						{{end}}
						{{if $ct}}
							{{$e.Set "color" 0x00ff7b}}
							{{dbSet 0 "accounts" $a}}
						{{end}}
					{{else}}
						{{$e.Set "description" (print "You have no account. Open one with `" $.Cmd " create` or wait to be added to a whitelist")}}
					{{end}}
				{{else if eq $o "list"}}
					{{$aL := cslice}}
					{{if $a.Get $uID}}
						{{$aL = $aL.Append $uID}}
					{{end}}
					{{range $cAs}}
						{{range (($a.Get .).accountSettings).whitelistedUsers}}
							{{- if in . $uID -}}
								{{- $aL = $aL.Append . -}}
							{{end}}
						{{end}}
					{{end}}
					{{if $aL}}
						{{$f := cslice}}
						{{range $aL}}
							{{- $f = $f.Append (sdict "Name" (print " Account ") "Value"  (print "`" . "`") "inline" false) -}}
						{{end}}
						{{$e.Set "fields" $f}}
						{{$e.Set "color" 0x00ff7b}}
					{{else}}
						{{$e.Set "description" (print "You are whitelisted in no accounts. Open one with `" $.Cmd " create` or wait to be added to a whitelist")}}
					{{end}}
				{{else if (reFind `(bal(ance)?|with(draw)?|dep(osit)?)` $o)}}
					{{if gt (len $.CmdArgs) 1}}
						{{$aT := (index $.CmdArgs 1)}}
						{{if (toInt $aT)}}
							{{if in $cAs (toString $aT)}}
								{{$ct := false}}
								{{if $a.Get $uID}}
									{{- $ct = true -}}
								{{end}}
								{{range (($a.Get (toString $aT)).accountSettings).whitelistedUsers}}
									{{- if in . $uID -}}
										{{- $ct = true -}}
									{{end}}
								{{end}}
								{{$aB := toInt ($a.Get (toString $aT)).accountBalance}}
								{{if and ((reFind `bal(ance)?` $o)) $ct}}
									{{$e.Set "description" (print "The account `" $aT "` has a balance of " $sb (humanizeThousands $aB))}}
									{{$e.Set "color" 0x00ff7b}}
								{{else if and ((reFind `with(draw)?|dep(osit)?` $o)) $ct}}
									{{if gt (len $.CmdArgs) 2}}
										{{$amt := (index $.CmdArgs 2)}}
										{{$lm := ($a.Get (toString $aT)).accountSettings.withdrawLimit}}
										{{if gt (toInt $amt) 0}}
											{{$amt = toInt $amt}}
											{{$c := or (dbGet $uID "cash").Value 0 | toInt}}
											{{$nAB := ""}}
											{{$aN := false}}
											{{$aR := ""}}
											{{if eq $o "with" "withdraw"}}
												{{$aR = "the account has"}}
												{{if and (gt $amt 0) (le $amt $aB)}}
													{{if le $amt $lm}}
														{{$aN = "withdrawn"}}
														{{$nAB = sub $aB $amt}}
														{{$c = add $c $amt}}
													{{else}}
														{{$aR = print $sb $lm}}
													{{end}}
												{{end}}
											{{else if eq $o "dep" "deposit"}}
												{{$aR = "you have"}}
												{{if and (gt $amt 0) (le $amt $c)}}
													{{$aN = "deposited"}}
													{{$nAB = add $aB $amt}}
													{{$c = sub $c $amt}}
												{{end}}
											{{end}}
											{{if $aN}}
												{{$mA := $a.Get (toString $aT)}}
												{{$mA.Set "accountBalance" $nAB}}
												{{$mA.Set "accountHistory"}}
												{{$e.Set "description" (print "You've just " $aN " " $sb $amt)}}
												{{$e.Set "color" 0x00ff7b}}
												{{$a.Set $aT $mA}}
												{{dbSet 0 "accounts" $a}}
												{{dbSet $uID "cash" $c}}
											{{else}}
												{{$e.Set "description" (print "You can't " $o " more than " $aR)}}
											{{end}}
										{{else}}
											{{$e.Set "description" (print "Invalid `amount` argument provided")}}
										{{end}}
									{{else}}
										{{$e.Set "description" (print "No `amount` argument provided")}}
									{{end}}
								{{else}}
									{{$e.Set "description" (print "You do not have access to this account")}}
								{{end}}
							{{else}}
								{{$e.Set "description" (print "This account does not exist")}}
							{{end}}
						{{else}}
							{{$e.Set "description" (print "Invalid account type provided")}}
						{{end}}
					{{else}}
						{{$e.Set "description" (print "No account provided\nView your current account(s) with `" $.Cmd " list`")}}
					{{end}}
				{{else}}
					{{$e.Set "description" (print "Invalid `option` argument passed" $oA)}}
				{{end}}
			{{end}}
		{{end}}
	{{else}}
		{{$e.Set "description" (print "No option argument passed" $oA)}}
	{{end}}
{{else}}
	{{$e.Set "description" (print "No `Settings` database found.\nPlease set it up with the default values using `" .ServerPrefix "server-set default`")}}
{{end}}
{{sendMessage nil (cembed $e)}}