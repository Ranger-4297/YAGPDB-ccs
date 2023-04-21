{{/*
        Made by Ranger (765316548516380732)

    Trigger Type: `Regex`
    Trigger: `\A(-|<@!?204255221017214977>\s*)(transfer)(\s+|\z)`

    ©️ Ranger 2020-Present
    GNU, GPLV3 License
    Repository: https://github.com/Ranger-4297/YAGPDB-ccs
*/}}


{{/* Only edit below if you know what you're doing (: rawr */}}

{{/* Transfers V6 data to V7 format */}}
{{$economyInfo := (dbGet .User.ID "EconomyInfo").Value}}
{{$cash := or $economyInfo.cash 0}}
{{$bank := or $economyInfo.bank 0}}
{{$inventory:= or $economyInfo.inventory sdict}}
{{$streaks := or $economyInfo.streaks (sdict "daily" 0 "weekly" 0 "monthly" 0)}}

{{dbSet .User.ID "cash" $cash}}

{{dbSet .User.ID "userEconData" (sdict "inventory" $inventory "streaks" $streaks)}}

{{$bankDB := or (dbGet 0 "bank").Value sdict}}
{{$bankDB.Set (toString .User.ID) $bank}}
{{dbSet 0 "bank" $bankDB}}

{{dbDel .User.ID "EconomyInfo"}}