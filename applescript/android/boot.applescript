on run argv

set deviceName to (item 1 of argv) 
set appName to "Terminal"

if application appName is running then
    tell application "Terminal"
        set currentTab to do script "cd ~/Library/Android/sdk/emulator;"
        delay 2
        do script "./emulator @" & deviceName & " -logcat '*:s GSM:i'" in front window
        activate
    end tell    
    return "Wiping device: " & deviceName
else
    tell application "Terminal"
        set currentTab to do script "cd ~/Library/Android/sdk/emulator;"
        delay 2
        do script "./emulator @" & deviceName & " -logcat '*:s GSM:i'" in front window
        activate
    end tell    
    return "Wiping device: " & deviceName
end if

end run