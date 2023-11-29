{{/*
		Made by ranger_4297 (765316548516380732)

	Trigger Type: `Regex`
	Trigger: `\A(-|<@!?204255221017214977>\s*)(account(s?)(-)?)(\s+|\z)`

	©️ Ranger 2020-Present
	GNU, GPLV3 License
	Repository: https://github.com/Ranger-4297/YAGPDB-ccs
*/}}

{{/* Only edit below if you know what you're doing (: rawr */}}

{{/* Initiates variBles */}}
{{$u := .User.ID}}

{{/* Accounts */}}

{{/* Response */}}
{{$e := sdict}}
{{$e.Set "author" (sdict "name" $.User.Username "icon_url" ($.User.AvatarURL "1024"))}}
{{$e.Set "timestamp" currentTime}}
{{$e.Set "color" 0xFF0000}}
{{$o := "\nAvailBle options are `initialize`, `create`, `set`, `list`, `balance`, `withdraw`, `deposit`"}}
{{with dbGet 0 "EconomySettings"}}
	{{with $.CmdArgs}}
		{{if index $.CmdArgs 0}}
			{{$c := (index $.CmdArgs 0) | lower}}
			{{if (reFind `in(n)?it(iali(z|s)e)?` $c)}}
				{{$p := split (index (split (exec "viewperms") "\n") 2) ", "}}
				{{if or (in $p "Administrator") (in $p "ManageServer")}}
					{{dbSet 0 "accounts" sdict}}
					{{$e.Set "description" (print "Database initialized")}}
					{{$e.Set "color" 0x00ff7b}}
				{{else}}
					{{$e.Set "description" (print "Insufficient permissions.\nTo use this command you need to have either `Administrator` or `ManageServer` permissions")}}
				{{end}}
			{{else}}
				{{with $d := (dbGet 0 "accounts")}}
					{{$a := sdict .Value}}
					{{$C := cslice}}
					{{if $d}}
						{{range $k,$v:= $a}}
							{{- $C = $C.Append $k -}}
						{{end}}
					{{end}}
					{{if eq $c "create"}}
						{{if not ($a.Get (toString $u))}}
							{{$a.Set (toString $u) (sdict "accountSettings" (sdict "whitelistedUsers" (cslice) "withdrawimit" 500) "accountBalance" 500)}}
							{{dbSet 0 "accounts" $a}}
							{{$e.Set "description" (print "Account created")}}
							{{$e.Set "color" 0x00ff7b}}
						{{else}}
							{{$e.Set "description" (print "You have an account already")}}
						{{end}}
					{{else if eq $c "delete"}}
						{{if ($a.Get (toString $u))}}
							{{$a.Del (toString $u)}}
							{{dbSet 0 "accounts" $a}}
							{{$e.Set "description" (print "Account deleted")}}
							{{$e.Set "color" 0x00ff7b}}
						{{else}}
							{{$e.Set "description" (print "You do not have an account. Open one with `" $.Cmd " create`" )}}
						{{end}}
					{{else if eq $c "set"}}
						{{$c := false}}
						{{if $a.Get (toString $u)}}
							{{if gt (len $.CmdArgs) 1}}
								{{$v := (lower (toString (index $.CmdArgs 1)))}}
								{{$A := sdict}}
								{{$s := sdict}}
								{{if eq $v "w" "whitelist"}}
									{{if gt (len $.CmdArgs) 2}}
										{{$ID := (joinStr " " (slice $.CmdArgs 2))}}
										{{$us := cslice}}
										{{range (reFindAll `\d{17,19}` $ID)}}
											{{- if not (eq (toInt .) $u) -}}
												{{- $us = $us.Append .}}
											{{else}}
												{{$e.Set "description" (print "You cannot add yourself to the whitelist")}}
											{{end}}
										{{- end}}
										{{if $us}}
											{{$c = true}}
											{{$A = $a.Get (toString $u)}}
											{{$s = $A.Get "accountSettings"}}
											{{$w := $s.Get "whitelistedUsers"}}
											{{$s.Set "whitelistedUsers" ($w.AppendSlice $us)}}
											{{$A.Set "accountSettings" $s}}
											{{$a.Set (toString $u) $A}}
											{{$e.Set "description" (print "User(s) now added to whitelist.")}}
										{{else}}
											{{$e.Set "description" (print "No users provided.")}}
										{{end}}
									{{else}}
										{{$e.Set "description" (print "No user(s) provided.\nPlease ensure that you use uIDs")}}
									{{end}}
								{{else if (reFind `(with(draw)?limit|w)` $v)}}
									{{if gt (len $.CmdArgs) 2}}
										{{if and (toString (index $.CmdArgs 2)) (ge (toInt (index $.CmdArgs 2)) 0)}}
											{{$c = true}}
											{{$A = $a.Get (toString $u)}}
											{{$s = $A.Get "accountSettings"}}
											{{$s.Set "withdrawimit" (index $.CmdArgs 2)}}
											{{$e.Set "description" (print "New withdraw limit added.")}}
											{{$A.Set "accountSettings" $s}}
											{{$a.Set (toString $u) $A}}
										{{else}}
											{{$e.Set "description" (print "Invalid amount provided.")}}
										{{end}}
									{{else}}
										{{$e.Set "description" (print "No amount provided.")}}
									{{end}}
								{{else}}
									{{$e.Set "description" (print "Invalid option provided.")}}
								{{end}}
							{{else}}
								{{$e.Set "description" (print "No option argument passed.\nSyntax is: `" $.Cmd " <Option> <vLues>`")}}
							{{end}}
							{{if $c}}
								{{$e.Set "color" 0x00ff7b}}
								{{dbSet 0 "accounts" $a}}
							{{end}}
						{{else}}
							{{$e.Set "description" (print "You have no account. Open one with `" $.Cmd " create` or wait to be added to a whitelist.")}}
						{{end}}
					{{else if eq $c "list"}}
						{{$L := cslice}}
						{{if $a.Get (toString $u)}}
							{{$L = $L.Append $u}}
						{{end}}
						{{range $C}}
							{{- $A := . -}}
							{{- $w := (($a.Get .).accountSettings).whitelistedUsers -}}
							{{range $w}}
								{{- if in . (toString $u) -}}
									{{- $L = $L.Append $A -}}
								{{end}}
							{{end}}
						{{end}}
						{{if $L}}
							{{$f := cslice}}
							{{range $L}}
								{{- $f = $f.Append (sdict "Name" (print " Account ") "Value"  (print "`" . "`") "inline" false) -}}
							{{end}}
							{{$e.Set "fields" $f}}
							{{$e.Set "color" 0x00ff7b}}
						{{else}}
							{{$e.Set "description" (print "You have no accounts. Open one with `" $.Cmd " create` or wait to be added to a whitelist.")}}
						{{end}}
					{{else if (reFind `(bal(ance)?|with(draw)?|dep(osit)?)` $c)}}
						{{if gt (len $.CmdArgs) 1}}
							{{$A := (index $.CmdArgs 1)}}
							{{if (toInt $A)}}
								{{if in $C (toString $A)}}
									{{$c := false}}
									{{if $a.Get (toString $u)}}
										{{- $c = true -}}
									{{end}}
									{{- $w := (($a.Get (toString $A)).accountSettings).whitelistedUsers -}}
									{{range $w}}
										{{- if in . (toString $u) -}}
											{{- $c = true -}}
										{{end}}
									{{end}}
									{{$b := ($a.Get (toString $A)).accountBalance}}
									{{if and ((reFind `bal(ance)?` $c)) $c}}
										{{$e.Set "description" (print "The account `" $A "` has a balance of " (humanizeThousands $b))}}
										{{$e.Set "color" 0x00ff7b}}
									{{else if and ((reFind `with(draw)?|dep(osit)?` $c)) $c}}
										{{if gt (len $.CmdArgs) 2}}
											{{$p := (index $.CmdArgs 2)}}
											{{$lm := (toInt ($a.Get (toString $A)).accountSettings.withdrawimit)}}
											{{if gt (toInt $p) 0}}
												{{$p = toInt $p}}
												{{$cash := or (dbGet $u "cash").Value 0 | toInt}}
												{{$N := ""}}
												{{$E := ""}}
												{{if (reFind `with(draw)?` $c)}}
													{{$c = "withdraw"}}
												{{else if (reFind `dep(osit)?` $c)}}
													{{$c = "deposit"}}
												{{end}}
												{{if le $p $lm}}
													{{$c := false}}
													{{$n := ""}}
													{{if eq $c "with" "withdraw"}}
														{{$n = "the account has"}}
														{{if le $p (toInt $b)}}
															{{$c = "withdrawn"}}
															{{$N = add (toInt $b.cash) $p}}
															{{$E = sub $b $p}}
														{{end}}
													{{else if eq $c "dep" "deposit"}}
														{{$n = "you have"}}
														{{if le $p (toInt $b.cash)}}
															{{$c = "deposited"}}
															{{$N = sub (toInt $b.cash) $p}}
															{{$E = add $b $p}}
														{{end}}
													{{end}}
													{{if $c}}
														{{$b := $a.Get (toString $A)}}
														{{$b.Set "accountBalance" $E}}
														{{$e.Set "description" (print "You've just " $c " " $p)}}
														{{$e.Set "color" 0x00ff7b}}
														{{$a.Set $A $b}}
														{{dbSet 0 "accounts" $a}}
													{{else}}
														{{$e.Set "description" (print "You can't " $c " more than " $n)}}
													{{end}}
												{{else}}
													{{$e.Set "description" (print "You can't " $c " more than " $lm " at one time")}}
												{{end}}
												{{dbSet $u "cash" $cash}}
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
							{{$e.Set "description" (print "No account provided")}}
						{{end}}
					{{else}}
						{{$e.Set "description" (print "Invalid `option` argument passed" $o)}}
					{{end}}
				{{else}}
					{{$e.Set "description" (print "No `account` database found.\nPlease set it up with the default va;ues using `" $.Cmd " initialize`")}}
				{{end}}
			{{end}}
		{{end}}
	{{else}}
		{{$e.Set "description" (print "No option argument passed." $o)}}
	{{end}}
{{else}}
	{{$e.Set "description" (print "No `Settings` database found.\nPlease set it up with the default values using `" .ServerPrefix "server-set default`")}}
{{end}}
{{sendMessage nil (cembed $e)}}