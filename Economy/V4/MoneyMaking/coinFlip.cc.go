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
                    {{$bet := (index . 1)}}
                    {{if $bet | toInt}}
                        {{$bet = $bet | toInt}}
                        {{if lt $bet 0}}
                            {{$embed.Set "description" (print "You cannot flip for lower than " $symbol "1")}}
                            {{$embed.Set "color" $errorColor}}
                        {{else}}
                            {{with (dbGet $userID "EconomyInfo")}}
                                {{$a = sdict .Value}}
                                {{$cash := $a.cash}}
                                {{$newCashBalance := ""}}
                                {{$int := randInt 1 3}}
                                {{if eq $int 1}} {{/* Win */}}
                                    {{$newCashBalance = $cash | add $bet}}
                                    {{$embed.Set "description" (print "You flipped " $side1 " and won " $symbol $bet)}}
                                    {{$embed.Set "color" $successColor}}
                                {{else}} {{/* Lose */}}
                                    {{$newCashBalance = $bet | sub $cash}}
                                    {{$embed.Set "description" (print "You flipped " $side1 " and lost.")}}
                                    {{$embed.Set "color" $errorColor}}
                                {{end}}
                                {{$sdict := (dbGet $userID "EconomyInfo").Value}}
                                {{$sdict.Set "cash" $newCashBalance}}
                                {{dbSet $userID "EconomyInfo" $sdict}}
                            {{end}}
                        {{end}}
                    {{else}}
                        {{$embed.Set "description" (print "No valid `Bet` argument provided.\nSyntax is `" $.Cmd " <Side:Head/Tails> <Bet:Amount>`")}}
                        {{$embed.Set "color" $errorColor}}
                    {{end}}
                {{else}}
                    {{$embed.Set "description" (print "No `Bet` argument provided.\nSyntax is `" $.Cmd " <Side:Head/Tails> <Bet:Amount>`")}}
                    {{$embed.Set "color" $errorColor}}
                {{end}}
            {{else}}
                {{$embed.Set "description" (print "No valid `Side` argument provided.\nSyntax is `" $.Cmd " <Side:Head/Tails> <Bet:Amount>`")}}
                {{$embed.Set "color" $errorColor}}
            {{end}}
        {{end}}
    {{else}}
        {{$embed.Set "description" (print "No `Side` argument provided.\nSyntax is `" $.Cmd " <Side:Head/Tails> <Bet:Amount>`")}}
        {{$embed.Set "color" $errorColor}}
    {{end}}
{{else}}
    {{$embed.Set "description" (print "No `Settings` database found.\nPlease set it up with the default values using `" $prefix "set default`")}}
    {{$embed.Set "color" $errorColor}}
{{end}}
{{sendMessage nil (cembed $embed)}}