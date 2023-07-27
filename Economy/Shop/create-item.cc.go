{{/*
		Made by Ranger (765316548516380732)
		Wait response logic by DZ (438789314101379072)

	Trigger Type: `Regex`
	Trigger: `.*`

	©️ Ranger 2020-Present
	GNU, GPLV3 License
	Repository: https://github.com/Ranger-4297/YAGPDB-ccs

	Note: Command is `create-item`/`new-item`. Use your severs default prefix
*/}}

{{/* Only edit below if you know what you're doing (: rawr */}}

{{/* Initiates variables */}}
{{$uI := .User.ID}}
{{$eC := 0xFF0000}}
{{$dV := toInt (dbGet .User.ID "waitResponse").Value}}
{{$cS := 0}}
{{$t := print `(` .ServerPrefix `|<@!?204255221017214977>\s*)((create|new)-?item)`}} 

{{/* Create item */}}

{{/* Response */}}
{{$e := sdict "title" "Item info" "footer" (sdict "text" "Type cancel to cancel the setup") "color" 0x00ff7b "timestamp" currentTime}}
{{$s := or (dbGet 0 "store").Value (sdict "items" sdict)}}
{{$i := $s.items}}
{{$iD := sdict "desc" "" "price" 0 "quantity" 0 "role" 0 "replyMsg" "" "expiry" 0}}
{{$p := split (index (split (exec "viewperms") "\n") 2) ", "}}
{{if not .ExecData}}
	{{with $eco := (dbGet 0 "EconomySettings").Value}}
		{{$sB := $eco.symbol}}
		{{if not $dV}}
			{{/* checks if message matches regex to begin tutorial */}}
			{{if reFind (print `\A(?i)` $t `(\s+|\z)`) $.Message.Content}}
				{{if or (in $p "Administrator") (in $p "ManageServer")}}
					{{$e.Set "fields" (cslice (sdict "name" "Name" "value" "⠀⠀"))}}
					{{$cS = 1}}
					{{$e := sendMessageRetID nil (complexMessage "content" "Please enter a name for the item (under 60 characters)" "embed" (cembed $e))}}
					{{dbSet 0 "createItem" (sdict "embed" (toString $e) "item" (sdict "name" "" "data" $iD) "user" (toString $uI))}}
					{{scheduleUniqueCC $.CCID nil 120 1 1}}
				{{else}}
					{{$e.Del "footer"}}{{$e.Del "title"}}
					{{$e.Set "author" (sdict "name" $.User.Username "icon_url" ($.User.AvatarURL "1024"))}}
					{{$e.Set "description" (print "Insufficient permissions.\nTo use this command you need to have either `Administrator` or `ManageServer` permissions")}}
					{{$e.Set "color" $eC}}
					{{sendMessage nil (cembed $e)}}
				{{end}}
			{{end}}
		{{else}}
			{{$cE := (dbGet 0 "createItem").Value}}
			{{if and $cE.user (eq (toString $cE.user) (toString $uI))}}
				{{if eq $dV 1}}
					{{if and (le (len (toRune $.Message.Content)) 60) (not (eq $.Message.Content "cancel"))}}
						{{$name := $.Message.Content}}
						{{$im := $cE.item}}
						{{$im.Set "name" $name}}
						{{dbSet 0 "createItem" $cE}}
						{{$e.Set "fields" (cslice (sdict "name" "Name" "value" $name "inline" true))}}
						{{editMessage nil (dbGet 0 "createItem").Value.embed (complexMessageEdit "content" "Please enter a price for the item" "embed" (cembed $e))}}
						{{$cS = 1}}
					{{else if not (eq $.Message.Content "cancel")}}
						{{$m := sendMessageRetID nil "Please try again and enter name under 60 characters"}}
						{{deleteTrigger 0}}
						{{deleteMessage nil $m 10}}
						{{$cS = 0}}
					{{end}}
					{{scheduleUniqueCC $.CCID nil 120 1 2}}
				{{else if eq $dV 2}}
					{{with toInt $.Message.Content}}
						{{if gt . 0}}
							{{$price := .}}
							{{$im := $cE.item}}
							{{$im.data.Set "price" $price}}
							{{dbSet 0 "createItem" $cE}}
							{{$m := structToSdict (index (getMessage nil (dbGet 0 "createItem").Value.embed).Embeds 0)}}
							{{$f := sdict "name" "Price" "value" (print $sB .) "inline" true}}
							{{$e.Set "fields" ((cslice.AppendSlice $m.Fields).Append $f)}}
							{{editMessage nil (dbGet 0 "createItem").Value.embed (complexMessageEdit "content" "Please enter a description for the item (under 200 characters)" "embed" (cembed $e))}}
							{{$cS = 1}}
						{{else}}
							{{$m := sendMessageRetID nil "Please try again and enter a valid number"}}
							{{deleteTrigger 0}}
							{{deleteMessage nil $m 10}}
							{{$cS = 0}}
						{{end}}
					{{else if not (eq (lower $.Message.Content) "cancel")}}
						{{$m := sendMessageRetID nil "Please try again and enter a valid number"}}
						{{deleteTrigger 0}}
						{{deleteMessage nil $m 10}}
						{{$cS = 0}}
					{{end}}
					{{scheduleUniqueCC $.CCID nil 120 1 3}}
				{{else if or (eq $dV 3) (eq $dV 6)}}
					{{if and (le (len (toRune $.Message.Content)) 200) (not (eq $.Message.Content "cancel"))}}
						{{$o := ""}}
						{{$fV := ""}}
						{{if eq $dV 3}}
							{{$o = "desc"}}
							{{$fV = "description"}}
						{{else}}
							{{$o = "replyMsg"}}
							{{$fV = "reply"}}
						{{end}}
						{{$im := $cE.item}}
						{{$im.data.Set $o $.Message.Content}}
						{{$m := structToSdict (index (getMessage nil (dbGet 0 "createItem").Value.embed).Embeds 0)}}
						{{$f := sdict "name" $fV "value" $.Message.Content}}
						{{$e.Set "fields" ((cslice.AppendSlice $m.Fields).Append $f)}}
						{{if eq $dV 3}}
							{{editMessage nil (dbGet 0 "createItem").Value.embed (complexMessageEdit "content" "How much of this item should the store stock?\nIf unlimited just reply `skip` or `inf`" "embed" (cembed $e))}}
						{{else}}
							{{editMessage nil (dbGet 0 "createItem").Value.embed (complexMessageEdit "content" "When should this item expire from the users inventory if not used\nIf never just reply `skip` or `inf`" "embed" (cembed $e))}}
						{{end}}
						{{dbSet 0 "createItem" $cE}}
						{{$cS = 1}}
						{{scheduleUniqueCC $.CCID nil 120 1 4}}
					{{else if not (eq $.Message.Content "cancel")}}
						{{$m := sendMessageRetID nil "Please try again and enter name under 60 characters"}}
						{{deleteTrigger 0}}
						{{deleteMessage nil $m 10}}
						{{$cS = 0}}
						{{scheduleUniqueCC $.CCID nil 120 1 5}}
					{{end}}
				{{else if eq $dV 4}}
					{{if or (gt (toInt $.Message.Content) 0) (eq (lower $.Message.Content) "inf" "skip")}}
						{{$qty := $.Message.Content}}
						{{if eq $qty "skip" "inf"}}
							{{$qty = "inf"}}
						{{end}}
						{{$im := $cE.item}}
						{{$im.data.Set "quantity" $qty}}
						{{$m := structToSdict (index (getMessage nil (dbGet 0 "createItem").Value.embed).Embeds 0)}}
						{{$f := sdict "name" "Stock" "value" (toString $qty) "inline" true}}
						{{$e.Set "fields" ((cslice.AppendSlice $m.Fields).Append $f)}}
						{{editMessage nil (dbGet 0 "createItem").Value.embed (complexMessageEdit "content" "What role do you want to be given when this item is used?\nType `skip` to skip this." "embed" (cembed $e))}}
						{{$cS = 1}}
						{{dbSet 0 "createItem" $cE}}
					{{else if not (eq $.Message.Content "cancel")}}
						{{$m := sendMessageRetID nil "Please try again and enter a valid number"}}
						{{deleteTrigger 0}}
						{{deleteMessage nil $m 10}}
						{{$cS = 0}}
					{{end}}
					{{scheduleUniqueCC $.CCID nil 120 1 6}}
				{{else if eq $dV 5}}
					{{$r := $.Message.Content}}
					{{if or ($.Guild.GetRole (toInt64 $r)) (eq (lower $r) "skip")}}
						{{if eq $r "skip"}}
							{{$r = 0}}
						{{end}}
						{{$im := $cE.item}}
						{{$im.data.Set "role" $r}}
						{{dbSet 0 "createItem" $cE}}
						{{if not $r}}
							{{$r = "none"}}
						{{else}}
							{{$r = print "<@&" $r ">"}}
						{{end}}
						{{$m := structToSdict (index (getMessage nil (dbGet 0 "createItem").Value.embed).Embeds 0)}}
						{{$f := sdict "name" "Role-given" "value" (toString $r) "inline" true}}
						{{$e.Set "fields" ((cslice.AppendSlice $m.Fields).Append $f)}}
						{{editMessage nil (dbGet 0 "createItem").Value.embed (complexMessageEdit "content" "What message should I reply with when the item is used? (under 200 character)" "embed" (cembed $e))}}
						{{$cS = 1}}
					{{else if not (eq (lower $r) "cancel")}}
						{{$m := sendMessageRetID nil "Please try again with a valid roleID"}}
						{{deleteTrigger 0}}
						{{deleteMessage nil $m 10}}
						{{$cS = 0}}
					{{end}}
					{{scheduleUniqueCC $.CCID nil 120 1 7}}
				{{else if eq $dV 7}}
					{{scheduleUniqueCC $.CCID nil 120 1 7}}
					{{$d := $.Message.Content}}
					{{if or (toDuration $d) (eq (lower $d) "skip" "inf")}}
						{{$eV := "none"}}
						{{if (toDuration $d)}}
							{{$d = (toDuration $d).Seconds}}
							{{$eV = humanizeDurationSeconds (mult $d $.TimeSecond)}}
						{{else}}
							{{$d = ""}}
						{{end}}
						{{$im := $cE.item}}
						{{$im.data.Set "expiry" $d}}
						{{$i.Set $cE.item.name $cE.item.data}}
						{{$m := structToSdict (index (getMessage nil (dbGet 0 "createItem").Value.embed).Embeds 0)}}
						{{$f := sdict "name" "Inventory expiry" "value" (toString $eV) "inline" true}}
						{{$e.Set "fields" ((cslice.AppendSlice $m.Fields).Append $f)}}
						{{editMessage nil (dbGet 0 "createItem").Value.embed (complexMessageEdit "content" "Item created! ✅" "embed" (cembed $e))}}
						{{dbSet 0 "store" $s}}
						{{dbDel 0 "createItem"}}
						{{dbDel $uI "waitResponse"}}
						{{cancelScheduledUniqueCC $.CCID 1}}
					{{else}}
						{{$m := sendMessageRetID nil "Please try again with a valid duration"}}
						{{deleteTrigger 0}}
						{{deleteMessage nil $m 10}}
						{{$cS = 0}}
					{{end}}
				{{end}}
				{{if eq (lower $.Message.Content) "cancel"}}
					{{sendMessage nil "Create-item was cancelled"}}
					{{dbDel $uI "waitResponse"}}
					{{dbDel 0 "createItem"}}
					{{$cS = 0}}
					{{cancelScheduledUniqueCC $.CCID 1}}
				{{end}}
			{{end}}
		{{end}}
	{{else}}
		{{$e.Del "footer"}}{{$e.Del "title"}}
		{{$e.Set "author" (sdict "name" $.User.Username "icon_url" ($.User.AvatarURL "1024"))}}
		{{$e.Set "description" (print "No `Settings` database found.\nPlease set it up with the default values using `" $.ServerPrefix "server-set default`")}}
		{{$e.Set "color" $eC}}
		{{sendMessage nil (cembed $e)}}
	{{end}}
{{else}}
	{{sendMessage nil "Create-item was cancelled"}}
	{{dbDel .User.ID "waitResponse"}}
	{{dbDel 0 "createItem"}}
{{end}}
{{if $cS}}
	{{dbSetExpire .User.ID "waitResponse" (str (add $dV 1)) 120}}
{{end}}