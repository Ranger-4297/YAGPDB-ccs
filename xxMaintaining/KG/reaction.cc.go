{{/*
		Made by ranger_4297 (765316548516380732)

	Trigger Type: `Reaction`
	Trigger: `Added reaction only`

	©️ Ranger 2020-Present
	GNU, GPLV3 License
*/}}


{{/* Configuration values start */}}
{{$enrollingRole := 1190603830065373224}}		{{/* roleID of ENROLLING role */}}
{{$onboardingRole := 1189413709454520360}}		{{/* roleID of onboarding role */}}
{{$serverRole := 1185703179254505582}}			{{/* roleID of SERVER role */}}
{{$allianceRole := 1185727965397516379}}		{{/* roleID of ALLIANNCE role */}}
{{$nameRole := 1185703261613863013}}			{{/* roleID of NAME role */}}
{{$rankRole := 1190609302218620938}}			{{/* roleID of RANK role */}}
{{$verifyRole := 1185731311927828540}}			{{/* roleID of VERIFIED role */}}
{{$languageChannel := 1187971295397281862}}		{{/* channelID of LANGUAGE channel*/}}
{{$guideChannel := 1189408420714004611}}		{{/* channelID of GUIDE channel*/}}
{{$serverChannel := 1185342810312945804}}		{{/* channelID of SERVER channel*/}}
{{$rankChannel := 1185723603086491648}}			{{/* channelID of RANK channel */}}
{{$rulesChannel := 1185730875644706826}}        {{/* channelID of RULES channel */}}
{{$publicChat := 1190146250125877288}}			{{/* channelID of PUBLIC CHAT channel*/}}
{{/* Configuration values end */}}

{{/* response */}}
{{sleep 1}}
{{$languages := cslice "english" "spanish" "french" "russian" "chinese" "arabic" "korean" "german" "vietnamese" "japanese" "turkish" "portugese" "malaysian" "fillipino" "ukranian" "indonesian" "greek" "dutch" "italian" "romanian" "danish" "polish" "hebrew"}}
{{$language := "english"}}
{{range (getMember .User.ID).Roles}}
	{{if in $languages (getRole .).Name}}
		{{$language = (getRole .).Name}}
		{{break}}
	{{end}}
{{end}}
{{$language = (dbGet 0 "languageDB").Value.Get $language}}
{{if eq .Channel.ID $languageChannel}}
	{{addRoleID $enrollingRole}}
	{{sleep 1}}
	{{$gm := sendMessageNoEscapeRetID $guideChannel (print (reReplace "<@!>" (reReplace `<#L>` (reReplace `<#S>` ($language.Get 20) (printf "<#%d>" $serverChannel)) (printf "<#%d>" 1187971295397281862)) .User.Mention))}}
	{{addMessageReactions $guideChannel $gm "✅"}}
	{{dbSetExpire .User.ID "Rrcooldown" true 60}}
	{{$m := sendMessageNoEscapeRetID $guideChannel (print (reReplace `<@!>` ($language.Get 15) .User.Mention))}}
	{{deleteMessage $guideChannel $m 60}}
{{else if eq .Channel.ID $guideChannel}}
	{{dbSetExpire .User.ID "Rrcooldown2" true 60}}
	{{addRoleID $onboardingRole}}
	{{$m := sendMessageNoEscapeRetID $guideChannel (print (reReplace `<#>` (reReplace `<@!>` ($language.Get 14) .User.Mention) (printf "<#%d>" $serverChannel)))}}
	{{deleteMessage $guideChannel $m 60}}
	{{$m2 := sendMessageNoEscapeRetID $serverChannel (print (reReplace `<@!>` ($language.Get 16) .User.Mention))}}
	{{deleteMessage $serverChannel $m2 60}}
{{else if eq .Channel.ID $rankChannel}}
	{{addRoleID $rankRole}}
	{{if $cd := dbGet .User.ID "Rrcooldown3"}}
		{{return}}
	{{end}}
	{{dbSetExpire .User.ID "Rrcooldown3" true 3}}
	{{$m := sendMessageNoEscapeRetID $rankChannel (print (reReplace `<#>` (reReplace `<@!>` ($language.Get 19) .User.Mention) (printf "<#%d>" $rulesChannel)))}}
	{{deleteMessage $rankChannel $m 60}}
	{{$m2 := sendMessageNoEscapeRetID $rulesChannel (print (reReplace `<@!>` ($language.Get 17) .User.Mention))}}
	{{deleteMessage $rulesChannel $m2 60}}
{{else if eq .Channel.ID $rulesChannel}}
	{{addRoleID $verifyRole}}
	{{range (cslice $serverRole $allianceRole $nameRole $enrollingRole $onboardingRole $rankRole)}}
		{{if (hasRoleID .)}}
			{{removeRoleID .}}
		{{end}}
	{{end}}
	{{if $cd := dbGet .User.ID "Rrcooldown4"}}
		{{return}}
	{{end}}
	{{dbSetExpire .User.ID "Rrcooldown5" true 60}}
	{{$m := sendMessageNoEscapeRetID $publicChat (print (reReplace `<@!>` ($language.Get 18) .User.Mention))}}
	{{deleteMessage $publicChat $m 60}}
{{end}}