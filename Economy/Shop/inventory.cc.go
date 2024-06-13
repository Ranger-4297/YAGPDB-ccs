{{/*
        Made by ranger_4297 (765316548516380732)

    Trigger Type: `Regex`
    Trigger: `\A(-|<@!?204255221017214977>\s*)(inv(entory)?)(\s+|\z)`

    ©️ RhykerWells 2020-Present
    GNU, GPLV3 License
    Repository: https://github.com/Ranger-4297/YAGPDB-ccs
*/}}


{{/* Only edit below if you know what you're doing (: rawr */}}

{{/* Initiates variables */}}
{{$user := .User}}
{{$successColor := 0x00ff7b}}
{{$errorColor := 0xFF0000}}
{{$prefix := .ServerPrefix }}

{{/* Server shop */}}

{{/* Response */}}
{{$embed := sdict "author" (sdict "name" .User.Username "icon_url" (.User.AvatarURL "1024")) "timestamp" currentTime "color" $errorColor}}
{{$economySettings := (dbGet 0 "EconomySettings").Value}}
{{if not $economySettings}}
	{{$embed.Set "description" (print "No `Settings` database found.\nPlease set it up with the default values using `" .ServerPrefix "server-set default`")}}
	{{sendMessage nil (cembed $embed)}}
	{{return}}
{{end}}
{{$symbol := $economySettings.symbol}}
{{with .CmdArgs}}
	{{if $newUser := (index . 0) | userArg}}
		{{$user = $newUser}}
	{{end}}
{{end}}
{{$page := "1"}}
{{if gt (len .CmdArgs) 1}}
	{{$page = (index .CmdArgs 1) | toInt}}
	{{if lt $page 1}}
		{{$page = 1}}
	{{end}}
{{end}}
{{$userdata := or (dbGet $user.ID "userEconData").Value (sdict "inventory" sdict)}}
{{$invStatus := or $userdata.settings.inventory "yes" | toString}}
{{$inventory := $userdata.inventory}}
{{if not $inventory }}
	{{$embed.Set "description" (print "This users inventory is empty :(\nThey should get some items from the shop!")}}
{{end}}
{{$field := cslice}}
{{$entry := cslice}}
{{range $k,$v := $inventory}}
	{{$item := $k}}
	{{$desc := $v.desc}}
	{{$qty := $v.quantity}}
	{{$role := $v.role}}
	{{if $role}}
		{{$role = print "<@&" $role ">"}}
	{{else}}
		{{$role = "none"}}
	{{end}}
	{{$expiry := $v.expiry}}
	{{$expires := "never"}}
	{{if $expiry}}
		{{$timeSeconds := toDuration (humanizeDurationSeconds (mult $expiry .TimeSecond))}}
		{{$expires = (print "<t:" (currentTime.Add $timeSeconds).Unix ":f>")}}
	{{end}}
	{{$entry = $entry.Append (sdict "Name" $item "value" (joinStr "\n" (print "Description: " $desc) (print "Quantity: " (humanizeThousands $qty)) (print "Role given: " $role) (print "Expiry: " $expires)) "inline" false)}}
{{end}}
{{$start := (mult 10 (sub $page 1))}}
{{$stop := (mult $page 10)}}
{{if ge $stop (len $entry)}}
	{{$stop = (len $entry)}}
{{end}}
{{if and (le $start $stop) (ge (len $entry) $start) (le $stop (len $entry))}}
	{{range (seq $start $stop)}}
		{{$field = $field.Append (index $entry .)}}
	{{end}}
{{else}}
	{{$embed.Set "description" (print "This page is empty")}}
{{end}}
{{if or (eq .User.ID $user.ID) (eq $invStatus "yes")}}
	{{if $inventory}}
		{{$embed.Set "title" (print "Inventory")}}
		{{$embed.Set "fields" $field}}
		{{$embed.Set "color" $successColor}}
		{{$embed.Set "footer" (sdict "text" (print "Page: " $page))}}
	{{end}}
	{{if not (and (eq .User.ID $user.ID) (eq $invStatus "no"))}}
		{{sendMessage nil (cembed $embed)}}
		{{return}}
	{{end}}
	{{sendDM (cembed $embed)}}
	{{$embed.Set "description" "Sent this to your DM as your inventory is on private"}}
	{{$embed.Del "footer"}}
	{{$embed.Del "fields"}}
	{{sendMessage nil (cembed $embed)}}
{{else}}
	{{sendMessage nil "This user has their inventory on private :("}}
{{end}}