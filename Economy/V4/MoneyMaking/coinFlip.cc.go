{{/*
		Made by Ranger (765316548516380732)

	Trigger Type: `Regex`
	Trigger: `\A(-|<@!?204255221017214977>\s*)(coin-?flip|cf|flip-?coin|fc)(\s+|\z)`

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

{{/* Flip */}}

{{/*
If the users aren't in the economy database 
It'll automatically add them
--
If there is no setting values
You'll be asked to set it up with default values
You can change these later
*/}}

{{/* Response */}}
{{$embed := sdict}}
{{$embed.Set "author" (sdict "name" $.User.Username "icon_url" ($.User.AvatarURL "1024"))}}
{{$embed.Set "timestamp" currentTime}}
{{with (dbGet 0 "EconomySettings")}}
    {{$a := sdict .Value}}
    {{$symbol := $a.symbol}}
    {{if not (dbGet $userID "EconomyInfo")}}
        {{dbSet $userID "EconomyInfo" (sdict "cash" 200 "bank" 0)}}
    {{end}}
    {{with $.CmdArgs}}
        {{if (index . 0)}}
            {{$side1 :=  (index . 0) | toString | lower}}
            {{$picker1 := ""}}
            {{$win := ""}}
            {{$lose := ""}}
            {{if (reFind `(t(ails?)?|h(eads?)?)` $side1)}}
                {{if eq $side1 "t" "tails" "tail"}}
                    {{$side1 = "tails"}}
                {{else if eq $side1 "h" "heads" "head"}}
                    {{$side1 = "heads"}}
                {{end}}
                {{if gt (len .) 1}}
                    {{$amount := (index . 1)}}
                    {{if $amount | toInt}}
                        {{$amount = $amount | toInt}}
                        {{if lt $amount 0}}
                            {{$embed.Set "description" (print "You cannot flip for lower than " $symbol "1")}}
                            {{$embed.Set "color" $errorColor}}
                        {{else}}
                            {{with (dbGet $userID "EconomyInfo")}}
                                {{$a = sdict .Value}}
                                {{$cash := $a.cash}}
                                {{$newCashBalance := ""}}
                                {{$int := randInt 1 3}}
                                {{if eq $int 1}} {{/* Win */}}
                                    {{$newCashBalance = $cash | add $amount}}
                                    {{$embed.Set "description" (print "You flipped " $side1 " and won " $symbol $amount)}}
                                    {{$embed.Set "color" $successColor}}
                                {{else}} {{/* Lose */}}
                                    {{$newCashBalance = $amount | sub $cash}}
                                    {{$embed.Set "description" (print "You flipped " $side1 " and lost.")}}
                                    {{$embed.Set "color" $errorColor}}
                                {{end}}
                                {{$sdict := (dbGet $userID "EconomyInfo").Value}}
                                {{$sdict.Set "cash" $newCashBalance}}
                                {{dbSet $userID "EconomyInfo" $sdict}}
                            {{end}}
                        {{end}}
                    {{else}}
                        {{$embed.Set "description" (print "No valid `Amount` argument provided.\nSyntax is `" $.Cmd " <Side:Head/Tails> <Amount:Amount>`")}}
                        {{$embed.Set "color" $errorColor}}
                    {{end}}
                {{else}}
                    {{$embed.Set "description" (print "No `Amount` argument provided.\nSyntax is `" $.Cmd " <Side:Head/Tails> <Amount:Amount>`")}}
                    {{$embed.Set "color" $errorColor}}
                {{end}}
            {{else}}
                {{$embed.Set "description" (print "No valid `Side` argument provided.\nSyntax is `" $.Cmd " <Side:Head/Tails> <Amount:Amount>`")}}
                {{$embed.Set "color" $errorColor}}
            {{end}}
        {{end}}
    {{else}}
        {{$embed.Set "description" (print "No `Side` argument provided.\nSyntax is `" $.Cmd " <Side:Head/Tails> <Amount:Amount>`")}}
        {{$embed.Set "color" $errorColor}}
    {{end}}
{{else}}
    {{$embed.Set "description" (print "No `Settings` database found.\nPlease set it up with the default values using `" $prefix "set default`")}}
    {{$embed.Set "color" $errorColor}}
{{end}}
{{sendMessage nil (cembed $embed)}}