{{/*
		Made by Ranger (765316548516380732)

	Trigger Type: `Regex`
	Trigger: `\A(-|<@!?204255221017214977>\s*)(offence|crime|commit-?felon)`
©️ Ranger 2021
MIT License
*/}}

{{/* Only edit below if you know what you're doing (: rawr */}}

{{/* Initiates variables */}}
{{$userID := .User.ID}}
{{$successColor := 0x00ff7b}}
{{$errorColor := 0xFF0000}}
{{$prefix := index (reFindAllSubmatches `.*?: \x60(.*)\x60\z` (execAdmin "Prefix")) 0 1 }}

{{/* Commits a crime with the chance of failing */}}

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
    {{$failRate := $a.failRate}}
    {{if not (dbGet $userID "EconomyInfo")}}
        {{dbSet $userID "EconomyInfo" (sdict "cash" 200 "bank" 0)}}
    {{end}}
    {{with (dbGet $userID "EconomyInfo")}}
        {{$a = sdict .Value}}
        {{$cash := $a.cash}}
        {{$amount := randInt $min $max}}
        {{$newCash := ""}}
        {{$int := randInt $failRate}}
        {{if gt $int (div $failRate 2)}}
            {{$newCash = $cash | add $amount}}
            {{sendMessage nil (cembed
            "author" (sdict "name" $.User.Username "icon_url" ($.User.AvatarURL "128"))
            "description" (print "You broke the law for a pretty penny! You made " $symbol $amount " in your crime spree today")
            "color" $successColor
            "timestamp" currentTime
            )}}
        {{else}}
            {{$newCash = $amount | sub $cash}}
            {{sendMessage nil (cembed
            "author" (sdict "name" $.User.Username "icon_url" ($.User.AvatarURL "128"))
            "description" (print "You broke the law trying to commit a felony! You were arrested and lost " $symbol $amount " due to your bail.")
            "color" $errorColor
            "timestamp" currentTime
            )}}
        {{end}}
        {{$sdict := (dbGet $userID "EconomyInfo").Value}}
        {{$sdict.Set "cash" $newCash}}
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