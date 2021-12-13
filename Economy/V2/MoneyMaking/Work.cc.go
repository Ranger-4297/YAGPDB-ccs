{{/*
		Made by Ranger (765316548516380732)

	Trigger Type: `Regex`
	Trigger: `\A(-|<@!?204255221017214977>\s*)(work|job|getpaid|labor)(\s+|\z)`

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

{{/* Works an honest job */}}

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
	{{$min := (toInt $a.min)}}
	{{$max := (toInt $a.max)}}
	{{$symbol := $a.symbol}}
	{{if not (dbGet $userID "EconomyInfo")}}
		{{dbSet $userID "EconomyInfo" (sdict "cash" 200 "bank" 0)}}
	{{end}}
	{{with (dbGet $userID "EconomyInfo")}}
		{{$a = sdict .Value}}
		{{$cash := $a.cash}}
		{{$workPay := (randInt $min) (randInt $max)}}
		{{$newCashBalance := $cash | add $workPay}}
		{{sendMessage nil (cembed
            "author" (sdict "name" $.User.Username "icon_url" ($.User.AvatarURL "1024"))
            "description" (print "You decided to work today! You got paid a hefty " $symbol $workPay)
            "color" 0x00ff7b
            "timestamp" currentTime
            )}}
		{{$sdict := (dbGet $userID "EconomyInfo").Value}}
        {{$sdict.Set "cash" $newCashBalance}}
        {{dbSet $userID "EconomyInfo" $sdict}}
	{{end}}
{{else}}
    {{sendMessage nil (cembed
            "author" (sdict "name" $.User.Username "icon_url" ($.User.AvatarURL "128"))
            "description" (print "No `Settings` database found.\nPlease set it up with the default values using `" $prefix "set default`")
            "color" $errorColor
            "timestamp" currentTime
            )}}
{{end}}